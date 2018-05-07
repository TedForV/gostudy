package bridge

import (
	"errors"
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}

func TestPrinterAPI2(t *testing.T) {
	testWriter := TestWriter{}
	api2 := PrinterImpl2{
		&testWriter,
	}
	expectedMessage := "Hello"
	err := api2.PrintMessage(expectedMessage)
	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n Actual:%s\n Expected: %s\n", testWriter.Msg, expectedErrorMessage)

		}
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello  io.Writer"

	normal := NormalPrinter{
		expectedMessage,
		&PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		expectedMessage,
		&PrinterImpl2{
			&testWriter,
		},
	}
	err = normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Errorf("Error message was not correct.\n Actual:%s\n Expected: %s\n", testWriter.Msg, expectedMessage)
	}
}
