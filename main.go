package main

import "bytes"
import "fmt"
import "io/ioutil"

// import "gopkg.in/yaml.v2"
import "log"
import "os"
import "os/exec"
import "strings"

func readGitConfig(key string) string {
	cmd := exec.Command("git", "config", "--global", key)

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

func getPairsFileBytes() []byte {
	ioutil.ReadFile(getPairsFilePath())

	return nil
}

func getPairsFilePath() string {
	var home = os.Getenv("HOME")
	if _, err := os.Stat("./.pairs"); err == nil {
		return ".pairs"
	} else if _, err := os.Stat(home + "/.pairs"); err == nil {
		return home + "/.pairs"
	}

	log.Fatal("No .pairs file found! Put one in your home directory or current working directory!")

	return ""
}

func main() {
	path := getPairsFilePath()
	fmt.Println(path)
	author, committer := getCurrentPairs()
	fmt.Println(author)
	fmt.Println(committer)
}
