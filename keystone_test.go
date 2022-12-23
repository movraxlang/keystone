package keystone

import (
	"fmt"
	"testing"
)

func TestNewEngine(t *testing.T) {
	_, err := NewEngine()
	if err != nil {
		t.Fatalf("error creating engine: %v", err)
	}
}

func TestAssemble(t *testing.T) {
	tests := []struct {
		code string
		want []string
	}{
		{"mov x0, #0x0",
			[]string{
				"00", "00", "80", "d2"},
		},
	}

	for _, tt := range tests {
		eng, err := NewEngine()
		if err != nil {
			t.Fatalf("%q creating engine: %v", tt.code, err)
		}

		bt, err := eng.Assemble(tt.code)
		if err != nil {
			t.Fatalf("assembling %q = %v", tt.code, err)
		}

		for idx, b := range bt {
			if got := fmt.Sprintf("%02x", b); got != tt.want[idx] {
				t.Errorf("bt[%d] = %02x; want = %s",
					idx, got, tt.want)
			}
		}
	}
}
