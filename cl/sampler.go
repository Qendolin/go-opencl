package cl

// #include "cl.h"
import "C"

func (ctx *Context) CreateSampler(normalizedCoords bool, addressingMode SamplerAddressingMode, filterMode SamplerFilterMode) (*Sampler, error) {
	var err C.cl_int
	clSampler := C.clCreateSampler(ctx.clContext, clBool(normalizedCoords), C.cl_uint(addressingMode), C.cl_uint(filterMode), &err)
	if err != C.CL_SUCCESS {
		return nil, toError(err)
	}
	if clSampler == nil {
		return nil, ErrUnknown
	}
	return newSampler(clSampler), nil
}
