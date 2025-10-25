package repo

import (
	"bufio"
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/config"
	"src/internal/utils"
	"strings"

	"github.com/spf13/cobra"
)

// ... (структуры InitSettings, TemplatingOptions, CreateRepositoryBody остаются без изменений) ...

// InitSettings cодержит опции для инициализации репозитория.
type InitSettings struct {
	DefaultBranch       string   `json:"default_branch,omitempty"`
	CreateReadme        bool     `json:"create_readme"`
	GitignorePresets    []string `json:"gitignore_presets,omitempty"`
	LicenseSlug         string   `json:"license_slug,omitempty"`
	SrcYamlTemplateSlug string   `json:"src_yaml_template_slug,omitempty"`
}

// TemplatingOptions cодержит опции для создания из шаблона.
type TemplatingOptions struct {
	TemplateID string `json:"template_id"`
}

// CreateRepositoryBody - это тело запроса для CОЗДАНИЯ РЕПОЗИТОРИЯ.
type CreateRepositoryBody struct {
	Name              string             `json:"name"`
	Slug              string             `json:"slug"`
	Description       string             `json:"description,omitempty"`
	Visibility        string             `json:"visibility,omitempty"`
	InitSettings      *InitSettings      `json:"init_settings,omitempty"`
	TemplatingOptions *TemplatingOptions `json:"templating_options,omitempty"`
}

