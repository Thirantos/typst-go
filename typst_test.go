package typst

import "testing"
import "github.com/h2non/filetype"

func TestGenerateSimplePdf(t *testing.T) {

	w := NewTypstWorld("./", `
	== Hello World!
	` /* lang:typst */)
	defer w.Free()

	c, e := w.Compile()
	defer c.Free()
	if e != nil {
		t.Error(e)
	}
	pdf, e := c.ToPdf()
	if e != nil {
		t.Error(e)
	}

	ft, e := filetype.Get(pdf)
	wanted := filetype.GetType("pdf")
	if e != nil {
		t.Error(e)
	}

	if ft != wanted {
		t.Errorf("Expected %v, got %v", wanted.Extension, ft.Extension)
	}

}

func TestGenerateSimpleInvalid(t *testing.T) {

	w := NewTypstWorld("./", `
	#this_it_not_compilable_typst
	` /* lang:typst */)
	defer w.Free()

	c, e := w.Compile()
	defer c.Free()
	if e == nil {
		t.Error("Expected error. (unknown_variable)")
	}
}

func TestGenerateImportPdf(t *testing.T) {

	w := NewTypstWorld("./", `
	#import "@preview/cetz:0.5.0"

	` /* lang:typst */)
	defer w.Free()

	c, e := w.Compile()
	defer c.Free()
	if e != nil {
		t.Error(e)
	}
	pdf, e := c.ToPdf()
	if e != nil {
		t.Error(e)
	}

	ft, e := filetype.Get(pdf)
	wanted := filetype.GetType("pdf")
	if e != nil {
		t.Error(e)
	}

	if ft != wanted {
		t.Errorf("Expected %v, got %v", wanted.Extension, ft.Extension)
	}

}
