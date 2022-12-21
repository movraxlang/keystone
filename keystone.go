package keystone

import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -lkeystone -lstdc++ -lm

#include <stdlib.h>
#include <keystone/keystone.h>

static unsigned char get_elem_from_array(unsigned char ** arr, int idx) {
	return arr[0][idx];
}
*/
import "C"

type Engine struct {
	ks *C.ks_engine
}

func NewEngine(arch KSArch, mode KSMode) (*Engine, error) {
	var ks *C.ks_engine
	err := C.ks_open(C.ks_arch(C.int(arch)), C.int(mode), &ks)
	if err != KS_ERR_OK {
		return nil, errors.New("error creating engine")
	}

	return &Engine{
		ks: ks,
	}, nil
}

func (e *Engine) Assemble(code string) ([]byte, error) {
	var count C.ulong
	var size C.ulong
	var encode *C.uchar
	assembly := C.CString(code)
	defer C.free(unsafe.Pointer(assembly))

	ret := C.ks_asm(e.ks, assembly, 0, &encode, &size, &count)
	if ret != 0 {
		return nil, errors.New(fmt.Sprintf("failed assembling with error code: %d", C.ks_errno(e.ks)))
	}

	bt := make([]byte, int(size))
	for i := 0; i < int(size); i++ {
		elem := C.get_elem_from_array(&encode, C.int(i))
		bt[i] = byte(elem)
	}
	return bt, nil
}
