package router

import (
	"bytes"
	"crypto/tls"
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
func (rs *RodentServer) RegisterWithServer(router Router) error {

	
	

	return nil

}

// Register a new rodent with primary C2 Server
func (rs *RodentServer) RegisterNewRodent(rdata rodent.Rodent) error {

	return nil
}

func (rs *RodentServer) GetRodentTasks(xid string) []rodent.Task {
	var mytasks []rodent.Task

	return mytasks
}

// Send API Request to complete or update a task
func (rs *RodentServer) UpdateRodentTask(xid string, rtask *rodent.Task) error {

	return nil

}

// Send Hearbeat for polling purposes.
func (rs *RodentServer) Heartbeat() {

}

// Send file to Rodent Server.
func RouterSendFile(xid string, filedata []byte) {

}

// Get File from Rodent Server.
func RouterGetFile(xid string) []byte {
	var filedat []byte

	return filedat

}
