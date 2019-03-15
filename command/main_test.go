package main

import "testing"

func TestFire(t *testing.T) {
	expected := "pew pew"
	fire := &Fire{}
	out := fire.Execute()

	if out != expected {
		t.Errorf("expected '%s' but got '%s'", expected, out)
	}
}

func TestJump(t *testing.T) {
	expected := "JUUUUUUMP"
	jump := &Jump{}
	out := jump.Execute()

	if out != expected {
		t.Errorf("expected '%s' but got '%s'", expected, out)
	}
}

// Test input handler

func TestInputHandleFire(t *testing.T) {
	expected := "pew pew"
	out := InputHandler("fire")

	if out != expected {
		t.Errorf("expected '%s' but got '%s'", expected, out)
	}
}

func TestInputHandleJump(t *testing.T) {
	expected := "JUUUUUUMP"
	out := InputHandler("jump")

	if out != expected {
		t.Errorf("expected '%s' but got '%s'", expected, out)
	}
}

func TestInputHandleCommandNotFound(t *testing.T) {
	expected := "command not found"
	out := InputHandler("asdfasdfasdf")

	if out != expected {
		t.Errorf("expected '%s' but got '%s'", expected, out)
	}
}
