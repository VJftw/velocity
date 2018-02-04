package build

import (
	"github.com/asdine/storm"
	uuid "github.com/satori/go.uuid"
)

type StreamManager struct {
	db *streamStormDB
}

func NewStreamManager(
	db *storm.DB,
) *StreamManager {
	m := &StreamManager{
		db: newStreamStormDB(db),
	}
	return m
}

func (m *StreamManager) new(
	s *Step,
	name string,
) *Stream {
	return &Stream{
		ID: uuid.NewV3(uuid.NewV1(), s.ID).String(),
		// Step: s,
		Name: name,
	}
}

func (m *StreamManager) save(s *Stream) error {
	return m.db.save(s)
}
