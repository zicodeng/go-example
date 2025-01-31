package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketsHandler is a handler for WebSocket upgrade requests.
type WebSocketsHandler struct {
	notifier *Notifier
	upgrader *websocket.Upgrader
}

// NewWebSocketsHandler constructs a new WebSocketsHandler.
func NewWebSocketsHandler(notifier *Notifier) *WebSocketsHandler {
	return &WebSocketsHandler{
		notifier: notifier,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

// ServeHTTP implements the http.Handler interface for the WebSocketsHandler.
func (wsh *WebSocketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("received websocket upgrade request")
	// Upgrade the connection to a WebSocket, and add the
	// new websock.Conn to the Notifier.
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Each request will be running on its own goroutine,
	// and represent a different client.
	// So whenever we receive a new request, and ServeHTTP is called,
	// we need to add that request as a new client to our Notifier's clients field.
	// Note that we don't want to spawn a new goroutine here
	// because the expectation is that this upgrade request never ends,
	// as the connection is upgraded into a persistent Websocket connection.
	wsh.notifier.AddClient(conn)
}

// Notifier is an object that handles WebSocket notifications.
type Notifier struct {
	// slice and channels are reference type.
	// We need to initialize it somehow
	// otherwise their zero value is nil,
	// and we might get nil pointer reference error.
	clients []*websocket.Conn
	eventQ  chan []byte
	// Add a mutex or other channels to
	// protect the `clients` slice from concurrent use.
	// Our NewNotifier() doesn't need to initialize mx field
	// and we are still able to use it,
	// because we are just using zero values for whatever in the Mutex struct fields.
	mx sync.Mutex
}

// NewNotifier constructs a new Notifier.
func NewNotifier() *Notifier {
	// Construct a new Notifier
	// and call the .start() method on
	// a new goroutine to start the
	// event notification loop.
	notifier := &Notifier{
		eventQ: make(chan []byte),
	}
	go notifier.start()
	return notifier
}

// AddClient adds a new client to the Notifier.
func (n *Notifier) AddClient(client *websocket.Conn) {
	log.Println("adding new WebSockets client")

	// Add the client to the `clients` slice
	// but since this can be called from multiple
	// goroutines, make sure you protect the `clients`
	// slice while you add a new connection to it!
	n.mx.Lock()
	n.clients = append(n.clients, client)
	n.mx.Unlock()

	// Process incoming control messages from the client.
	// Once this client is added to the list, it will constantly
	// send control messages to our server. If at one point,
	// we get an error when reading those control messages,
	// it informs us that connection is lost, and we need to
	// remove it from the list.
	for {
		if _, _, err := client.NextReader(); err != nil {
			client.Close()
			// Remove it from the list
			n.mx.Lock()
			for i, c := range n.clients {
				if c == client {
					n.clients = append(n.clients[:i], n.clients[i+1:]...)
				}
			}
			n.mx.Unlock()
			break
		}
	}
}

// Notify broadcasts the event to all WebSocket clients
// by sending an event to the eventQ.
func (n *Notifier) Notify(event []byte) {
	log.Printf("adding event to the queue")
	// Add `event` to the `n.eventQ`
	n.eventQ <- event
}

// Start starts the notification loop.
func (n *Notifier) start() {
	log.Println("starting notifier loop")
	// Start a never-ending loop that reads
	// new events out of the `n.eventQ` and broadcasts
	// them to all WebSocket clients.
	for msg := range n.eventQ {
		n.mx.Lock()
		// Loop through all the existing clients,
		// and send messages to all of them.
		for i, c := range n.clients {
			// If we encounter an error writing messages out,
			// it means this connection is lost,
			// and we need to remove it from the list.
			if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
				c.Close()
				n.clients = append(n.clients[:i], n.clients[i+1:]...)
				log.Println(err)
			}
		}
		n.mx.Unlock()
	}
}
