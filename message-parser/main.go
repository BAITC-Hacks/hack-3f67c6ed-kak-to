package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		message := scanner.Text()

		category, answer := parseMessage(message)

		fmt.Println("Обращение:", message)
		fmt.Println("Категория:", category)
		fmt.Println("Черновик ответа:", answer)
		fmt.Println("------------------------------")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}
}

func parseMessage(message string) (string, string) {
	text := strings.ToLower(message)

	switch {
	case strings.Contains(text, "справк"):
		return "справка",
			"Здравствуйте! Справку о месте учёбы можно получить через учебную часть. Уточните, пожалуйста, если нужна информация о конкретном способе получения."

	case strings.Contains(text, "столов") ||
		strings.Contains(text, "холодн") ||
		strings.Contains(text, "wi-fi") ||
		strings.Contains(text, "wifi") ||
		strings.Contains(text, "пропал"):
		return "жалоба",
			"Здравствуйте! Спасибо за обращение. Передадим информацию ответственным сотрудникам для проверки ситуации."

	default:
		return "другое",
			"Здравствуйте! Спасибо за обращение. Уточним информацию и подскажем, куда лучше обратиться по вашему вопросу."
	}
}