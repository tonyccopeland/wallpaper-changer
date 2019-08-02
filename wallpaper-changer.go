// Tony Copeland <tonyccopeland@gmail.com>
//
// Simple Desktop Wallpaper Changer for my Gnome based Linux Desktop every 10 minutes
//  (really just a small app for me to tinker with the go language)
//
// Build
// rm -f ~/go/bin/wallpaper-changer;go build -o ~/go/bin/wallpaper-changer wallpaper-changer.go
// Install - add the following to .bashrc
// pkill -f wallpaper-changer; sleep 1;~/go/bin/wallpaper-changer <directory with pictures>/dev/null 2>&1 &
package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// init a slice of type string
	var fs []string

	// read the pcitures directory
	files, ferr := ioutil.ReadDir(os.Args[1])
	if ferr != nil {
		log.Fatal(ferr)
	}

	// append list of file names to slice
	for _, f := range files {
		// limit my file types to pictures
		if strings.HasSuffix(f.Name(), ".jpg") {
			//fmt.Println(f.Name())
			fs = append(fs, f.Name())
		}
	}

	// randomize the seed once
	rand.Seed(time.Now().UnixNano())

	// loop forever
	for {
		// shuffle the slices
		rand.Shuffle(len(fs), func(i, j int) {
			fs[i], fs[j] = fs[j], fs[i]
		})

		// run gsettings to change desktop
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", os.Args[1]+fs[0])
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)

		}

		// sleep in between background picture changes
		scmd := exec.Command("sleep", "10m")
		serr := scmd.Run()
		if serr != nil {
			log.Fatal(serr)
		}
	}
}
