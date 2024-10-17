package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"os"
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
	log.Println("current version is " + version)
	n, _ := strconv.Atoi(strings.Split(version, ".")[2])

	log.Println("git add .")
	cmd = exec.Command("git", "add", ".")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("git commit -m " + *msg)
	cmd = exec.Command("git", "commit", "-m", *msg)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("git tag " + " v0.0." + strconv.Itoa(n+1) + " master")
	cmd = exec.Command("git", "tag", "v0.0."+strconv.Itoa(n+1), "master")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("git push -u origin " + "v0.0." + strconv.Itoa(n+1))
	cmd = exec.Command("git", "push", "-u", "origin", "v0.0."+strconv.Itoa(n+1))
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("git push")
	cmd = exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
