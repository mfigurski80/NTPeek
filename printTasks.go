package main

import (
	"fmt"
	"math"
	"time"
)

func formatRelativeDate(date string) string {
	// convert "2022-07-01" to "next Monday"
	t, _ := time.Parse("2006-01-02", date[:10])
	hoursDiff := time.Until(t).Hours()
	diff := math.Ceil((hoursDiff) / 24)
	// fmt.Printf("%v -> %v", hoursDiff/24, diff)
	if diff == 0 {
		return "Today"
	}
	if diff == 1 {
		return "Tomorrow"
	}
	if diff == -1 {
		return "Yesterday"
	}
	if diff < 10 && diff > -7 {
		if diff >= 8 {
			return "next " + t.Weekday().String()
		}
		if diff > 0 {
			return t.Weekday().String()
		}
		return "last " + t.Weekday().String()
	}
	if diff > 0 {
		return "in " + fmt.Sprint(diff) + " days"
	}
	return fmt.Sprint(-diff) + " days ago"
}

var colorReset = "\033[0m"
var bold = "\033[1m"
var weak = "\033[2m"
var colorMap = map[string]string{
	"pink":    "\033[48;5;217;30m",
	"red":     "\033[48;5;203;30m",
	"orange":  "\033[48;5;208;30m",
	"yellow":  "\033[48;5;229;30m",
	"green":   "\033[48;5;120;30m",
	"blue":    "\033[48;5;45;30m",
	"purple":  "\033[48;5;147;30m",
	"brown":   "\033[48;5;101m",
	"gray":    "\033[48;5;250;30m",
	"default": "\033[48;5;240m",
}

func printTasks(tasks []Task) {
	maxClassLen := 0
	classLengths := make([]int, len(tasks))
	for i, task := range tasks {
		if len(task.Class) > maxClassLen {
			maxClassLen = len(task.Class)
		}
		classLengths[i] = len(task.Class)
	}

	for i, task := range tasks {
		classOffset := maxClassLen - classLengths[i]
		class := colorMap[task.ClassColor] + task.Class + colorReset
		name := bold + task.Name + colorReset
		due := formatRelativeDate(task.Due)
		// fmt.Printf("%s %s %s\n", class, name, due)
		fmt.Printf("%*s%s | %s  %s(%s)%s\n", classOffset, "", class, name, weak, due, colorReset)
	}
}
