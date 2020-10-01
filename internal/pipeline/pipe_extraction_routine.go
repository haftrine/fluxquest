package pipeline

import (
	"sync"

	"github.com/haftrine/fluxquest/internal/extraction"
)

type extractorRoutineArgs struct {
	wg *sync.WaitGroup
	e  extraction.Extractor
	eb func(error)
	ec chan error
}

func extractorRoutine(args *extractorRoutineArgs) {
	err := args.e.Start(args.ec)
	if err != nil {
		args.eb(err)
	}
	args.wg.Done()
}
