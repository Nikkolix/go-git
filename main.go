package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var msg = flag.String("msg", "msg", "commit message")

// git add .
// git commit -m "{msg}"
// git tag v0.0.{n} master
// git push -u origin v0.0.{n}
// git push

func main() {
	cmd := exec.Command("git", "describe", "--tags")

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	cmd.Stdout = writer

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	version := b.String()
	version = strings.Trim(version, "\r\n")
	if !regexp.MustCompile(`^v\d+\.\d+\.\d+$`).MatchString(version) {
		log.Fatal("version not matched: \"" + version + "\"")
	}

	n, _ := strconv.Atoi(strings.Split(version, ".")[2])

	err = exec.Command("git", "add", ".").Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("git", "commit", "-m", *msg).Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("git", "tag", "v0.0."+strconv.Itoa(n+1), "master").Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("git", "push", "-u", "origin", "v0.0."+strconv.Itoa(n+1)).Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("git", "push").Run()
	if err != nil {
		log.Fatal(err)
	}
}
