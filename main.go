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
var RECLAIMSTRING = regexp.MustCompile("^([\\s\\S]*?)(!STR!([0-9]+)!STR!)([\\s\\S]*)$") //Has s/e anchors
var ESCAPEDOLLARS = regexp.MustCompile("\\$")
var REPLACE020 = regexp.MustCompile("\020")

//Moonshine
var COMMENT = regexp.MustCompile("--((?:[^\\n])*)\\n")
var BLOCKCOMMENT = regexp.MustCompile("---((?:.|\\n)*?)---")
var ISCONDITION = regexp.MustCompile("\\sis\\s")
var ISNTCONDITION = regexp.MustCompile("\\sisnt\\s")
var FUNCACCESSOR = regexp.MustCompile("::")
var ZEROOP = regexp.MustCompile("((?:[a-zA-Z_]+(?:[a-zA-Z0-9_]*))|(?:\\([^)(\\n]+\\)))\\?")
var EXISACCESSOR = regexp.MustCompile("([a-zA-Z_]+(?:[a-zA-Z0-9_.?]*))\\?\\.")
var INCACCESSOR = regexp.MustCompile("([^\\s\\n]*)\\[(.*)##(.*)\\]")
var PINCACCESSOR = regexp.MustCompile("([^\\s\\n]*)\\[\\+\\]")
var WHITESPACEREMOVE = regexp.MustCompile("(?m)(?:\\n|\\r)+(?: |\\t)*(::|\\.|\\\\)")
var INCREMENTOPERATOR = regexp.MustCompile("\\+\\+")
var DOUBLESTATEMENTOP = regexp.MustCompile("([\t ]*)(.*)\\s&&\\s")

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
		fmt.Print("Shining: ", path + "")
		err := compile(path)
		if err != nil {return err}
		fmt.Print(" DONE \n")
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
	sArr := make([]string, 0, 0)

	isString := false
	stringOpener := ""
	hasEscape := false
	hasInterp := false
	lastCh := ""
	strBuff := ""
	locBuff := ""
	for _, c := range input {
		char, _ := strconv.Unquote(strconv.QuoteRuneToASCII(c))
    if isString {
			//Switch when string opened
			switch (char) {
				case "{":
					if lastCh == "#" {
						hasInterp = true
					}
					strBuff += char
				case "}":
					hasInterp = false
					strBuff += char
				case "\\": //escape symbol, toggle escaping
					hasEscape = !hasEscape
					strBuff += char
				case "'", "\"": //close string
					if (!hasEscape) && (!hasInterp) && stringOpener == char{
						isString = false
						if hasEscape {hasEscape = false}
						locBuff += char + "!STR!" + strconv.Itoa(len(sArr)) + "!STR!" + char
						sArr = append(sArr, strBuff)
						strBuff = "" //reset string buffer
					} else {
						strBuff += char
					}

					if (hasEscape) {
						if hasEscape {hasEscape = false}
					}

				default:
					strBuff += char
					if hasEscape {hasEscape = false}
			}

		} else {
			//Switch when string closed
			switch (char) {
				case "'", "\"": //open string
					if !hasEscape {
						isString = true
						stringOpener = char
					}
				default: //add char to nonstring buffer
					locBuff += char
			}

		}

		//record last character (for interpolation check)
		lastCh = char
	}

	return locBuff, sArr
}

//replaces all string tokens with their original value from an array
func showStrings(input string, sArr []string) (string, error) {
	local := input
	for {
		if RECLAIMSTRING.MatchString(local) == false {break}
		id, err := strconv.Atoi(RECLAIMSTRING.ReplaceAllString(local, "$3"))
		if err != nil {return "", err}
		get := "$1\020" + ESCAPEDOLLARS.ReplaceAllString(sArr[id],"$\020") + "$4"
		local = RECLAIMSTRING.ReplaceAllString(local, get) //the $1 doesnt like to be next to a string, so we put the space char code in
		local = REPLACE020.ReplaceAllString(local, "")
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

	//Remove breaking whitespace
	local = WHITESPACEREMOVE.ReplaceAllString(local, "$1\020$2\020")
	local = REPLACE020.ReplaceAllString(local, "")

	//Change && to \n
	local = DOUBLESTATEMENTOP.ReplaceAllString(local, "$1\020$2\020\n$1\020")
	local = REPLACE020.ReplaceAllString(local, "")

	//Change " is " to " == "
	local = ISCONDITION.ReplaceAllString(local, " == ")

	//Change " isnt " to " != "
	local = ISNTCONDITION.ReplaceAllString(local, " != ")

	//Change "::" to "\"
	local = FUNCACCESSOR.ReplaceAllString(local, "\\")

	//Change a[.##.] to a[.#a.]
	local = INCACCESSOR.ReplaceAllString(local, "$1\020[$2\020#$1\020$3\020]")
	local = REPLACE020.ReplaceAllString(local, "")

	//Change a[+] to a[#a+1]
	local = PINCACCESSOR.ReplaceAllString(local, "$1\020[#$1\020+1]")
	local = REPLACE020.ReplaceAllString(local, "")

	//Change ++ to += 1
	local = INCREMENTOPERATOR.ReplaceAllString(local, "+=1")

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
