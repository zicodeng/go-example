package main

import (
	"fmt"
	"github.com/zicodeng/go-example/websocket/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

// NotificationsHandler handles requests for the /notifications resource.
type NotificationsHandler struct {
	notifier *handlers.Notifier
}

// NewNotificationsHandler constructs a new NotificationsHandler.
func NewNotificationsHandler(notifier *handlers.Notifier) *NotificationsHandler {
	return &NotificationsHandler{notifier}
}

// ServeHTTP handles HTTP requests for the NotificationsHandler.
func (nh *NotificationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	msg := fmt.Sprintf("Notification pushed from the server at %s", time.Now().Format("15:04:05"))
	nh.notifier.Notify([]byte(msg))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = "localhost:4000"
	}

	notifier := handlers.NewNotifier()

	mux := http.NewServeMux()
	mux.Handle("/websockets", handlers.NewWebSocketsHandler(notifier))
	mux.Handle("/notifications", NewNotificationsHandler(notifier))

	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
