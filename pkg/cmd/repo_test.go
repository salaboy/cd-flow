package cmd

import (
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

func Test_setExtensionForRepoEvents(t *testing.T) {

	repoName = "foo"
	cloudEvent := createTestCloudEvent()
	setExtensionForRepoEvents(cloudEvent)

	assert.Equal(t, "foo", cloudEvent.Extensions()["cdfreponame"])
}

func createTestCloudEvent() cloudevents.Event {
	event := cloudevents.NewEvent()
	event.SetID("abc-123") //Generate with UUID
	event.SetSource("cdf-events")
	event.SetType("CDF.Repository.Created")
	event.SetTime(time.Now())
	return event
}
