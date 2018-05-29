package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"code.cloudfoundry.org/runtimeschema/cc_messages"

	"github.com/julienschmidt/httprouter"
	"github.com/julz/cube/opi"
	"github.com/julz/cube/sink"
)

func runLRPHandler(c *sink.Converger) {
	handler := createHandler(c)
	http.ListenAndServe("0.0.0.0:8076", handler)
}

func createHandler(c *sink.Converger) http.Handler {
	handler := httprouter.New()

	lrpHandler := LRPHandler{
		converger: c,
	}
	handler.POST("/v1/lrp", lrpHandler.Desire)
	handler.GET("/v1/lrps", lrpHandler.List)

	return handler
}

type ListLRPResponseBody struct {
	Infos []DesiredLRPSchedulingInfo `json:"desired_lrp_scheduling_infos"`
}

type DesiredLRPSchedulingInfo struct {
	Key DesiredLRPKey `json:"desired_lrp_key"`
}

type DesiredLRPKey struct {
	Guid string `json:"process_guid"`
}

type LRPHandler struct {
	converger *sink.Converger
}

func (l *LRPHandler) Desire(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("failed to read request body", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	var msg cc_messages.DesireAppRequestFromCC
	if err := json.Unmarshal(body, &msg); err != nil {
		fmt.Println("failed to deserialize request body", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := l.converger.ConvergeOnce(context.Background(), []cc_messages.DesireAppRequestFromCC{msg}); err != nil {
		fmt.Println("failed to desire lrp", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusAccepted)
}

func (l *LRPHandler) List(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	lrps, err := l.converger.Desirer.List(context.Background())
	if err != nil {
		fmt.Println("failed to list lrp", err.Error())
		resp.WriteHeader(http.StatusInternalServerError)
	}
	responseBody := toResponseBody(lrps)
	body, err := json.Marshal(responseBody)
	if err != nil {
		fmt.Println("failed to serialize list response body", responseBody)
		resp.WriteHeader(http.StatusInternalServerError)
	}

	resp.Write(body)
	resp.WriteHeader(http.StatusOK)
}

func toResponseBody(lrps []opi.LRP) *ListLRPResponseBody {
	infos := []DesiredLRPSchedulingInfo{}
	for _, l := range lrps {
		info := DesiredLRPSchedulingInfo{
			Key: DesiredLRPKey{
				Guid: l.Name,
			},
		}
		infos = append(infos, info)
	}

	return &ListLRPResponseBody{
		Infos: infos,
	}
}
