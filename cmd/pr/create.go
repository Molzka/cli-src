package pr

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

type CreatePullRequestBody struct {
	Title        string   `json:"title"`
	Description  string   `json:"description,omitempty"`
	SourceBranch string   `json:"source_branch"`
	TargetBranch string   `json:"target_branch"`
	ForkRepoID   string   `json:"fork_repo_id,omitempty"`
	ReviewerIDs  []string `json:"reviewer_ids,omitempty"`
	Publish      bool     `json:"publish"`
}

var payload CreatePullRequestBody

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создать Pull Request",
	Long: `Создает Pull Request. 
Можно использовать флаги (--title, --body) или запустить команду 
в интерактивном режиме (промпт-ответ) без флагов.`,

	// Используем RunE для возврата ошибок
	RunE: func(cmd *cobra.Command, args []string) error {
		// Создаем 'reader' один раз для всех вводов
		reader := bufio.NewReader(os.Stdin)

		// Создаем 'client' для http requests
		token, err := config.LoadToken()

		if err != nil {
			return fmt.Errorf("Ошибка инициализации http клиента. Проверьте Ваш Token.")
		}

		client := api.NewSourceCraftClient(token)

		// ----------------------------------------------------
		// 1. Reponame
		// ----------------------------------------------------
		var reponame string

		for {
			fmt.Print("Вставьте название репозитория (username/repo): ")
			reponame = readLine(reader)
			_, err := client.DoRequest("GET", fmt.Sprintf("/repos/%s", reponame), nil)
			if err != nil {
				fmt.Print("[Ошибка] Репозиторий не существует. Попробуйте снова.")
				continue
			}
			break
		}

		// ----------------------------------------------------
		// 2. Title (Обязательно, макс 1024)
		// ----------------------------------------------------
		// Получаем 'title' из флага
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			// Если флаг не указан, запускаем промпт
			fmt.Println("ℹ️ Заголовок PR не указан. Включен интерактивный режим.")
			for {
				fmt.Print("Title: ")
				title = readLine(reader)
				if title == "" {
					fmt.Println("   [Ошибка] Заголовок не может быть пустым. Попробуйте снова.")
					continue
				}
				if len(title) > 1024 {
					fmt.Printf("   [Ошибка] Заголовок слишком длинный: %d символов (макс 1024). Попробуйте снова.\n", len(title))
					continue
				}
				break // Все в порядке
			}
		}
		payload.Title = title

		// ----------------------------------------------------
		// 3. Description (Опционально)
		// ----------------------------------------------------
		description, _ := cmd.Flags().GetString("body")
		if !cmd.Flags().Changed("body") {
			// Промпт, только если флаг НЕ был установлен (даже если он был пустым)
			fmt.Print("Description (optional): ")
			description = readLine(reader)
		}
		payload.Description = description

		// ----------------------------------------------------
		// 4. Source Branch (Обязательно)
		// ----------------------------------------------------
		sourceBranch, _ := cmd.Flags().GetString("head")
		if sourceBranch == "" {
			// Если флаг не указан, пытаемся получить текущую ветку
			currentBranch := getCurrentGitBranch()
			prompt := "Source Branch: "
			if currentBranch != "" {
				prompt = fmt.Sprintf("Source Branch (default: %s): ", currentBranch)
			}

			for {
				fmt.Print(prompt)
				sourceBranch = readLine(reader)
				if sourceBranch == "" && currentBranch != "" {
					sourceBranch = currentBranch // Используем default
				}
				if sourceBranch == "" {
					fmt.Println("   [Ошибка] Source Branch не может быть пустым. Попробуйте снова.")
					continue
				}
				break
			}
		}
		payload.SourceBranch = sourceBranch

		// ----------------------------------------------------
		// 5. Target Branch (Обязательно, с default)
		// ----------------------------------------------------
		// У флага уже есть default "main", поэтому промпт не нужен.
		// Cobra автоматически присвоит "main", если флаг не указан.
		targetBranch, _ := cmd.Flags().GetString("base")
		payload.TargetBranch = targetBranch

		// ----------------------------------------------------
		// 6. Reviewer IDs (Опционально, список)
		// ----------------------------------------------------
		reviewers, _ := cmd.Flags().GetStringSlice("reviewers")
		if !cmd.Flags().Changed("reviewers") {
			fmt.Print("Reviewer IDs (optional, через запятую): ")
			reviewerInput := readLine(reader)
			if reviewerInput != "" {
				reviewers = parseCommaSeparatedList(reviewerInput)
			}
		}
		payload.ReviewerIDs = reviewers

		// ----------------------------------------------------
		// 7. Publish (Логика 'draft')
		// ----------------------------------------------------
		// API ожидает 'publish', а флаг удобнее назвать '--draft'
		isDraft, _ := cmd.Flags().GetBool("draft")

		if !cmd.Flags().Changed("draft") {
			// Если флаг --draft не был использован, спрашиваем
			for {
				// По умолчанию (N) - 'false', что создаст Draft
				fmt.Print("Опубликовать сразу (не как черновик)? (y/N): ")
				publishInput := strings.ToLower(readLine(reader))

				if publishInput == "y" || publishInput == "yes" {
					isDraft = false // (т.е. Publish = true)
					break
				}
				if publishInput == "n" || publishInput == "no" || publishInput == "" {
					isDraft = true // (т.е. Publish = false)
					break
				}
				fmt.Println("   [Ошибка] Введите 'y' (да) или 'n' (нет).")
			}
		}
		// Инвертируем, так как API ожидает 'publish'
		payload.Publish = !isDraft

		// ----------------------------------------------------
		// 8. ОТПРАВКА ДАННЫХ
		// ----------------------------------------------------
		fmt.Println("\n✅ Данные успешно собраны!")
		fmt.Println("----------------------")
		fmt.Printf("Title: %s\n", payload.Title)
		fmt.Printf("Source: %s -> Target: %s\n", payload.SourceBranch, payload.TargetBranch)
		fmt.Printf("Publish (Ready): %t\n", payload.Publish)
		fmt.Println("----------------------")
		fmt.Println("Вызов вашей функции sendRequest(payload)...")

		//
		// ЗДЕСЬ ВЫЗЫВАЙТЕ ВАШУ ФУНКЦИЮ
		//

		_, err = client.DoRequest("POST", fmt.Sprintf("/repos/%s/pulls", reponame), payload)

		if err != nil {
			return fmt.Errorf("ошибка при отправке запроса: %w", err)
		}

		fmt.Println("Pull Request успешно создан!")
		return nil // Возвращаем nil в случае успеха
	},
}

