package internal

import (
	"net/http"

	"github.com/video-hosting-platform/subscription-service/internal/handler"
)

func Run() {
	handler := handler.New(nil)
	http.ListenAndServe(":8081", handler.Init())
}
