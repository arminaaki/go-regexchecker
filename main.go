package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const banner = `
_______      ________     ______    ________   ____  ____                                 
|_   __ \    |_   __  |  .' ___  |  |_   __  | |_  _||_  _|                                
  | |__) |     | |_ \_| / .'   \_|    | |_ \_|   \ \  / /                                  
  |  __ /      |  _| _  | |   ____    |  _| _     > '' <                                   
 _| |  \ \_   _| |__/ | \ '.___]  |  _| |__/ |  _/ /''\ \_                                 
|____| |___| |________|  '._____.'  |________| |____||____|                                
   ______   ____  ____   ________     ______   ___  ____    ________   _______      _   _  
 .' ___  | |_   ||   _| |_   __  |  .' ___  | |_  ||_  _|  |_   __  | |_   __ \    | | | | 
/ .'   \_|   | |__| |     | |_ \_| / .'   \_|   | |_/ /      | |_ \_|   | |__) |   | | | | 
| |          |  __  |     |  _| _  | |          |  __'.      |  _| _    |  __ /    | | | | 
\ '.___.'\  _| |  | |_   _| |__/ | \  .___.'\  _| |  \ \_   _| |__/ |  _| |  \ \_  |_| |_| 
 '.____ .' |____||____| |________|  '.____ .' |____||____| |________| |____| |___| (_) (_) 
==========================================================================================
`

var (
	Match    = Green
	NonMatch = Red
	Captures = White
	Banner   = Purple
	Green    = color("\033[1;32m%s\033[0m")
	Red      = color("\033[1;31m%s\033[0m")
	White    = color("\033[1;37m%s\033[0m")
	Purple   = color("\033[1;34m%s\033[0m")
)

// Checker struct
type Checker struct {
	Rule *regexp.Regexp
	file *os.File
}

// Run function
func (c *Checker) Run() error {

	fmt.Println(Banner(banner))
	defer c.file.Close()

	scanner := bufio.NewScanner(c.file)
	for scanner.Scan() {
		text := scanner.Text()
		isMatch, captures := c.Validate(text)
		if isMatch {
			fmt.Println(Match(fmt.Sprintf("✔ %s", text)))
			fmt.Printf("%s\n", Captures(formatMap(captures)))
		} else {
			fmt.Println(NonMatch(fmt.Sprintf("✘ %s", text)))
		}

		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// Validate function
func (c *Checker) Validate(str string) (bool, map[string]string) {
	matches := c.Rule.FindStringSubmatch(str)
	c.Rule.FindStringSubmatch(str)

	if len(matches) == 0 {
		return false, nil
	}

	if len(matches) == 1 {
		return true, nil
	}

	expNames := c.Rule.SubexpNames()
	result := make(map[string]string)
	for i := 1; i < len(matches); i++ {

		groupName := expNames[i]
		if groupName == "" {
			result[fmt.Sprintf("%d.", i)] = matches[i]
		} else {
			result[groupName] = matches[i]
		}
	}
	return true, result
}

func validateFile(filename string) (*os.File, error) {
	if filename == "" {
		return os.Stdin, nil
	}

	return os.Open(filename)
}

func validatePattern(str string) (*regexp.Regexp, error) {

	// Remove forward slashes if present
	r := regexp.MustCompile(`^/(?P<RegxBody>.*)/$`)
	if result := r.FindStringSubmatch(str); len(result) > 1 {
		str = result[1]
	}

	// Convert to Golang capture patterns
	m := regexp.MustCompile(`(\?<(\w+)>)`)
	str = m.ReplaceAllString(str, "?P<$2>")

	return regexp.Compile(str)
}

// NewChecker function
func NewChecker(filename string, rule string) (*Checker, error) {
	file, err := validateFile(filename)
	if err != nil {
		return nil, err
	}

	reRule, err := validatePattern(rule)
	if err != nil {
		return nil, fmt.Errorf("validatePattern: %s", err)
	}
	return &Checker{
		file: file,
		Rule: reRule,
	}, nil
}

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func formatMap(m map[string]string) string {
	var result []string
	if len(m) == 0 {
		return ""
	}

	maxLength := 0
	for k:= range m {
		if len(k) > maxLength {
			maxLength = len(k)
		}
	}
	paddedKeyFormat := fmt.Sprintf("%%-%ds => %%s", maxLength)
	for k, v := range m {
		result = append(result, fmt.Sprintf(paddedKeyFormat, k, v))
	}

	return strings.Join(result, "\n")
}

func main() {

	var filename = flag.String("filename", "", "The input filename. Default: STDIN")
	var rule = flag.String("rule", "", "The Regex rule")
	flag.Parse()

	checker, err := NewChecker(*filename, *rule)
	if err != nil {
		log.Fatal(err)
	}
	if err := checker.Run(); err != nil {
		log.Fatal(err)
	}
}
