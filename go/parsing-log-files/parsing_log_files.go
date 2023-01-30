package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`".*(?i)password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +(\w+)`)
	logs := []string{}
	for _, line := range lines {
		if re.MatchString(line) {
			line = "[USR] " + re.FindStringSubmatch(line)[1] + " " + line
		}
		logs = append(logs, line)
	}
	return logs
}
