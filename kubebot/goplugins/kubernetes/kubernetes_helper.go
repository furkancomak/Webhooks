package kubernetes

import (
	"fmt"
	"github.com/startupheroes/kubebot/util"
	"strings"
)

const (
	unlockMessage = "OK, i have just unlocked `%s/%s`."
	okResponse    = "Roger that!\nthis is the response to your request: ```%s``` "
	ListName      = "```%s```"
)

type envLock struct {
	User        string
	Environment string
	Time        string
	Version     string
	App         string
	Message     string
}
type deployInfo struct {
	Project     Project
	User        string
	Environment string
	Version     string
	Message     string
	Lock		envLock
}

func deployCommand(project Project, user, version, environment string) string {
	var response string
	deployments := project.Deployment
	for appName, imageName := range deployments {
		outputMessage, _ := deploy(appName, imageName, version, environment)
		response += outputMessage + "\n"
		err := doTag(imageName, version, environment, appName)
		if err != nil {
			response += "Can not tag " + version + "of " + appName + " as latest-in" + environment + ". Be advise\n"
		}
	}
	addPrometheusEvent(project.Name, environment, user, version)
	return fmt.Sprintf(okResponse, response)
}

func getEnvLockKey(projectName string, environment string) string {
	return projectName + "|" + environment
}

func restartCommand(project Project, environment string) string {
	var response string
	deployments := project.Deployment
	for appName := range deployments {
		outputMessage, _ := restart(appName, environment)
		response += outputMessage + "\n"
	}
	return fmt.Sprintf(okResponse, response)
}

func rollbackCommand(project Project, environment string) string {
	var response string
	deployments := project.Deployment
	for appName := range deployments {
		outputMessage, _ := rollback(appName, environment)
		response += outputMessage + "\n"
	}
	return fmt.Sprintf(okResponse, response)
}

func listCommand(projects []Project, ) string {
	var response = fmt.Sprintf("|%-20s|%-45s|\n", "NAME", "ALIAS")
	for _, v := range projects {
		name := v.Name
		alias := strings.Join(v.Alias, ", ")
		response += fmt.Sprintf("|%-20s|%-45s|\n", name, alias)

	}
	return fmt.Sprintf(ListName, response)
}

func deploy(appName string, imageName string, version string, environment string) (string, int) {
	return util.Execute("", "kubectl", "set", "image", "deployment", appName, "application="+imageName+":"+version, "--namespace="+environment)
}

func restart(appName string, environment string) (string, int) {
	return util.Execute("", "kubectl", "rollout", "restart", "deployment", appName, "--namespace="+environment)
}

func rollback(appName string, environment string) (string, int) {
	return util.Execute("", "kubectl", "rollout", "undo", "deployment", appName, "--namespace="+environment)
}

type Error struct {
	message string
}

func (error *Error) Error() string {
	return error.message
}
