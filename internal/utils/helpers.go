package utils

import (
	"bufio"
	"os/exec"
	"strings"
)

// readLine читает одну строку из STDIN
func ReadLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// parseCommaSeparatedList очищает список, введенный через запятую
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

// getCurrentGitBranch (бонус) пытается получить имя текущей ветки git
func GetCurrentGitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "" // Не в git репозитории или git не найден
	}
	return strings.TrimSpace(string(output))
}
