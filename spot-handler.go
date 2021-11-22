package main

import (
	"fmt"
	"os"
	"strconv"

	"os/exec"
	"strings"
	"time"
)

type SNSMessage struct {
	Type             string    `json:"Type"`
	MessageID        string    `json:"MessageId"`
	TopicArn         string    `json:"TopicArn"`
	Subject          string    `json:"Subject"`
	Message          Notice    `json:"Message"`
	Timestamp        time.Time `json:"Timestamp"`
	SignatureVersion string    `json:"SignatureVersion"`
	Signature        string    `json:"Signature"`
	SigningCertURL   string    `json:"SigningCertURL"`
	UnsubscribeURL   string    `json:"UnsubscribeURL"`
}

type Notice struct {
	Version    string    `json:"version"`
	ID         string    `json:"id"`
	DetailType string    `json:"detail-type"`
	Source     string    `json:"source"`
	Account    string    `json:"account"`
	Time       time.Time `json:"time"`
	Region     string    `json:"region"`
	Resources  []string  `json:"resources"`
	Detail     struct {
		InstanceID     string `json:"instance-id"`
		InstanceAction string `json:"instance-action"`
	} `json:"detail"`
}

func (n Notice) GetRequestId() string {
	return n.ID
}
func (n Notice) GetInstanceId() string {
	return strings.Trim(n.Detail.InstanceID, " ")
}

func (n Notice) GetInstanceAction() string {
	return n.Detail.InstanceAction
}

func (n Notice) ExecuteDrain(hostname string) {
	Logger.Info(fmt.Sprintf(
		"Executing node drain for instance %s hostname %s on request id %s for action %s",
		n.GetInstanceId(),
		hostname,
		n.GetRequestId(),
		n.GetInstanceAction()))

	context := ""
	if os.Getenv("KUBECTL_CONFIG") != "" {
		context = os.Getenv("KUBECTL_CONFIG")
	}
	
	// kubectl --kubeconfig /var/lib/kubelet/kubeconfig drain node_name
	command := fmt.Sprintf("kubectl --kubeconfig %s  drain %s", context, hostname)
	// command := "sleep 10 && echo 'done'"
	Logger.Info("executing:" + command)
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Start(); err != nil {
		Logger.Error(err.Error())
		return
	}
	pid := cmd.Process.Pid
	Logger.Info("Command running with pid: " + strconv.Itoa(pid))
	go func() {
		err := cmd.Wait()
		Logger.Info(fmt.Sprintf("Command finished with error: %v", err))
	}()
}
