package main

import "bytes"
import "fmt"

// import "io/ioutil"

// import "gopkg.in/yaml.v2"
import "log"
import "os"
import "os/exec"
import "strings"
import "errors"

type cmd interface {
	Run() error
}

var command func(string, ...string) cmd = func(name string, args ...string) cmd {
	return exec.Command(name, args...)
}

func readGitConfig(key string) string {
	cmd := command("git", "config", "--global", key)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("key missing from git config: %s", key))
	}

	result := strings.TrimSpace(out.String())
	fmt.Println(result)
	if result == "" {
		log.Fatal(fmt.Sprintf("key missing from git config: %s", key))
	}

	return result
}

func nameForInitials(initials string) string {
	return ""
}

func getCurrentPairs() (string, string) {
	author := readGitConfig("dbl.author")
	committer := readGitConfig("dbl.committer")
	return author, committer
}

// func getPairsFileBytes() []byte {
// 	ioutil.ReadFile(getPairsFilePath())
//
// 	return nil
// }

var getenv func(string) string = os.Getenv
var stat func(string) (os.FileInfo, error) = os.Stat

func getPairsFilePath() (string, error) {
	var home = getenv("HOME")
	if _, err := stat(".pairs"); err == nil {
		return ".pairs", nil
	} else if _, err := stat(home + "/.pairs"); err == nil {
		return home + "/.pairs", nil
	}

	return "", errors.New("No .pairs file found! Put one in your home directory or current working directory!")
}

func main() {
	path, _ := getPairsFilePath()
	fmt.Println(path)
	author, committer := getCurrentPairs()
	fmt.Println(author)
	fmt.Println(committer)

	fmt.Println(nameForInitials(author))
}
