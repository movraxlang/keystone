package keystone

import (
	"fmt"
	"testing"
)

func TestNewEngine(t *testing.T) {
	_, err := NewEngine(KS_ARCH_ARM, KS_MODE_ARM)
	if err != nil {
		t.Fatalf("error creating engine: %v", err)
	}
}

func TestAssemble(t *testing.T) {
	tests := []struct {
		code string
		arch KSArch
		mode KSMode
		want []string
	}{
		{"add eax, ecx", KS_ARCH_X86, KS_MODE_16,
			[]string{
				"66", "01", "c8"},
		},
		{"sub r1, r2, r5", KS_ARCH_ARM, KS_MODE_ARM,
			[]string{
				"05", "10", "42", "e0"},
		},
	}

	for _, tt := range tests {
		eng, err := NewEngine(tt.arch, tt.mode)
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
