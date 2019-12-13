package kubernetes

import (
	"context"
	"fmt"
	"github.com/fatih/flags"
	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"github.com/lnxjedi/gopherbot/bot"
	"github.com/nlopes/slack"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/startupheroes/kubebot/util"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

const (
	issueInfoMessage     = "\nPR Info: ```%s\n%s```"
	prNotApprovedMessage = "Looks like PR is not approved. Make sure that at least 2 reviewers approved the PR so I can deploy it."
)

var (
	lockEnvMap = make(map[string]envLock)
)

type Config struct {
	GithubToken            string
	GithubOrg              string
	SlackVerificationToken string
	Projects               []Project
	Environments           []string
	ProtectedEnvironments  []string
}

var githubClient *github.Client

func kubernetes(robot *bot.Robot, command string, args ...string) (retval bot.TaskRetVal) {
	var config *Config
	robot.GetTaskConfig(&config)
	if command == "init" {
		githubClient = newGithubClient(config.GithubToken)
		go func() {
			glog.Info("Starting prometheus metrics.")
			http.Handle("/metrics", promhttp.Handler())
			http.Handle("/slack/interactive", InteractionHandler{VerificationToken: config.SlackVerificationToken})
			glog.Warning(http.ListenAndServe(":8080", nil))
		}()
		return
	}

	if command == "list" {
		robot.Reply(listCommand(config.Projects))
		return bot.Success
	}
	if command == "heapdump" {
		podName := args[0]
		environment := args[1]
		output, _ := util.Execute("", "kubectl", "exec", "-it", podName, "-n"+environment, "heapdump")
		robot.Reply(output)
		return bot.Success
	}
	environment, err := getEnvironment(config.Environments, command, args...)
	if err != nil {
		robot.Reply(err.Error())
		return
	}

	project, err := getProject(config.Projects, command, args...)
	if err != nil {
		robot.Reply(err.Error())
		return
	}
	var response string
	switch command {
	case "deploy":
		version := args[0]
		merged := flags.Has("--merged", args)
		issueInfo, issueApproved := getPRInfo(config.GithubOrg, project.Repo, version, merged)
		if !issueApproved && !robot.CheckAdmin() && util.Contains(config.ProtectedEnvironments, environment) {
			robot.Reply(prNotApprovedMessage)
			return
		}

		message := ""
		if flags.Has("--message", args) {
			message, _ = flags.Value("--message", args)
		}
		envLockKey := getEnvLockKey(project.Name, environment)
		if envLock, exist := lockEnvMap[envLockKey]; exist && envLock.User != robot.User {
			client := robot.Incoming.Client.(*slack.Client)
			promptVerificationForDeploy(client, envLock, project, robot.User, environment, version, robot.Channel, message)
			return
		}
		output := deployCommand(project, robot.User, version, environment)
		response = output
		if issueInfo != "" {
			response += issueInfo
		}
		lockIfNotProd(environment, message, robot.User, version, project, envLockKey)
	case "rollback":
		response = rollbackCommand(project, environment)
	case "restart":
		response = restartCommand(project, environment)
	case "unlock":
		envLockKey := getEnvLockKey(project.Name, environment)
		if _, exist := lockEnvMap[envLockKey]; exist {
			delete(lockEnvMap, envLockKey)
		}
		response = fmt.Sprintf(unlockMessage, project.Name, environment)
	}
	robot.Reply(response)
	return bot.Success
}

func lockIfNotProd(environment string, messageInfo string, user, version string, project Project, envLockKey string) {
	if environment != "prod" {
		location, _ := time.LoadLocation("Europe/Moscow")
		lock := envLock{
			Message:     messageInfo,
			User:        user,
			Environment: environment,
			Version:     version,
			Time:        time.Now().In(location).Format("2006-01-02 15:04"),
			App:         project.Name,
		}
		lockEnvMap[envLockKey] = lock
	}
}

func getPRInfo(owner, repo, version string, merged bool) (string, bool) {
	opt := &github.SearchOptions{}
	searchCondition := "is:open"
	if merged {
		searchCondition = "is:merged"
	}
	query := fmt.Sprintf("review:approved is:pr %s repo:%s/%s %s", searchCondition, owner, repo, version)
	searchResult, _, _ := githubClient.Search.Issues(context.Background(), query, opt)
	approved := len(searchResult.Issues) > 0
	if approved {
		return fmt.Sprintf(issueInfoMessage, *searchResult.Issues[0].Title, *searchResult.Issues[0].HTMLURL), approved
	}
	return "", approved
}

func newGithubClient(token string) *github.Client {
	return github.NewClient(oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		),
	))
}

func init() {
	bot.RegisterPlugin("kubernetes", bot.PluginHandler{
		Handler: kubernetes,
		Config:  &Config{},
	})
}
