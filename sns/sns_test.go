package sns

import (
	"bytes"
	"log"
	"net/http"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"launchpad.net/gocheck"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	gocheck.TestingT(t)
}

var _ = gocheck.Suite(&RequestSuite{})

type RequestSuite struct {
	request              *http.Request
	expectedNotification *Notification
}

func (self *RequestSuite) SetUpTest(c *gocheck.C) {
	body := `{
  "Type" : "Notification",
  "MessageId" : "da41e39f-ea4d-435a-b922-c6aae3915ebe",
  "TopicArn" : "arn:aws:sns:us-east-1:123456789012:MyTopic",
  "Subject" : "test",
  "Message" : "test message",
  "Timestamp" : "2012-04-25T21:49:25.719Z",
  "SignatureVersion" : "1",
  "Signature" : "EXAMPLElDMXvB8r9R83tGoNn0ecwd5UjllzsvSvbItzfaMpN2nk5HVSw7XnOn/49IkxDKz8YrlH2qJXj2iZB0Zo2O71c4qQk1fMUDi3LGpij7RCW7AW9vYYsSqIKRnFS94ilu7NFhUzLiieYr4BKHpdTmdD6c0esKEYBpabxDSc=",
  "SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-f3ecfb7224c7233fe7bb5f59f96de52f.pem",
  "UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:123456789012:MyTopic:2bcfbf39-05c3-41de-beaa-fcfcc21c8f55"
}`
	buf := bytes.NewReader([]byte(body))
	req, err := http.NewRequest("POST", "http://example.com", buf)

	if err != nil {
		log.Fatal("Failed to build request:", err)
	}

	self.request = req

	self.expectedNotification = &Notification{
		Type:             "Notification",
		MessageId:        "da41e39f-ea4d-435a-b922-c6aae3915ebe",
		TopicArn:         "arn:aws:sns:us-east-1:123456789012:MyTopic",
		Subject:          "test",
		Message:          "test message",
		Timestamp:        "2012-04-25T21:49:25.719Z",
		SignatureVersion: "1",
		Signature:        "EXAMPLElDMXvB8r9R83tGoNn0ecwd5UjllzsvSvbItzfaMpN2nk5HVSw7XnOn/49IkxDKz8YrlH2qJXj2iZB0Zo2O71c4qQk1fMUDi3LGpij7RCW7AW9vYYsSqIKRnFS94ilu7NFhUzLiieYr4BKHpdTmdD6c0esKEYBpabxDSc=",
		SigningCertURL:   "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-f3ecfb7224c7233fe7bb5f59f96de52f.pem",
		UnsubscribeURL:   "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:123456789012:MyTopic:2bcfbf39-05c3-41de-beaa-fcfcc21c8f55",
	}
}

func (self *RequestSuite) TestNewFromRequest(c *gocheck.C) {
	Convey("Parse SNS POST request into Notification struct", c, func() {

		req := self.request

		notification := NewFromRequest(req)
		expectedNotification := self.expectedNotification

		So(notification, ShouldHaveSameTypeAs, expectedNotification)
		So(notification.Type, ShouldEqual, expectedNotification.Type)
		So(notification.MessageId, ShouldEqual, expectedNotification.MessageId)
		So(notification.TopicArn, ShouldEqual, expectedNotification.TopicArn)
		So(notification.Subject, ShouldEqual, expectedNotification.Subject)
		So(notification.Message, ShouldEqual, expectedNotification.Message)
		So(notification.Timestamp, ShouldEqual, expectedNotification.Timestamp)
		So(notification.SignatureVersion, ShouldEqual, expectedNotification.SignatureVersion)
		So(notification.Signature, ShouldEqual, expectedNotification.Signature)
		So(notification.SigningCertURL, ShouldEqual, expectedNotification.SigningCertURL)
		So(notification.UnsubscribeURL, ShouldEqual, expectedNotification.UnsubscribeURL)
	})

}
