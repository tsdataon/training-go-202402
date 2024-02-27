package main

import "testing"

func TestPerson(t *testing.T) {
	var john = person{
		name: "John Doe",
		age:  25,
	}
	status(&john)

	if john.name != "John Doe" {
		t.Error("Expected John Doe but got", john.name)
	}

	if john.age != 25 {
		t.Error("Expected 25 but got", john.age)
	}

	if john.status != "belum menikah" {
		t.Error("Expected belum menikah but got", john.status)
	}
}
