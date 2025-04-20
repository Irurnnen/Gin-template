package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Irurnnen/gin-template/internal/config"
)

type Application struct {
	Config config.Config
	Debug  bool
}

func New() *Application {
	return &Application{
		Config: *config.NewConfigExample(),
		Debug:  false,
	}
}

func NewDebug() *Application {
	return &Application{
		Config: *config.NewConfigExample(),
		Debug:  true,
	}
}

func (a *Application) Run() error {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":8080", nil)
	return err
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
