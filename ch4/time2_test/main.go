// time2_test
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var err error

// TMPFile создаёт шаблон файла и вызывает текстовый редактор для его заполнения.
// Возвращает путь ко временному файлу типа string.
func TMPFile() string {
	fpath := os.TempDir() + "/github_issues.tmp"

	if err = ioutil.WriteFile(fpath, []byte("Title: \nBody: "), 0644); err != nil {
		log.Fatalf("Ошибка при создании шаблона файла: %v", err)
	}

	cmd := exec.Command("vim", fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Ошибка: %v\n", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Ошибка в время редактирования. Error: %v\n", err)
	} else {
		log.Printf("Успешно отредактировано.\n")
	}
	return fpath
}

// ParseFile сканирует врменный текстовый файл, затем удаляет его.
// Создаёт отображение, в котором записаны Title и Body Issue.
// Возвращает указатель на отображение.
func main() {
	fpath := TMPFile()
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

	// Debug info
	// for item := range data {
	// 	fmt.Printf("Key: %s\tValue: %s\n", item, data[item])
	// }

	// if err = os.Remove(fpath); err != nil {
	// 	log.Fatalf("Ошибка при удалении временного файла: %v", err)
	// }
	fmt.Printf("Отображение:\n%s\n", data)
}
