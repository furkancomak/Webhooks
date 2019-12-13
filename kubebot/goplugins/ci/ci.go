package ci

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/flags"
	"github.com/golang/glog"
	"github.com/lnxjedi/gopherbot/bot"
	"github.com/startupheroes/kubebot/util"
	"net/http"
	"strings"
	"unicode"
)

type CircleCIRequest struct {
	BuildParameters map[string]string `json:"build_parameters"`
}

const apiUrl = "https://circleci.com/api/v1.1/project/github/startupheroes/startupheroes-release?circle-token="
const apiTestUrl = "https://circleci.com/api/v1.1/project/github/startupheroes/testify?circle-token="

type Config struct {
	Token              string
	CircleJob          string
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	DockerUser         string
	DockerPass         string
	ReleaseResponseUrl string
	TestResponseUrl    string
}

func ci(r *bot.Robot, command string, args ...string) (retval bot.TaskRetVal) {
	if command == "init" { // ignore init
		return
	}
	var config *Config
	r.GetTaskConfig(&config)
	switch command {
	case "release":
		release(r, config, args...)
	case "test":
		test(r, config, args...)
	}
	return bot.Success
}

func release(r *bot.Robot, config *Config, args ...string) {
	userId := r.GetSenderAttribute("internalID").Attribute
	request := newReleaseRequest(args, config, userId, r)
	doRequest(request, r, apiUrl+config.Token)
	r.Reply("Your release request has been passed to CircleCI.")
}

func test(r *bot.Robot, config *Config, args ...string) {
	userId := r.GetSenderAttribute("internalID").Attribute
	request := newTestRequest(args, config, userId, r)
	doRequest(request, r, apiTestUrl+config.Token)
	r.Reply("Your test request has been passed to CircleCI.")
}

func doRequest(request *CircleCIRequest, r *bot.Robot, apiUrl string) {
	marshal, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(marshal))
	if err != nil {
		r.Reply("Could prepare JSON request :(")
		glog.Error("Error: ", err)
		return
	}
	client := http.Client{}
	req.Header.Add("Content-Type", "application/json")
	glog.Info("Json was %v", request)
	resp, err := client.Do(req)
	if err != nil {
		r.Reply("Could not pass your request to CircleCi :(")
		glog.Error("Resp was", resp)

	}
}

func newReleaseRequest(args []string, config *Config, userId string, r *bot.Robot) *CircleCIRequest {
	buildParameters := make(map[string]string)
	buildParameters["text"] = strings.Join(args, " ")
	buildParameters["response_url"] = config.ReleaseResponseUrl
	buildParameters["user_id"] = util.StripUserId(userId)
	buildParameters["user_name"] = r.User
	request := &CircleCIRequest{BuildParameters: buildParameters}
	return request
}

func newTestRequest(args []string, config *Config, userId string, r *bot.Robot) *CircleCIRequest {
	buildParameters := make(map[string]string)
	buildParameters["CIRCLE_JOB"] = config.CircleJob
	buildParameters["AWS_REGION"] = config.AwsRegion
	buildParameters["AWS_ACCESS_KEY_ID"] = config.AwsAccessKeyId
	buildParameters["AWS_SECRET_ACCESS_KEY"] = config.AwsSecretAccessKey
	buildParameters["DOCKER_USER"] = config.DockerUser
	buildParameters["DOCKER_PASS"] = config.DockerPass
	buildParameters["RESPONSE_URL"] = config.TestResponseUrl
	buildParameters["USER_ID"] = util.StripUserId(userId)
	buildParameters["USER_NAME"] = r.User
	isAll := flags.Has("all", args)
	if !isAll {
		tags, tagErr := flags.Value("tags", args)
		if tagErr == nil {
			buildParameters["TEST_TAGS"] = "--tags " + tags
		}

	}
	apps, appErr := flags.Value("apps", args)
	if appErr == nil {
		appsWithVersion := strings.Split(stripSpace(apps), ",")
		for _, a := range appsWithVersion {
			app := strings.Split(a, ":")
			if len(app) > 1 {
				buildParameters[strings.ToUpper(app[0])+"_APP"] = app[1]
			}
		}
	}
	request := &CircleCIRequest{BuildParameters: buildParameters}
	return request
}

func stripSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func init() {
	bot.RegisterPlugin("ci", bot.PluginHandler{
		Handler: ci,
		Config:  &Config{},
	})
}
