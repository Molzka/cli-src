package utils

import (
	"bufio"
	"os/exec"
	"strings"
	"time"
)

func ReadLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ParseCommaSeparatedList(input string) []string {
	ids := strings.Split(input, ",")
	cleanedIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		trimmedID := strings.TrimSpace(id)
		if trimmedID != "" {
			cleanedIDs = append(cleanedIDs, trimmedID)
		}
	}
	return cleanedIDs
}

func GetCurrentGitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func FormatDate(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("02.01.2006 15:04")
}
