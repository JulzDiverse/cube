package sink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
	"github.com/julz/cube"
	"github.com/julz/cube/opi"
)

func Convert(
	msg cc_messages.DesireAppRequestFromCC,
	registryUrl string,
	cfClient cube.CfClient,
	client *http.Client,
	log lager.Logger,
) opi.LRP {
	if len(msg.ProcessGuid) > 36 {
		msg.ProcessGuid = msg.ProcessGuid[:36]
	}

	if msg.DockerImageUrl == "" {
		msg.DockerImageUrl = dropletToImageURI(msg, cfClient, client, registryUrl, log)
	}

	return opi.LRP{
		Name:            msg.ProcessGuid,
		Image:           msg.DockerImageUrl,
		TargetInstances: msg.NumInstances,
	}
}

func dropletToImageURI(
	msg cc_messages.DesireAppRequestFromCC,
	cfClient cube.CfClient,
	client *http.Client,
	registryUrl string,
	log lager.Logger,
) string {
	var appInfo cube.AppInfo
	for _, v := range msg.Environment {
		if v.Name == "VCAP_APPLICATION" {
			err := json.Unmarshal([]byte(v.Value), &appInfo)
			if err != nil {
				log.Error("failed-to-decode-environment-json-from-cc_message", err)
				panic(err)
			}
		}
	}

	dropletBytes, err := cfClient.GetDropletByAppGuid(appInfo.AppGuid)
	if err != nil {
		log.Error("failed-to-get-droplet-from-cloud-controller", err, lager.Data{"app-guid": appInfo.AppGuid})
		panic(err)
	}

	digest := stageRequest(client, registryUrl, appInfo, msg.DropletHash, dropletBytes, log)

	return fmt.Sprintf("10.244.0.142:8080/cloudfoundry/app-name@%s", digest)
}

func stageRequest(
	client *http.Client,
	registryUrl string,
	appInfo cube.AppInfo,
	dropletHash string,
	dropletBytes []byte,
	log lager.Logger,
) string {
	registryStageUri := registryStageUri(registryUrl, appInfo.SpaceName, appInfo.AppName, dropletHash)

	log.Info("sending-request-to-registry", lager.Data{"request": registryStageUri})

	req, err := http.NewRequest("POST", registryStageUri, bytes.NewReader(dropletBytes))
	if err != nil {
		log.Error("failed-to-create-http-request", err, nil)
		panic(err)
	}

	req.Header.Set("Content-Type", "application/gzip")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("stage-request-to-registry-failed", err, lager.Data{"request": registryStageUri})
		return ""
	}

	log.Info("request-successful", lager.Data{"response_status": resp.StatusCode})

	digest, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("read-response-failed", err)
		return ""
	}

	return string(digest)

}

func dropletDownloadUri(baseUrl string, appGuid string) string {
	return fmt.Sprintf("%s/v2/apps/%s/droplet/download", baseUrl, appGuid)
}

func registryStageUri(baseUrl string, space string, appname string, guid string) string {
	return fmt.Sprintf("%s/v2/%s/%s/blobs/?guid=%s", baseUrl, space, appname, guid)
}
