package faiss

/*
#include <faiss/c_api/IndexScalarQuantizer_c.h>
#include <faiss/c_api/Index_c.h>
*/
import "C"

// IndexScalarQuantizer is an index that stores quantized vectors and performs exhaustive
// search.
type IndexScalarQuantizer struct {
	Index
}

// NewIndexScalarQuantizer creates a new scalar quantizer index.
func NewIndexScalarQuantizer(d int, quantizer int, metric int) (*IndexScalarQuantizer, error) {
	var idx faissIndex
	if c := C.faiss_IndexScalarQuantizer_new_with(
		&idx.idx,
		C.idx_t(d),
		C.FaissQuantizerType(quantizer),
		C.FaissMetricType(metric),
	); c != 0 {
		return nil, getLastError()
	}
	return &IndexScalarQuantizer{&idx}, nil
}
