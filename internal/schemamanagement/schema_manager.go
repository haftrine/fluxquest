package schemamanagement

import (
	"github.com/haftrine/fluxquest/internal/idrf"
	"github.com/haftrine/fluxquest/internal/schemamanagement/schemaconfig"
)

// SchemaManager defines methods for schema discovery and preparation
type SchemaManager interface {
	DiscoverDataSets() ([]string, error)
	FetchDataSet(dataSetIdentifier string) (*idrf.DataSet, error)
	PrepareDataSet(*idrf.DataSet, schemaconfig.SchemaStrategy) error
}
