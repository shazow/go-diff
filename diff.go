package diff

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

type Differ interface {
	// Diff writes a generated patch to out for the diff between a and b.
	Diff(out io.Writer, a io.Reader, b io.Reader) error
}

// Object is the minimum representation for generating diffs with git-style headers.
type Object struct {
	io.Reader
	// ID is the sha1 of the object
	ID [20]byte
	// Path is the root-relative path of the object
	Path string
	// Mode is the entry mode of the object
	Mode int
}

// ErrEmptyComparison is used when both src and dst are EmptyObjects
var ErrEmptyComparsion = errors.New("no objects to compare, both are empty")

// EmptyObject can be used when there is no corresponding src or dst entry, such as during deletions or creations.
var EmptyObject Object

// Writer writes diffs using a given Differ including git-style headers between each patch.
type Writer struct {
	io.Writer
	Differ
	SrcPrefix string
	DstPrefix string
}

// WriteHeader writes only the header for the comparison between the src and dst Objects.
func (w *Writer) WriteHeader(src, dst Object) error {
	var srcPath, dstPath string
	if src == EmptyObject && dst == EmptyObject {
		return ErrEmptyComparsion
	}
	if src != EmptyObject {
		srcPath = filepath.Join(w.SrcPrefix, src.Path)
		dstPath = srcPath
	}
	if dst != EmptyObject {
		dstPath = filepath.Join(w.DstPrefix, dst.Path)
		if srcPath == "" {
			srcPath = dstPath
		}
	}
	// TODO: Detect renames?

	fmt.Fprintf(w, "diff --git %s %s\n", srcPath, dstPath)
	if src == EmptyObject {
		fmt.Fprintf(w, "new file mode %o\n", dst.Mode)
		fmt.Fprintf(w, "index %x..%x\n", src.ID, dst.ID)
		fmt.Fprintf(w, "--- /dev/null\n")
		fmt.Fprintf(w, "+++ %s\n", dstPath)
	} else if dst == EmptyObject {
		fmt.Fprintf(w, "deleted file mode %o\n", src.Mode)
		fmt.Fprintf(w, "index %x..%x\n", src.ID, dst.ID)
		fmt.Fprintf(w, "--- %s\n", srcPath)
		fmt.Fprintf(w, "+++ /dev/null\n")
	} else {
		fmt.Fprintf(w, "index %x..%x %o\n", src.ID, dst.ID, dst.Mode)
		fmt.Fprintf(w, "--- %s\n", srcPath)
		fmt.Fprintf(w, "+++ %s\n", dstPath)
	}
	return nil
}

// WriteDiff performs a Diff between a and b and writes only the resulting diff. It does not write the header.
func (w *Writer) WriteDiff(a, b io.Reader) error {
	return w.Differ.Diff(w, a, b)
}
