package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/richardlehane/unwarc"
	"github.com/richardlehane/webarchive"
)

var (
	target = flag.String("d", "", "target directory e.g. -d dump")
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("Expecting name of one or more WARC or ARC files to unpack e.g. unwarc blackbooks.warc.gz")
	}

	var rdr webarchive.Reader

	for i, v := range flag.Args() {
		var dir string
		if *target != "" {
			dir = *target
		} else {
			dir = unwarc.Base(v)
		}
		f, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			rdr, err = webarchive.NewReader(f)
		} else {
			err = rdr.Reset(f)
		}
		if err != nil {
			log.Fatal(err)
		}
		for record, err := rdr.NextPayload(); err == nil; record, err = rdr.NextPayload() {
			rel, fn := unwarc.Sanitise(record.URL())
			if rel == "" {
				rel = dir
			} else {
				rel = filepath.Join(dir, rel)
			}
			if err := os.MkdirAll(rel, 0666); err != nil {
				log.Fatal(err)
			}
			out, err := os.Create(filepath.Join(rel, fn))
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(out, record)
			if err != nil {
				log.Fatal(err)
			}
			out.Close()
		}
		rdr.Close()
		f.Close()
	}
}
