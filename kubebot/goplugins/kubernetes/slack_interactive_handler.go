package kubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
	"strings"
)

// interactionHandler handles interactive message response.
type InteractionHandler struct {
	VerificationToken string
}

func (h InteractionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/slack/interactive" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("incorrect path: %s", r.URL.Path)))
		return
	}

	// slack API calls the data POST a 'payload'
	reply := r.PostFormValue("payload")
	if len(reply) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("could not find payload"))
		return
	}

	var payload slack.InteractionCallback
	err := json.NewDecoder(strings.NewReader(reply)).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("could not process payload"))
		return
	}

	if payload.Token != h.VerificationToken {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("verification token does not match"))
		return
	}

	callback := payload.CallbackID
	var response slack.Msg
	var status int
	switch callback {
	case deployCallbackId:
		response, status = handleDeployCallback(payload)
	case askCallbackId:
		response, status = handleAskCallback(payload)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}
