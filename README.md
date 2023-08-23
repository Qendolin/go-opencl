# OpenCL bindings for Go

Documentation at <http://godoc.org/github.com/jgillich/go-opencl/cl>.

See the [test](cl/cl_test.go) for usage examples.

By default, the OpenCL 1.2 API is exported. To get OpenCL 1.0, set the build tag `cl10`.

# Modifications of this fork (Qendolin)

Changed `data []byte` to a `length int, data unsafe.Pointer` for `CreateImage`, `CreateImageSimple`, `CreateBuffer`, `EnqueueReadImage` and `EnqueueWriteImage`.
`unsafe.Pointer` allows for any data type and is more useful than forcing a `[]byte`.
Because of this change `CreateBufferFloat32` has been removed.

Updated `ByteSlice` to use the new `unsafe.Slice`.