package kubernetes

import (
	"fmt"
	"github.com/startupheroes/kubebot/util"
)

const (
	unknownEnvironmentMessage = "There is no environment called '%s'"
	unknownProjectMessage     = "We don't have project named '%s', do we?"
)

type Project struct {
	Name           string
	Repo           string
	Alias          []string
	Deployment     map[string]string
}

func getProject(projects []Project, command string, args ...string) (Project, error) {

	var alias string
	switch command {
	case "deploy":
		alias = args[1]
	default:
		alias = args[0]
	}

	for _, project := range projects {
		if util.Contains(project.Alias, alias) {
			return project, nil

		}
	}
	return Project{}, &Error{fmt.Sprintf(unknownProjectMessage, alias)}

}

func getEnvironment(environments []string, command string, args ...string) (string, error) {
	var environment string
	switch command {
	case "deploy":
		environment = args[2]
	default:
		environment = args[1]
	}

	if !util.Contains(environments, environment) {
		return "", &Error{fmt.Sprintf(unknownEnvironmentMessage, environment)}
	}
	return environment, nil
}
