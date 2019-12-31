package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/aerth/modem"
)

func main() {
	var (
		outfileName = flag.String("out", "out.wav", "output file (and extension)")
		infileName  = flag.String("in", "-", "input file or '-'")
		limitSize   = flag.Int64("limit", -1, "Limit write (-1 to disable)")
		silence     = flag.Bool("s", false, "silent, no logs")
		frequency   = flag.Uint64("f", 1000, "base frequency, probably not higher than 100000")
		samplerate  = flag.Int("b", 22400, "samplerate, probably use 22000, 44100, 48000")
	)
	flag.Parse()
	log.SetFlags(0)
	if *silence {
		log.SetOutput(ioutil.Discard)
	}

	var (
		outFile *os.File
		inFile  *os.File
		err     error
	)

	if *outfileName == "-" {
		outFile = os.Stdout
	}

	if *infileName == "-" {
		inFile = os.Stdin
	}

	if outFile == nil {
		outFile, err = os.OpenFile(*outfileName, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		defer outFile.Close()
	}
	if inFile == nil {
		inFile, err = os.Open(*infileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer inFile.Close()
	}

	log.Println("reading file:", *infileName)
	log.Println("writing file:", *outfileName)

	err = modem.Modulate(modem.ModulateConfig{
		In:         inFile,
		Out:        outFile,
		Frequency:  modem.Frequency(*frequency),
		SampleRate: *samplerate,
		Limit:      *limitSize,
	})

	if err != nil {
		log.Fatalf("fatal modulation error: %v", err)
	}
	return
}
