package auth

import (
	"fmt"
	"os"
	"src/internal/api"
	"src/internal/config"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func init() {
	AuthCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Войти в систему",
	Run:   runLogin,
}

func runLogin(cmd *cobra.Command, args []string) {
	fmt.Print("Введите Personal Access Token: ")
	byteToken, err := term.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println("\nОшибка при чтении токена:", err)
		return
	}

	token := string(byteToken)

	if !api.ValidateToken(token) {
		fmt.Println("Неверный токен")
		return
	}

	if err := config.SaveToken(token); err != nil {
		fmt.Println("Ошибка сохранения токена:", err)
		return
	}
	fmt.Println("Вход успешно совершен!")
}
