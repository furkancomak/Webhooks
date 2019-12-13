package vpn

import (
	"fmt"
	"github.com/startupheroes/kubebot/util"
	"github.com/lnxjedi/gopherbot/bot"
	"github.com/nlopes/slack"
	"strings"
)

func vpn(robot *bot.Robot, command string, args ...string) (retval bot.TaskRetVal) {
	if command == "init" { // ignore init
		return
	}
	switch command {
	case "generate":
		robot.Reply("I'll send a DM when i am done. Please wait...")
		clientName := args[0]
		podName, _ := util.Execute("", "kubectl", "get", "pod", "-nvpn", "-l", "app=openvpn-operator", "-o", "jsonpath='{.items[0].metadata.name}'")
		podName = strings.Replace(podName, "'", "", -1)
		output, _ := util.Execute("", "kubectl", "exec", "-nvpn", podName, "--", command, clientName)
		robot.SendUserMessage(robot.User, output)
		fileName := clientName + ".ovpn"
		content, _ := util.Execute("", "kubectl", "exec", "-nvpn", podName, "--", "cat", "/etc/openvpn/client_configs/"+fileName)
		userId := robot.GetSenderAttribute("internalID").Attribute
		client := robot.Incoming.Client.(*slack.Client)
		params := slack.FileUploadParameters{
			Filename: fileName,
			Content:  content,
			Channels: []string{util.StripUserId(userId)},
		}
		_, err := client.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
	case "revoke":
		clientName := args[0]
		podName, _ := util.Execute("", "kubectl", "get", "pod", "-nvpn", "-l", "app=openvpn-operator", "-o", "jsonpath='{.items[0].metadata.name}'")
		podName = strings.Replace(podName, "'", "", -1)
		output, _ := util.Execute("", "kubectl", "exec", "-nvpn", podName, "--", command, clientName)
		robot.Reply(output)
	default:
		robot.Reply("Unsupported operation: " + command)
	}
	return bot.Success
}

func init() {
	bot.RegisterPlugin("vpn", bot.PluginHandler{
		Handler:       vpn,
	})
}
