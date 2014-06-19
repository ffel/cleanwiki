package interfaces

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ffel/cleanwiki/usecase"
)

func TestPageFileRepo_existing(t *testing.T) {
	repo := PageFileRepo{}

	notes := usecase.Notes{&Logger{}, &repo}

	orig := []byte("walking on the grass with bare feet")

	notes.Add("aFineTitle", orig)

	file, err := ioutil.ReadFile("aFineTitle.txt")

	if err != nil {
		t.Errorf("unexpected error %d\n", err)
	}

	if bytes.Compare(orig, file) != 0 {
		t.Error("unexpected difference in contents")
	}

	body, err := notes.Read("aFineTitle")

	if err != nil {
		t.Errorf("unexpected error %d\n", err)
	}

	if bytes.Compare(orig, body) != 0 {
		t.Error("unexpected difference in contents")
	}

	// clean up file
	_ = os.Remove("aFineTitle.txt")
}

// logger mock - now it needs to be public for usecases to use

type Logger struct {
	log string
}

func (l *Logger) Log(message string) error {
	l.log += message + "\n"

	fmt.Println(message)

	return nil
}
