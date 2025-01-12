package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Wilderman"
	want := regexp.MustCompile(`\b` + name + `\b`)
	fmt.Println("want: ", want)
	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Wilderman") = %q, %v, quiere coincidencia para %#q, nil`, message, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, message, err)
	}
}
