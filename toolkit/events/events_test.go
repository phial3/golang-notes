package events_test

import (
	"github.com/phial3/gotoolkit/events"
	"testing"
)

type UserCreatedEvent struct {
	ID       string
	UserName string
}

type UserUpdatedEvent struct {
	ID string
}

func TestPublishEvent(t *testing.T) {
	eventManager := events.NewEventManager(events.NewMemoryEventStore(false))
	eventManager.Listen(func(evt UserCreatedEvent) {
		t.Logf("user created: id=%s, name=%s", evt.ID, evt.UserName)

		if evt.ID != "111" {
			t.Error("test failed")
		}
	})

	eventManager.Listen(func(evt UserUpdatedEvent) {
		t.Logf("user updated-1: id=%s", evt.ID)

		if evt.ID != "121" {
			t.Error("test failed")
		}
	})

	eventManager.Listen(func(evt UserUpdatedEvent) {
		t.Logf("user updated-2: id=%s", evt.ID)

		if evt.ID != "121" {
			t.Error("test failed")
		}
	})

	eventManager.Publish(UserCreatedEvent{
		ID:       "111",
		UserName: "李逍遥",
	})

	eventManager.Publish(UserUpdatedEvent{
		ID: "121",
	})
}
