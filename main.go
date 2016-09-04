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
var STRING1 = regexp.MustCompile("^(.*?)(\")(?:\\\\[\"]|[^\"])*(\")(.*)$")
var STRING2 = regexp.MustCompile("^(.*?)(')(?:\\\\[']|[^'])*(')(.*)$")

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
		fmt.Println("compiling", path + "...")
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
	translated := translate(contents)
	err = ioutil.WriteFile(FILETYPE.ReplaceAllString(path,".moon"), []byte(translated), 0644)
	if err != nil {fmt.Println(err)}
	return nil
}

func translate(input string) string {
	local := input

	//Hide Strings
	sArr := make([]string, 0, 0)
	//Hide " strings
	for {
		if STRING1.MatchString(local) == false {break}
		found := STRING1.FindString(local)
		local = STRING1.ReplaceAllString(local, "!STR!" + strconv.Itoa(len(sArr)) + "!STR!")
		fmt.Println(found)
		sArr = append(sArr, found)
	}

	//Hide ' strings
	for {
		if STRING2.MatchString(local) == false {break}
		found := STRING2.FindString(local)
		local = STRING2.ReplaceAllString(local, "!STR!" + strconv.Itoa(len(sArr)) + "!STR!")
		sArr = append(sArr, found)
	}

	//Delete multiline's
	local = BLOCKCOMMENT.ReplaceAllString(local, "")

	return local
}