var payload CreateRepositoryBody
var orgSlug string // Переменная для хранения org_slug

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создать новый репозиторий в организации",
	Long: `Создает новый репозиторий в указанной организации.
Можно использовать флаги (--org, --name, --slug) или запустить команду
в интерактивном режиме (промпт-ответ) без флагов.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		token, err := config.LoadToken()
		if err != nil {
			return fmt.Errorf("Ошибка инициализации http клиента. Проверьте Ваш Token.")
		}
		client := api.NewSourceCraftClient(token)

		// ----------------------------------------------------
		// 1. Org Slug (Обязательно для URL)
		// ----------------------------------------------------
		orgSlug, _ = cmd.Flags().GetString("org")
		if orgSlug == "" {
			fmt.Println("ℹ️ Не указана организация. Включен интерактивный режим.")
			for {
				fmt.Print("Slug организации (org_slug): ")
				orgSlug = utils.ReadLine(reader)
				if orgSlug == "" {
					fmt.Println("   [Ошибка] Slug организации не может быть пустым. Попробуйте снова.")
					continue
				}
				break
			}
		}

		// ----------------------------------------------------
		// 2. Name (Обязательно, макс 256)
		// ----------------------------------------------------
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			// (Оставил 'if name == ""' на случай, если --org указан, а --name нет)
			fmt.Println("ℹ️ Название репозитория не указано. Включен интерактивный режим.")
			for {
				fmt.Print("Название репозитория (name): ")
				name = utils.ReadLine(reader)
				if name == "" {
					fmt.Println("   [Ошибка] Название не может быть пустым. Попробуйте снова.")
					continue
				}
				if len(name) > 256 {
					fmt.Printf("   [Ошибка] Название слишком длинное: %d символов (макс 256). Попробуйте снова.\n", len(name))
					continue
				}
				break
			}
		}
		payload.Name = name

		// ----------------------------------------------------
		// 3. Slug (Обязательно, макс 256)
		// ----------------------------------------------------
		slug, _ := cmd.Flags().GetString("slug")
		if slug == "" {
			prompt := fmt.Sprintf("Slug репозитория (default: %s): ", name)
			for {
				fmt.Print(prompt)
				slug = utils.ReadLine(reader)
				if slug == "" {
					slug = name
				}
				if len(slug) > 256 {
					fmt.Printf("   [Ошибка] Slug слишком длинный: %d символов (макс 256). Попробуйте снова.\n", len(slug))
					slug = ""
					continue
				}
				break
			}
		}
		payload.Slug = slug

		// ... (Остальная логика: Description, Visibility, Templating, Init Settings) ...
		// (Код для шагов 4, 5, 6 не изменился)

		// ----------------------------------------------------
		// 4. Description (Опционально)
		// ----------------------------------------------------
		description, _ := cmd.Flags().GetString("description")
		if !cmd.Flags().Changed("description") {
			fmt.Print("Описание (description, optional): ")
			description = utils.ReadLine(reader)
		}
		payload.Description = description

		// ----------------------------------------------------
		// 5. Visibility (Опционально, enum)
		// ----------------------------------------------------
		visibility, _ := cmd.Flags().GetString("visibility")
		if !cmd.Flags().Changed("visibility") {
			for {
				fmt.Print("Сделать репозиторий приватным? (y/N): ")
				input := strings.ToLower(utils.ReadLine(reader))

				if input == "y" || input == "yes" {
					visibility = "private"
					break
				}
				if input == "n" || input == "no" || input == "" {
					visibility = "public"
					break
				}
				fmt.Println("   [Ошибка] Введите 'y' (да) или 'n' (нет).")
			}
		}
		if visibility != "public" && visibility != "private" && visibility != "internal" {
			fmt.Printf("[Warning] Неверное значение --visibility: '%s'. Используется 'public'.\n", visibility)
			visibility = "public"
		}
		payload.Visibility = visibility

		// ----------------------------------------------------
		// 6. Templating Options (Опционально)
		// ----------------------------------------------------
		templateID, _ := cmd.Flags().GetString("template-id")
		if !cmd.Flags().Changed("template-id") {
			fmt.Print("Создать из шаблона? (y/N): ")
			input := strings.ToLower(utils.ReadLine(reader))
			if input == "y" || input == "yes" {
				for {
					fmt.Print("ID репозитория-шаблона (template_id): ")
					templateID = utils.ReadLine(reader)
					if templateID == "" {
						fmt.Println("   [Ошибка] ID шаблона не может быть пустым. Попробуйте снова.")
						continue
					}
					break
				}
			}
		}
		if templateID != "" {
			payload.TemplatingOptions = &TemplatingOptions{TemplateID: templateID}
		}

		// ----------------------------------------------------
		// 7. Init Settings (Опционально, *если не из шаблона*)
		// ----------------------------------------------------
		if payload.TemplatingOptions == nil {
			fmt.Print("Инициализировать репозиторий (README, .gitignore)? (y/N): ")
			input := strings.ToLower(utils.ReadLine(reader))
			if input == "y" || input == "yes" {
				initSettings := &InitSettings{}

				fmt.Print("Создать README.md? (Y/n): ")
				readmeInput := strings.ToLower(utils.ReadLine(reader))
				if readmeInput == "y" || readmeInput == "yes" || readmeInput == "" {
					initSettings.CreateReadme = true
				}

				fmt.Print("Название default-ветки (optional, default: main): ")
				branchInput := utils.ReadLine(reader)
				if branchInput != "" {
					initSettings.DefaultBranch = branchInput
				}

				fmt.Print("Шаблоны .gitignore (optional, через запятую, e.g., 'go,node'): ")
				gitignoreInput := utils.ReadLine(reader)
				if gitignoreInput != "" {
					initSettings.GitignorePresets = utils.ParseCommaSeparatedList(gitignoreInput)
				}

				fmt.Print("Лицензия (optional, e.g., 'mit', 'apache-2.0'): ")
				initSettings.LicenseSlug = utils.ReadLine(reader)

				fmt.Print("Шаблон CI (optional, src_yaml_template_slug): ")
				initSettings.SrcYamlTemplateSlug = utils.ReadLine(reader)

				payload.InitSettings = initSettings
			}
		}

		// ----------------------------------------------------
		// 8. ОТПРАВКА ДАННЫХ
		// ----------------------------------------------------
		fmt.Println("\n✅ Данные для создания репозитория собраны!")
		fmt.Println("----------------------")
		fmt.Printf("Организация: %s\n", orgSlug) // <-- Добавлен вывод orgSlug
		fmt.Printf("Name: %s\n", payload.Name)
		fmt.Printf("Slug: %s\n", payload.Slug)
		fmt.Printf("Visibility: %s\n", payload.Visibility)
		// ... (остальной вывод) ...
		fmt.Println("----------------------")

		// --- ИЗМЕНЕННЫЙ ENDPOINT ---
		endpoint := fmt.Sprintf("/orgs/%s/repos", orgSlug)
		fmt.Printf("Вызов POST %s...\n", endpoint)

		_, err = client.DoRequest("POST", endpoint, payload)

		if err != nil {
			return fmt.Errorf("ошибка при отправке запроса: %w", err)
		}

		fmt.Println("Репозиторий успешно создан!")
		return nil
	},
}

func init() {
	// rootCmd.AddCommand(createCmd)

	// --- ДОБАВЛЕН ФЛАГ 'org' ---
	createCmd.Flags().String("org", "", "Slug организации (org_slug) (обязательно)")

	createCmd.Flags().StringP("name", "n", "", "Название репозитория (обязательно)")
	createCmd.Flags().StringP("slug", "s", "", "Slug репозитория (e.g., 'my-repo'). По умолч. - 'name'")
	createCmd.Flags().StringP("description", "d", "", "Описание репозитория")
	createCmd.Flags().String("visibility", "", "Видимость: 'public', 'private' или 'internal'")
	createCmd.Flags().String("template-id", "", "ID репозитория-шаблона для создания")
}
