package cl

// #include "cl.h"
import "C"
import (
	"github.com/Qendolin/go-opencl/cl"
)

func (ctx *Context) CreateSampler(normalizedCoords bool, addressingMode cl.SamplerAddressingMode, filterMode cl.SamplerFilterMode) (*MemObject, error) {
	var err C.cl_int
	clSampler := C.clCreateSampler(ctx.clContext, clBool(normalizedCoords), addressingMode, filterMode, &err)
	if err != C.CL_SUCCESS {
		return nil, toError(err)
	}
	if clSampler == nil {
		return nil, ErrUnknown
	}
	return newSampler(clSampler), nil
}
