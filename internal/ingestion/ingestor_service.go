package ingestion

import (
	"github.com/haftrine/fluxquest/internal/connections"
	"github.com/haftrine/fluxquest/internal/ingestion/config"
	"github.com/haftrine/fluxquest/internal/ingestion/ts"
	tsSchema "github.com/haftrine/fluxquest/internal/schemamanagement/ts"
)

// IngestorService exposes methods to create new ingestors
type IngestorService interface {
	NewTimescaleIngestor(dbConn connections.PgxWrap, config *config.IngestorConfig) Ingestor
}

// NewIngestorService creates an instance of the IngestorService
func NewIngestorService() IngestorService {
	return &ingestorService{}
}

type ingestorService struct {
}

// NewIngestor creates a new instance of an Ingestor with a specified config, for a specified
// data set and data channel
func (i *ingestorService) NewTimescaleIngestor(dbConn connections.PgxWrap, config *config.IngestorConfig) Ingestor {
	schemaManager := tsSchema.NewTSSchemaManager(dbConn, config.Schema, config.ChunkTimeInterval)
	return &ts.TSIngestor{
		DbConn:           dbConn,
		Config:           config,
		IngestionRoutine: ts.NewRoutine(),
		SchemaManager:    schemaManager,
	}
}
