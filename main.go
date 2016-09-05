package main

import (
	"path/filepath"
	"io/ioutil"
	"os"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

/*
Need to make string regex's only match first match using anchors.

this works (kinda)
^(.*?)((\")(?:\\\\[\"]|[^\"])*(\"))(.*)$


*/


//Compiler
var FILETYPE = regexp.MustCompile("(\\.shine|\\.mshine)")
var STRING1 = regexp.MustCompile("^([\\s\\S]*?)(?:\")((?:\\\\[\"]|[^\"])*)(?:\")([\\s\\S]*)$") //Has s/e anchors
var STRING2 = regexp.MustCompile("^([\\s\\S]*?)(?:')((?:\\\\[']|[^'])*)(?:')([\\s\\S]*)$") //Has s/e anchors
var RECLAIMSTRING = regexp.MustCompile("^([\\s\\S]*?)(!STR!([0-9]+)!STR!)([\\s\\S]*)$") //Has s/e anchors

//Moonshine
var BLOCKCOMMENT = regexp.MustCompile("---((?:.|\\n)*)---")

func main() {

	//get the file/directory from the argument
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Println("Runtime Error:", err)
		return
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {return err}

	//Does the string contain the file extension?
	matched:= FILETYPE.MatchString(strings.Replace(path, "\\", "\\\\", -1))

	if matched == true {
		//Compile it then
		fmt.Println("Shining:", path + "...")
		err := compile(path)
		if err != nil {return err}
		fmt.Println("done")
	}
	return nil
}

func compile(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {return err}
	contents := string(bs)
	translated, err := translate(contents)
	if err != nil {return err}
	err = ioutil.WriteFile(FILETYPE.ReplaceAllString(path,".moon"), []byte(translated), 0644)
	if err != nil {fmt.Println(err)}
	return nil
}

func translate(input string) (string, error) {
	local := input

	//Hide Strings
	local, sArr := hideStrings(local)

	//Delete multiline's
	local = BLOCKCOMMENT.ReplaceAllString(local, "")

	//Show Strings
	local, err := showStrings(local, sArr)
	if err != nil {return "", err}

	return local, nil
}

func hideStrings(input string) (string, []string) {
	local := input
	sArr := make([]string, 0, 0)
	//Hide " strings
	for {
		if STRING1.MatchString(local) == false {break}
		found := STRING1.ReplaceAllString(local, "$2")
		local = STRING1.ReplaceAllString(local, "$1!STR!" + strconv.Itoa(len(sArr)) + "!STR!$3")
		sArr = append(sArr, found)
	}

	//Hide ' strings
	for {
		if STRING2.MatchString(local) == false {break}
		found := STRING2.ReplaceAllString(local, "$2")
		local = STRING2.ReplaceAllString(local, "$1!STR!" + strconv.Itoa(len(sArr)) + "!STR!$3")
		sArr = append(sArr, found)
	}
	return local, sArr
}

func showStrings(input string, sArr []string) (string, error) {
	local := input
	for {
		if RECLAIMSTRING.MatchString(local) == false {break}
		id, err := strconv.Atoi(RECLAIMSTRING.ReplaceAllString(local, "$3"))
		if err != nil {return "", err}
		local = RECLAIMSTRING.ReplaceAllString(local, "$1\""+sArr[id]+"\"$4")
	}
	return local, nil
}
