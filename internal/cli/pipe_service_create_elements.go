package cli

import (
	"fmt"

	"github.com/haftrine/fluxquest/internal/connections"
	"github.com/haftrine/fluxquest/internal/extraction"
	extrConfig "github.com/haftrine/fluxquest/internal/extraction/config"
	"github.com/haftrine/fluxquest/internal/ingestion"
	ingConfig "github.com/haftrine/fluxquest/internal/ingestion/config"
	influx "github.com/influxdata/influxdb/client/v2"
)

func (p *pipeService) createElements(
	infConn influx.Client,
	tsConn connections.PgxWrap,
	extrConf *extrConfig.ExtractionConfig,
	ingConf *ingConfig.IngestorConfig) (extraction.Extractor, ingestion.Ingestor, error) {
	extractor, err := p.extractorService.InfluxExtractor(infConn, extrConf)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create extractor\n%v", err)
	}

	ingestor := p.ingestorService.NewTimescaleIngestor(tsConn, ingConf)
	return extractor, ingestor, nil
}
