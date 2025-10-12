package routes

import (
	"fmt"
	"net/http"

	"github.com/Vlad06013/BotConstructor.git/domain/client/backendRedirecter/http/controllers"
)

func InitRoutes() {
	http.HandleFunc("/api/hello", controllers.SendMessage)

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
