package ingestion

import (
	"github.com/haftrine/fluxquest/internal/idrf"
)

// Ingestor takes a data channel of idrf rows and inserts them in a target database
type Ingestor interface {
	ID() string
	Prepare(conn *idrf.Bundle) error
	Start(chan error) error
}