// init() регистрирует команду и ее флаги
func init() {
	// Добавляем эту команду в вашу корневую команду (rootCmd)
	// Убедитесь, что rootCmd определен в cmd/root.go
	// rootCmd.AddCommand(prCreateCmd)

	// Определяем флаги
	// Имя, Сокращение, Значение по умолч., Описание
	createCmd.Flags().StringP("title", "t", "", "Заголовок Pull Request")
	createCmd.Flags().StringP("body", "b", "", "Описание Pull Request (body)")

	createCmd.Flags().StringP("head", "H", "", "Исходная ветка (source branch). По умолч. - текущая")
	createCmd.Flags().StringP("base", "B", "main", "Целевая ветка (target branch)")

	createCmd.Flags().StringSliceP("reviewers", "r", nil, "Список Reviewer ID (через запятую)")

	// Флаг 'draft' (черновик)
	// По умолчанию (false) PR будет опубликован (payload.Publish = true)
	createCmd.Flags().Bool("draft", false, "Создать PR как черновик (publish=false)")
}

/*
// ЗАГЛУШКА ДЛЯ ВАШЕЙ ФУНКЦИИ (поместите ее куда нужно)
func sendRequest(payload CreatePullRequestBody) error {
	// ... ваша логика net/http
	fmt.Printf("Отправка JSON: %+v\n", payload)
	return nil
}
*/
