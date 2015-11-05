package event

import "testing"

func TestCreateMachine(t *testing.T) {
	m := NewMachine()
	m.Start()
}

func TestRunningEvent(t *testing.T) {
	done := make(chan bool)

	m := NewMachine()

	m.OnceEvent("core::running", func(e Event) {
		t.Log(e)
		done <- true
	})

	m.Start()
	<-done
}

func TestPersonalEvent(t *testing.T) {
	done := make(chan bool)

	m := NewMachine()

	m.OnceEvent("user::myevent", func(e Event) {
		t.Log(e)
		done <- true
	})

	m.Push(CreateEvent("user::myevent"))

	m.Start()
	<-done
}

func TestManyEvents(t *testing.T) {
	done := make(chan bool)

	m := NewMachine()

	i := 0
	m.OnEvent("user::myevent", func(e Event) {
		t.Log(e)

		if i >= 2 {
			done <- true
		} else {
			i++
		}
	})

	m.Push(CreateEventWithData("user::myevent", []byte("A")))
	m.Push(CreateEventWithData("user::myevent", []byte("B")))
	m.Push(CreateEventWithData("user::myevent", []byte("C")))

	m.Start()
	<-done
}

func TestStatusCommand(t *testing.T) {
	done := make(chan bool)

	m := NewMachine()

	m.Command("core::status", func(e Event) {
		t.Log(e)
		done <- true
	})

	m.Start()
	<-done
}

func TestCreateCommand(t *testing.T) {
	done := make(chan bool)

	m := NewMachine()

	m.OnceEvent("user::test", func(ev Event) {
		t.Log(ev)
		newEvent := CreateResponseEvent(ev)
		m.Push(newEvent)
	})

	m.Command("user::test", func(ev Event) {
		t.Log(ev)
		done <- true
	})

	m.Start()
	<-done
}
