package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func expandField(field string, min int, max int) []int {
	var result []int

	if field == "*" {
		for i := min; i <= max; i++ {
			result = append(result, i)
		}
		return result
	}

	parts := strings.Split(field, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			num, _ := strconv.Atoi(part)
			result = append(result, num)
		}
	}

	return result
}

func parseCronExpression(expression string) map[string][]int {
	fields := strings.Fields(expression)
	if len(fields) != 6 {
		fmt.Println("Invalid cron expression. It should have 6 fields.")
		os.Exit(1)
	}

	minute := expandField(fields[0], 0, 59)
	hour := expandField(fields[1], 0, 23)
	dayOfMonth := expandField(fields[2], 1, 31)
	month := expandField(fields[3], 1, 12)
	dayOfWeek := expandField(fields[4], 0, 6)

	return map[string][]int{
		"minute":      minute,
		"hour":        hour,
		"day of month": dayOfMonth,
		"month":       month,
		"day of week": dayOfWeek,
	}
}

func formatOutput(parsedData map[string][]int) {
	for key, value := range parsedData {
		fmt.Printf("%-14s %s\n", key, formatList(value))
	}
}

func formatList(list []int) string {
	var strList []string
	for _, num := range list {
		strList = append(strList, strconv.Itoa(num))
	}
	return strings.Join(strList, " ")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron_parser_go \"*/15 0 1,15 * 1-5 /usr/bin/find\"")
		os.Exit(1)
	}

	cronExpression := os.Args[1]
	parsedData := parseCronExpression(cronExpression)
	formatOutput(parsedData)
}
