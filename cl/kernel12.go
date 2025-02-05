//go:build !cl10

package cl

// #include "cl.h"
import "C"
import "unsafe"

func (k *Kernel) ArgName(index int) (string, error) {
	var strC [1024]byte
	var strN C.size_t
	if err := C.clGetKernelArgInfo(k.clKernel, C.cl_uint(index), C.CL_KERNEL_ARG_NAME, 1024, unsafe.Pointer(&strC[0]), &strN); err != C.CL_SUCCESS {
		return "", toError(err)
	}
	return C.GoString((*C.char)(unsafe.Pointer(&strC))), nil
}
