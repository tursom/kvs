package pipe

import "github.com/tursom/GoCollections/lang"

type (
	PipelineHandler[V1, V2 any] interface {
		lang.Object
		Encode1(input []V1) []V2
		Encode2(input []V2) []V1
	}
)
