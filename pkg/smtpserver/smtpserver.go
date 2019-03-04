package smtpserver

import (
	"fmt"

	guerrilla "github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/log"
)

// https://github.com/flashmob/go-guerrilla/wiki/Using-as-a-package#starting-a-server---custom-listening-interface

var smtpDaemon *guerrilla.Daemon

// Start server
func Start() error {
	appConfig := &guerrilla.AppConfig{
		LogFile: log.OutputStdout.String(),
	}

	serverConfig := guerrilla.ServerConfig{
		ListenInterface: "127.0.0.1:2525",
		IsEnabled:       true,
	}
	appConfig.Servers = append(appConfig.Servers, serverConfig)

	smtpDaemon = &guerrilla.Daemon{Config: appConfig}

	if err := smtpDaemon.Start(); err != nil {
		return err
	}

	fmt.Println("SMTP server started")

	return nil
}

// Stop server
func Stop() {
	smtpDaemon.Shutdown()
	fmt.Println("SMTP server closed")
}
