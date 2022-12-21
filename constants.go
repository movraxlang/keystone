package keystone

// #include <keystone/keystone.h>
import "C"

const (
	KS_ERR_OK = C.KS_ERR_OK
)

type KSArch int

const (
	KS_ARCH_X86 KSArch = C.KS_ARCH_X86

	KS_ARCH_ARM   KSArch = C.KS_ARCH_ARM
	KS_ARCH_ARM64 KSArch = C.KS_ARCH_ARM64
)

type KSMode int

const (
	KS_MODE_16 KSMode = C.KS_MODE_16
	KS_MODE_32 KSMode = C.KS_MODE_32
	KS_MODE_64 KSMode = C.KS_MODE_64

	KS_MODE_LITTLE_ENDIAN KSMode = C.KS_MODE_LITTLE_ENDIAN
	KS_MODE_BIG_ENDIAN    KSMode = C.KS_MODE_BIG_ENDIAN

	KS_MODE_ARM   KSMode = C.KS_MODE_ARM
	KS_MODE_THUMB KSMode = C.KS_MODE_THUMB
)
