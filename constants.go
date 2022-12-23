package keystone

// #include <keystone/keystone.h>
import "C"

const (
	KS_ERR_OK = C.KS_ERR_OK
)

type KSArch int

const (
	KS_ARCH_ARM64 KSArch = C.KS_ARCH_ARM64
)

type KSMode int

const (
	KS_MODE_LITTLE_ENDIAN KSMode = C.KS_MODE_LITTLE_ENDIAN
)
