package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const usage = `Usage: archive_dates [<archive>]

Prints times for each file in an archive.
	zip: modification
	tar: modification, change, and access

Examples:
	archive_dates archive.zip
	archive_dates archive.tar.gz
	bzip -c -d archive.tar.bz2 | archive_dates
`

func main() {
	if len(os.Args) > 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	var r io.Reader

	if len(os.Args) == 2 {
		filename := os.Args[1]
		if strings.HasSuffix(filename, ".zip") {
			zr, err := zip.OpenReader(filename)
			try(err)
			defer zr.Close()
			for _, f := range zr.File {
				fmt.Printf("%s\t%s\n", fmtTime(f.Modified), f.Name)
			}
			return
		}

		f, err := os.Open(filename)
		try(err)
		defer f.Close()
		r = f
		if strings.HasSuffix(filename, ".gz") {
			r, err = gzip.NewReader(f)
			try(err)
		}
	} else {
		r = os.Stdin
	}

	tr := tar.NewReader(r)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		try(err)
		mt := fmtTime(h.ModTime)
		ct := fmtTime(h.ChangeTime)
		at := fmtTime(h.AccessTime)
		fmt.Printf("%s\t%s\t%s\t%s\n", mt, ct, at, h.Name)
	}
}

func fmtTime(t time.Time) string {
	t = t.UTC()
	if t == (time.Time{}) {
		return "-"
	}
	return t.String()
}

func try(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
