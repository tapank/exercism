package grep

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Search(pattern string, flags, files []string) []string {
	// process arguments
	var linenum, list, ignorecase, invert, matchline bool
	for _, flag := range flags {
		switch flag {
		case "-n":
			linenum = true
		case "-l":
			list = true
		case "-i":
			ignorecase = true
		case "-v":
			invert = true
		case "-x":
			matchline = true
		}
	}

	// compile regexp
	regexpflags := ""
	if ignorecase {
		regexpflags += "(?i)"
	}
	if matchline {
		pattern = "^" + pattern + "$"
	}
	reg, err := regexp.Compile(regexpflags + pattern)
	if err != nil {
		return nil
	}

	// process files
	result := []string{}
	addFileName := len(files) > 1
	for _, fname := range files {
		// set up scanner
		f, err := os.Open(fname)
		if err != nil {
			return nil
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)

		// process file
		lnum := 0
		var foundinfile bool
		for scanner.Scan() {
			lnum++
			line := scanner.Text()
			foundmatch := reg.MatchString(line)
			if invert {
				foundmatch = !foundmatch
			}
			if foundmatch {
				foundinfile = true
				if !list {
					if linenum {
						line = strconv.Itoa(lnum) + ":" + line
					}
					if addFileName {
						line = fname + ":" + line
					}
					result = append(result, line)
				}
			}
		}
		if list && foundinfile {
			result = append(result, fname)
		}
	}

	return result
}
