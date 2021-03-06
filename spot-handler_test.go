package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	rr := httptest.NewRecorder()

	r := server()
	routes(r)
	request, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	r.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
}

func TestSpotHandler(t *testing.T) {
	registerLogger()
	ec2auth()

	r := server()

	routes(r)
	t.Run("Check path exist", func(t *testing.T) {
		rr := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/api/v1/notice", nil)
		assert.NoError(t, err)

		r.ServeHTTP(rr, request)
		assert.NotEqual(t, 404, rr.Code)
	})

	t.Run("Accepts confirmation request", func(t *testing.T) {
		rr := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/api/v1/notice", nil)
		request.Header.Set("x-amz-sns-message-type", "SubscriptionConfirmation")
		assert.NoError(t, err)

		r.ServeHTTP(rr, request)

		assert.Equal(t, rr.Body.String(), "Confirmed")
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("Process Notice", func(t *testing.T) {
		rr := httptest.NewRecorder()

		body := strings.NewReader("{ \"Type\": \"Notification\", \"MessageId\": \"22b80b92-fdea-4c2c-8f9d-bdfb0c7bf324\", \"TopicArn\": \"arn:aws:sns:us-west-2:123456789012:MyTopic\", \"Subject\": \"My First Message\", \"Message\": { \"version\": \"0\", \"id\": \"1e5527d7-bb36-4607-3370-4164db56a40e\", \"detail-type\": \"EC2 Spot Instance Interruption Warning\", \"source\": \"aws.ec2\", \"account\": \"123456789012\", \"time\": \"1970-01-01T00:00:00Z\", \"region\": \"us-east-1\", \"resources\": [\"arn:aws:ec2:us-east-1b:instance/i-0b662ef9931388ba0\"], \"detail\": { \"instance-id\": \"i-0b65b0ddcc9c89298\", \"instance-action\": \"terminate\" } }, \"Timestamp\": \"2012-05-02T00:54:06.655Z\", \"SignatureVersion\": \"1\", \"Signature\": \"EXAMPLEw6JRN...\", \"SigningCertURL\": \"https://sns.us-west-2.amazonaws.com/SimpleNotificationService-f3ecfb7224c7233fe7bb5f59f96de52f.pem\", \"UnsubscribeURL\": \"https://sns.us-west-2.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-west-2:123456789012:MyTopic:c9135db0-26c4-47ec-8998-413945fb5a96\"}")
		request, err := http.NewRequest("POST", "/api/v1/notice", body)
		request.Header.Set("x-amz-sns-message-type", "Notification")
		assert.NoError(t, err)

		r.ServeHTTP(rr, request)
		fmt.Println("Body: " + rr.Body.String())

		assert.NotEmpty(t, rr.Body.String())
		assert.Equal(t, 200, rr.Code)
	})

}
