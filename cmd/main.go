package main

// go mod init mod - чтобы разрешить go использовать модули с github. go get ...ссылка с гитхаб, чтобы установить модуль в проект
import (
	"log"
	//"net/http"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
