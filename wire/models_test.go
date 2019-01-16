package wire

import (
	"reflect"
	"testing"
)

func TestNewMessage(t *testing.T) {
	tests := []struct {
		name string
		want Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDI(t *testing.T) {

}
