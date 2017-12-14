package websocket

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// Controller - Handles Websockets
type Controller struct {
	logger  *log.Logger
	render  *render.Render
	manager *Manager
}

// NewController - returns a new Controller for client Websockets.
func NewController(websocketManager *Manager) *Controller {
	return &Controller{
		logger:  log.New(os.Stdout, "[controller:websocket]", log.Lshortfile),
		render:  render.New(),
		manager: websocketManager,
	}
}

// Setup - Sets up the Websocket Controller
func (c Controller) Setup(router *mux.Router) {

	// /v1/ws
	router.Handle("/v1/ws", negroni.New(
		negroni.Wrap(http.HandlerFunc(c.wsClientHandler)),
	)).Methods("GET")

	c.logger.Println("Set up Websocket controller.")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c Controller) wsClientHandler(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	wsClient := NewClient(ws)

	c.manager.Save(wsClient)

	// Monitor for Messages
	go c.monitor(wsClient)
}

func (c *Controller) monitor(client *Client) {
	for {
		message := &PhoenixMessage{}
		err := client.ws.ReadJSON(message)
		if err != nil {
			log.Println(err)
			log.Println("Closing Client WebSocket")
			client.ws.Close()
			c.manager.Remove(client)
			return
		}
		client.HandleMessage(message)
	}
}
