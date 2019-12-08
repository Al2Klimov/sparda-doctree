package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var filePattern = regexp.MustCompile(`\A(\w+?)_(.+)_Nr_(\d+)_(\d+-\d+-\d+).pdf\z`)

func main() {
	fmt.Fprintln(os.Stderr, "ls .")

	entries, errRD := ioutil.ReadDir(".")
	if errRD != nil {
		fmt.Fprintln(os.Stderr, errRD.Error())
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			if match := filePattern.FindStringSubmatch(entry.Name()); match != nil {
				dir := path.Join(match[1], strings.Replace(match[2], "_", " ", -1), match[4])
				fmt.Fprintf(os.Stderr, "mkdir -p %#v\n", dir)

				if errMA := os.MkdirAll(dir, 0777); errMA != nil {
					fmt.Fprintln(os.Stderr, errMA.Error())
					return
				}

				dest := path.Join(dir, fmt.Sprintf("Nr %s.pdf", match[3]))
				fmt.Fprintf(os.Stderr, "mv %#v %#v\n", entry.Name(), dest)

				if errRn := os.Rename(entry.Name(), dest); errRn != nil {
					fmt.Fprintln(os.Stderr, errRn.Error())
					return
				}
			}
		}
	}
}
