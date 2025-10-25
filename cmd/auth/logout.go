package auth

import (
	"fmt"
	"src/internal/config"

	"github.com/spf13/cobra"
)

func init() {
	AuthCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Выход из системы",
	Run:   runLogout,
}

func runLogout(cmd *cobra.Command, args []string) {
	token, err := config.LoadToken()

	if err != nil {
		if token == "" {
			fmt.Println("Токен уже удален.")
			return
		}
		fmt.Println("Ошибка при удалении токена:", err)
		return
	} else {
		err := config.DeleteToken()
		if err != nil {
			fmt.Println("Ошибка при удалении токена:", err)
			return
		}
	}
	fmt.Println("Токен успешно удален!")

}
