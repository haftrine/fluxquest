package cli

import (
	"fmt"

	"github.com/haftrine/fluxquest/internal/transformation"
	influx "github.com/influxdata/influxdb/client/v2"
)

const (
	transformerIDTemplate = "%s_transfomer_%s"
)

func (p *pipeService) createTransformers(pipeID string, infConn influx.Client, measure string, inputDb string, conf *MigrationConfig) ([]transformation.Transformer, error) {
	transformers := []transformation.Transformer{}

	if conf.TagsAsJSON {
		id := fmt.Sprintf(transformerIDTemplate, pipeID, "tagsAsJSON")
		tagsTransformer, err := p.transformerService.TagsAsJSON(infConn, id, inputDb, conf.RetentionPolicy, measure, conf.TagsCol)
		if err != nil {
			return nil, err
		}
		// if measurement has no tags, a nil transformer is returned
		if tagsTransformer != nil {
			transformers = append(transformers, tagsTransformer)
		}
	}

	if conf.FieldsAsJSON {
		id := fmt.Sprintf(transformerIDTemplate, pipeID, "fieldsAsJSON")
		fieldsTransformer, err := p.transformerService.FieldsAsJSON(infConn, id, inputDb, conf.RetentionPolicy, measure, conf.FieldsCol)
		if err != nil {
			return nil, err
		}
		transformers = append(transformers, fieldsTransformer)
	}

	return transformers, nil
}
