package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
) 
type Task struct {
	ID     int
	Title  string
	Status string
}

const fileName = "tasks.txt"

func LoadTasks() []Task {
file, err := os.Open(fileName)
	if err != nil {
		return []Task{}
	}
	defer file.Close()

	var tasks []Task
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, ";", 3)
		if len(parts) < 3 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		task := Task{
			ID:     id,
			Status: parts[1],
			Title:  parts[2],
			}
		tasks = append(tasks, task)
	}

	return tasks
}
