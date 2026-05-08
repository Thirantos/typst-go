package typst

/*
#cgo CFLAGS: -I .
#cgo darwin,arm64 LDFLAGS: -L ./lib/aarch64-apple-darwin/ -ltypst -framework Security -framework CoreFoundation
#cgo darwin,amd64 LDFLAGS: -L ./lib/x86_64-apple-darwin/ -ltypst -framework Security -framework CoreFoundation

#cgo linux,arm64 LDFLAGS: -L ./lib/aarch64-unknown-linux-gnu/ -ltypst
#cgo linux,amd64 LDFLAGS: -L ./lib/x86_64-unknown-linux-gnu/ -ltypst

#cgo windows,amd64 LDFLAGS: -L ./lib/x86_64-pc-windows-msvc/ -ltypst

#include <stdlib.h>
#include "typst.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type TypstWorld struct{ world *C.TypstWorld }
type TypstDocument struct{ doc *C.TypstDocument }

// generates a new TypstWorld Object
// Must be manually freed using `.Free()`
func NewTypstWorld(path string, source string) TypstWorld {
	path_c := C.CString(path)
	defer C.free(unsafe.Pointer(path_c))

	source_c := C.CString(source)
	defer C.free(unsafe.Pointer(source_c))

	_world := C.typst_world_new(path_c, source_c)

	return TypstWorld{_world}
}

// Frees a TypstWorld Object
func (world *TypstWorld) Free() {
	C.typst_world_free(world.world)
}

// Compiles a TypstWorld Object
// Must be manually freed using `.Free()`
func (world *TypstWorld) Compile() (TypstDocument, error) {
	var err C.TypstError
	doc := C.typst_world_compile(world.world, &err)
	err_msg := C.GoString((err.message))

	if err_msg == "" {

		return TypstDocument{doc}, nil
	}
	return TypstDocument{doc}, fmt.Errorf("Typst Error: %v", err_msg)

}

// Frees a TypstDocument Object
func (doc *TypstDocument) Free() {
	C.typst_document_free(doc.doc)
}
func (doc *TypstDocument) ToPdf() ([]byte, error) {

	var data *C.uint8_t
	var length C.uintptr_t

	var err C.TypstError
	C.typst_document_to_pdf(doc.doc, &length, &data, &err)
	defer C.typst_pdf_free(data, length)
	bytes := C.GoBytes(unsafe.Pointer(data), C.int(length))

	err_msg := C.GoString((err.message))

	if err_msg == "" {
		return bytes, nil
	}

	return bytes, fmt.Errorf("Typst Error: %v", err_msg)
}
