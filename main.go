package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/wellington/go-libsass"
)

func main() {
	if os.Args[1] == "" {
		log.Fatalln("Usage: gosass filename.sass")
		os.Exit(1)
	}
	sass, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	bufsass := bytes.NewBuffer(sass)
	bufscss := bytes.NewBuffer([]byte{})
	err = libsass.ToScss(bufsass, bufscss)

	comp, err := libsass.New(os.Stdout, bufscss)
	if err != nil {
		log.Fatal(err)

	}

	if err := comp.Run(); err != nil {
		log.Fatal(err)

	}

}
