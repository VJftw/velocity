package websocket

import (
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ID:            uuid.NewV4().String(),
		ws:            ws,
		subscriptions: []string{},
	}
}

type Client struct {
	ID            string
	ws            *websocket.Conn
	subscriptions []string
}

func (c *Client) Subscribe(s string, ref uint64) {
	c.subscriptions = append(c.subscriptions, s)
	c.ws.WriteJSON(PhoenixMessage{
		Event: PhxReplyEvent,
		Topic: s,
		Ref:   ref,
		Payload: PhoenixReplyPayload{
			Status:   "ok",
			Response: map[string]string{},
		},
	})
}

func (c *Client) Unsubscribe(s string, ref uint64) {
	var element int
	for i, v := range c.subscriptions {
		if v == s {
			element = i
			break
		}
	}
	if element < len(c.subscriptions) {
		c.subscriptions = append(c.subscriptions[:element], c.subscriptions[element+1:]...)
	}
	c.ws.WriteJSON(PhoenixMessage{
		Event: PhxReplyEvent,
		Topic: s,
		Ref:   ref,
		Payload: PhoenixReplyPayload{
			Status:   "ok",
			Response: map[string]string{},
		},
	})
}

func (c *Client) HandleHeartbeat(ref uint64) {
	c.ws.WriteJSON(PhoenixMessage{
		Event: PhxReplyEvent,
		Topic: PhxSystemTopic,
		Ref:   ref,
		Payload: PhoenixReplyPayload{
			Status:   "ok",
			Response: map[string]string{},
		},
	})
}

func (c *Client) HandleMessage(m *PhoenixMessage) {
	if m.Topic == PhxSystemTopic {
		switch m.Event {
		case PhxHeartbeatEvent:
			c.HandleHeartbeat(m.Ref)
			break
		}
		return
	}

	switch m.Event {
	case PhxJoinEvent:
		c.Subscribe(m.Topic, m.Ref)
		break
	case PhxLeaveEvent:
		c.Unsubscribe(m.Topic, m.Ref)
		break
	}
}

// Channel Event constants
const (
	PhxCloseEvent     = "phx_close"
	PhxErrorEvent     = "phx_error"
	PhxJoinEvent      = "phx_join"
	PhxReplyEvent     = "phx_reply"
	PhxLeaveEvent     = "phx_leave"
	PhxHeartbeatEvent = "heartbeat"
	PhxSystemTopic    = "phoenix"
)

// Velocity specific events
const (
	VNewProject    = "project:new"
	VUpdateProject = "project:update"
	VDeleteProject = "project:delete"

	VNewCommit    = "commit:new"
	VUpdateCommit = "commit:update"
	VDeleteCommit = "commit:delete"

	VNewBranch    = "branch:new"
	VUpdateBranch = "branch:update"
	VDeleteBranch = "branch:delete"

	VNewKnownHost    = "knownhost:new"
	VDeleteKnownHost = "knownhost:delete"

	VNewBuild    = "build:new"
	VUpdateBuild = "build:update"
	VDeleteBuild = "build:delete"

	VUpdateStep = "step:update"
	VDeleteStep = "step:delete"

	VUpdateStream = "stream:update"
	VDeleteStream = "stream:delete"

	VNewStreamLine = "streamLine:new"

	VUpdateSlave = "slave:update"
	VNewSlave    = "slave:new"
	VDeleteSlave = "slave:delete"
)

type PhoenixMessage struct {
	Event   string      `json:"event"`
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
	Ref     uint64      `json:"ref"`
}

type PhoenixReplyPayload struct {
	Status   string      `json:"status"`
	Response interface{} `json:"response"`
}

type BuildMessage struct {
	Step   uint64     `json:"step"`
	Status string     `json:"status"`
	Log    LogMessage `json:"log"`
}

type LogMessage struct {
	Timestamp time.Time `json:"timestamp"`
	Output    string    `json:"output"`
}
