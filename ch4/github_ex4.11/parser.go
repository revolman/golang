// Упражнение 4.11 - часть программы для создания Issues
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Edit - вызов эдитора
func Edit(e string, fpath string) {
	cmd := exec.Command(e, fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Ошибка: %v\n", err)
	}
	fmt.Println("Успешно отредактировано.")
}

// ParseFile сканирует врменный текстовый файл, затем удаляет его.
// Создаёт отображение, в котором записаны Title и Body Issue.
// Возвращает указатель на отображение.
func ParseFile(fpath string) map[string]string {
	// fpath := TMPFile()
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make(map[string]string)
	var body string
	afterBody := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Title: ") {
			line = strings.TrimPrefix(line, "Title: ")
			data["title"] += line
			continue
		}

		if strings.HasPrefix(line, "State: ") {
			line = strings.TrimPrefix(line, "State: ")
			data["state"] += line
			continue
		}

		if strings.HasPrefix(line, "Body: ") {
			body = strings.TrimPrefix(line, "Body: ")
			afterBody = true
			continue
		}

		if afterBody {
			body = body + "\n" + line
		}
	}
	data["body"] += body

	if err = os.Remove(fpath); err != nil {
		log.Fatalf("Ошибка при удалении временного файла: %v", err)
	}
	return data
}
