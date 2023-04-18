package pipe

import "github.com/tursom/GoCollections/lang"

type (
	PipelineHandler[V1, V2 any] interface {
		lang.Object
		encode1(input []V1) []V2
		encode2(input []V2) []V1
	}
)
