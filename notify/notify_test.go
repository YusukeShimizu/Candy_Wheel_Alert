package notify

import (
	"testing"

	"github.com/Candy_Wheel_Alert/env"
)

func TestNotify(t *testing.T) {
	config, err := env.Process()
	if err != nil {
		t.Fatal(err)
	}
	notifyer, err := NewNotifyer(&config)
	if err != nil {
		t.Fatal(err)
	}
	notifyer.Notify("Hello Candy!")
}
