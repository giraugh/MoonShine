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

//Compiler
var FILETYPE = regexp.MustCompile("(\\.shine|\\.mshine)")
var STRING1 = regexp.MustCompile("^([\\s\\S]*?[^\\\\])(?:\")((?:\\\\[\"]|[^\"])*)(?:\")([\\s\\S]*)$") //Has s/e anchors
var STRING2 = regexp.MustCompile("^([\\s\\S]*?[^\\\\])(?:')((?:\\\\[']|[^'])*)(?:')([\\s\\S]*)$") //Has s/e anchors
var RECLAIMSTRING = regexp.MustCompile("^([\\s\\S]*?)(!STR!([0-9]+)!STR!)([\\s\\S]*)$") //Has s/e anchors

//Moonshine
var COMMENT = regexp.MustCompile("--((?:[^\\n])*)\\n")
var BLOCKCOMMENT = regexp.MustCompile("---((?:.|\\n)*?)---")
var ISCONDITION = regexp.MustCompile("\\sis\\s")
var ISNTCONDITION = regexp.MustCompile("\\sisnt\\s")
var FUNCACCESSOR = regexp.MustCompile("::")
var ZEROOP = regexp.MustCompile("((?:[a-zA-Z_]+(?:[a-zA-Z0-9_]*))|(?:\\([^)(\\n]+\\)))\\?")
var EXISACCESSOR = regexp.MustCompile("([a-zA-Z_]+(?:[a-zA-Z0-9_.?]*))\\?\\.")

//Get args and walk filepath
func main() {

	//get the file/directory from the argument
	if len(os.Args) < 2 {
		fmt.Println("### Incorrect Usage: Didnt supply file or directory to compile. ###\n\n\t\tUsage: MoonShine <file or directory>")
		return
	}
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Println("Runtime Error:", err)
		return
	}
}

//Called for each compiling file, calls translate
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

//Reads file, translates, writes file
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

//writes all strings to array and replaces them with a token
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

//replaces all string tokens with their original value from an array
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

//translates moonshine to moonscript (where the magic happens)
func translate(input string) (string, error) {
	local := input

	//Delete comments
	local = BLOCKCOMMENT.ReplaceAllString(local, "")
	local = COMMENT.ReplaceAllString(local, "\n")

	//Hide Strings
	local, sArr := hideStrings(local)

	//Change " is " to " == "
	local = ISCONDITION.ReplaceAllString(local, " == ")

	//Change " isnt " to " != "
	local = ISNTCONDITION.ReplaceAllString(local, " != ")

	//Change "::" to "\"
	local = FUNCACCESSOR.ReplaceAllString(local, "\\")

	//Existential accessor, must come before zero op
	for {
		if EXISACCESSOR.MatchString(local) == false {break}
		local = EXISACCESSOR.ReplaceAllString(local, "($1 or {}).")
	}

	//Zero operator
	local = ZEROOP.ReplaceAllString(local, "($1 != \"\" and $1 != 0)")

	//Show Strings
	local, err := showStrings(local, sArr)
	if err != nil {return "", err}

	return local, nil
}
