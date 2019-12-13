package main

import (
	"github.com/lnxjedi/gopherbot/bot"
	_ "github.com/lnxjedi/gopherbot/brains/file"
	_ "github.com/lnxjedi/gopherbot/connectors/slack"
	_ "github.com/lnxjedi/gopherbot/goplugins/knock"
	_ "github.com/lnxjedi/gopherbot/goplugins/ping"
	_ "github.com/startupheroes/kubebot/goplugins/ci"
	_ "github.com/startupheroes/kubebot/goplugins/kubernetes"
	_ "github.com/startupheroes/kubebot/goplugins/vpn"
)

var versionInfo = bot.VersionInfo{
	Version: "v2.0.0",
	Commit:  "(manual build)",
}

func main() {
	bot.Start(versionInfo)
}
