package router

import (
	"crypto/tls"
	"github.com/RodentFramework/core/rodent"
)

/*
Helper functions for Routers that are written in golang.


*/

type RodentServer struct {
	TlsConf tls.Config
	Address string
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
