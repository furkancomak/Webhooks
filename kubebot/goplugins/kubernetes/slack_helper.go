package kubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/nlopes/slack"
	"net/http"
)

const (
	actionDoNothing      = "do_nothing"
	actionDeployItAnyway = "deploy_it_anyway"
	actionAsk            = "ask_to_user"
	actionYes            = "yes"
	actionNo             = "no"
	deployCallbackId     = "deploy"
	askCallbackId        = "ask"
	lockedEnvironment    = "`%s` is locked by `%s` at `%s` by `deploy %s %s/%s`\n" +
		"What would you like me to do ?"
	lockedEnvironmentWithMessage = "Deployed with: `%s`\n" + lockedEnvironment
	askResponse                  = "`%s` asks you to unlock & deploy `%s` by `deploy %s %s/%s` \n" +
		"What would you like me to do ?"
	askResponseWithMessage = "`%s` asks you to unlock & deploy `%s` by `deploy %s %s/%s` with\n" +
		"`%s`\n" +
		"what would you like me to do?"
)

func promptVerificationForDeploy(client *slack.Client, envLock envLock, project Project, user, environment, version, channel, message string) {
	deployInfo := deployInfo{project, user, environment, version, message, envLock}
	attachment := newVerificationAttachment(deployInfo)
	messageAttachment := slack.MsgOptionAttachments(attachment)
	channelID, timestamp, err := client.PostMessage(channel, slack.MsgOptionText("", false), messageAttachment)
	if err != nil {
		glog.Errorf("Could not send message: %v", err)
	}
	glog.Infof("Message with buttons sucessfully sent to channel %s at %s", channelID, timestamp)
	return
}

func newVerificationAttachment(deployInformation deployInfo) slack.Attachment {
	bytes, _ := json.Marshal(deployInformation)
	deployInfoString := string(bytes)
	preText := getPreText(deployInformation.Lock)

	attachment := slack.Attachment{
		Pretext:    preText,
		Fallback:   "We don't currently support your client",
		CallbackID: deployCallbackId,
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			{
				Name:  actionDoNothing,
				Text:  "Do nothing",
				Type:  "button",
				Value: deployInfoString,
				Style: "primary",
			},
			{
				Name:  actionAsk,
				Text:  "Poke @" + deployInformation.Lock.User,
				Type:  "button",
				Value: deployInfoString,
				Style: "default",
			},
			{
				Name:  actionDeployItAnyway,
				Text:  "Deploy it anyway",
				Type:  "button",
				Value: deployInfoString,
				Style: "danger",
			},
		},
	}
	return attachment
}
func newAskAttachment(askInformation deployInfo) slack.Attachment {
	stringByte, _ := json.Marshal(askInformation)
	askString := string(stringByte)
	preText := getAskPreText(askInformation)
	attachment := slack.Attachment{
		Pretext:    preText,
		Fallback:   "We don't currently support your client",
		CallbackID: askCallbackId,
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			{
				Name:  actionYes,
				Text:  "Do it",
				Type:  "button",
				Value: askString,
				Style: "primary",
			},
			{
				Name:  actionNo,
				Text:  "Don't do it",
				Type:  "button",
				Value: askString,
				Style: "danger",
			},
		},
	}
	return attachment
}

func getPreText(lock envLock) string {
	if lock.Message != "" {
		return fmt.Sprintf(lockedEnvironmentWithMessage, lock.Message, lock.Environment, lock.User, lock.Time, lock.Version,
			lock.App, lock.Environment)
	}
	return fmt.Sprintf(lockedEnvironment, lock.Environment, lock.User, lock.Time, lock.Version, lock.App, lock.Environment)
}

func getAskPreText(askInfo deployInfo) string {
	if askInfo.Message != "" {
		return fmt.Sprintf(askResponseWithMessage, askInfo.User, askInfo.Lock.Environment, askInfo.Version,
			askInfo.Lock.App, askInfo.Environment, askInfo.Message)
	}
	return fmt.Sprintf(askResponse, askInfo.User, askInfo.Lock.Environment, askInfo.Version, askInfo.Lock.App,
		askInfo.Environment)
}

func handleDeployCallback(payload slack.InteractionCallback) (slack.Msg, int) {
	var deployInfo deployInfo
	err := json.Unmarshal([] byte(payload.Actions[0].Value), &deployInfo)
	if err != nil {
		message := "Could not extract deploy info from payload"
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		return response, http.StatusOK
	}
	user := deployInfo.User
	environment := deployInfo.Environment
	project := deployInfo.Project
	envLockKey := getEnvLockKey(project.Name, environment)
	version := deployInfo.Version
	lockUser := deployInfo.Lock.User
	action := payload.Actions[0].Name
	switch action {
	case actionDeployItAnyway:
		message := "Batman has no limits, but you do, sir."
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		if user == payload.User.Name {
			message := deployCommand(project, user, version, environment)
			lockIfNotProd(environment, deployInfo.Message, user, version, project, envLockKey)
			response = newResponse(user, message, slack.Attachment{})
		}
		return response, http.StatusOK

	case actionAsk:
		message := "This request is not yours."
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		if user == payload.User.Name {
			attachment := newAskAttachment(deployInfo)
			response = newResponse(lockUser, "", attachment)
		}
		return response, http.StatusOK
	case actionDoNothing:
		message := "Today, I don't want to."
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		if user == payload.User.Name {
			message := "Of course, as you wish"
			response = newResponse(user, message, slack.Attachment{})
		}
		return response, http.StatusOK

	default:
		message := "This action is not supported."
		returnMessage := newResponse(payload.User.Name, message, slack.Attachment{})
		return returnMessage, http.StatusNotAcceptable
	}
}
func handleAskCallback(payload slack.InteractionCallback) (slack.Msg, int) {
	var deployInfo deployInfo
	err := json.Unmarshal([] byte(payload.Actions[0].Value), &deployInfo)
	if err != nil {
		message := "Could not extract deploy info from payload"
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		return response, http.StatusOK
	}
	user := deployInfo.User
	environment := deployInfo.Environment
	project := deployInfo.Project
	envLockKey := getEnvLockKey(project.Name, environment)
	version := deployInfo.Version
	lockUser := deployInfo.Lock.User
	action := payload.Actions[0].Name
	switch action {
	case actionYes:
		message := "Know your limits, sir."
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		if lockUser == payload.User.Name {
			message := deployCommand(project, user, version, environment)
			lockIfNotProd(environment, deployInfo.Message, user, version, project, envLockKey)
			response = newResponse(deployInfo.User, message, slack.Attachment{})
		}
		return response, http.StatusOK
	case actionNo:
		message := "You crossed the line, sir."
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		if lockUser == payload.User.Name {
			message := "Okay, It stays locked"
			response = newResponse(payload.User.Name, message, slack.Attachment{})
		}
		return response, http.StatusOK
	default:
		message := "This action is not supported."
		response := newResponse(payload.User.Name, message, slack.Attachment{})
		return response, http.StatusNotAcceptable
	}
}

func newResponse(user, message string, attachment slack.Attachment) slack.Msg {
	return slack.Msg{
		Text:            "<@" + user + ">, " + message,
		ReplaceOriginal: true,
		ResponseType:    "in_channel",
		Attachments:     []slack.Attachment{attachment}}
}
