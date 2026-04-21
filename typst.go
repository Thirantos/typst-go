package typst

/*
#cgo CFLAGS: -I .
#cgo darwin,arm64 LDFLAGS: -L ./lib/aarch64-apple-darwin/ -ltypst -framework Security -framework CoreFoundation
#cgo darwin,amd64 LDFLAGS: -L ./lib/x86_64-apple-darwin/ -ltypst -framework Security -framework CoreFoundation

#cgo linux,amd64 LDFLAGS: -L ./lib/x86_64-unknown-linux-gnu/
#cgo windows,amd64 LDFLAGS: -L ./lib/x86_64-pc-windows-msvc/

#include <stdlib.h>
#include <string.h>
#include "typst.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type TypstWorld struct{ w *C.TypstWorld }
type TypstDocument struct{ d *C.TypstDocument }

// NewTypstWorld Generates a new Typst World
// NOTICE: must be maunally freed using `.Free()`
func NewTypstWorld(path string, source string) TypstWorld {
	path_c := C.CString(path)
	defer C.free(unsafe.Pointer(path_c))

	source_c := C.CString(source)
	defer C.free(unsafe.Pointer(source_c))

	_world := C.typst_world_new(path_c, source_c)

	return TypstWorld{_world}
}

func (w *TypstWorld) Free() {
	C.typst_world_free(w.w)
}

func (w *TypstWorld) Compile() (TypstDocument, error) {
	var e C.TypstError
	d := C.typst_world_compile(w.w, &e)
	err_msg := C.GoString((e.message))

	if err_msg == "" {

		return TypstDocument{d}, nil
	}
	return TypstDocument{d}, fmt.Errorf("Typst Error: %v", err_msg)

}

func (d *TypstDocument) Free() {
	C.typst_document_free(d.d)
}
func (d *TypstDocument) ToPdf() ([]byte, error) {

	var data *C.uint8_t
	var leng C.uintptr_t

	var e C.TypstError
	C.typst_document_to_pdf(d.d, &leng, &data, &e)
	defer C.typst_pdf_free(data, leng)
	b := C.GoBytes(unsafe.Pointer(data), C.int(leng))

	err_msg := C.GoString((e.message))

	if err_msg == "" {
		return b, nil
	}

	return b, fmt.Errorf("Typst Error: %v", err_msg)
}
