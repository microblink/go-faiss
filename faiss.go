// Package faiss provides bindings to Faiss, a library for vector similarity
// search.
// More detailed documentation can be found at the Faiss wiki:
// https://github.com/facebookresearch/faiss/wiki.
package faiss

/*
#cgo LDFLAGS: -lfaiss_c

#include <faiss/c_api/IndexScalarQuantizer_c.h>
#include <faiss/c_api/Index_c.h>
#include <faiss/c_api/error_c.h>
*/
import "C"
import "errors"

func getLastError() error {
	return errors.New(C.GoString(C.faiss_get_last_error()))
}

// Metric type
const (
	MetricInnerProduct  = C.METRIC_INNER_PRODUCT
	MetricL2            = C.METRIC_L2
	MetricL1            = C.METRIC_L1
	MetricLinf          = C.METRIC_Linf
	MetricLp            = C.METRIC_Lp
	MetricCanberra      = C.METRIC_Canberra
	MetricBrayCurtis    = C.METRIC_BrayCurtis
	MetricJensenShannon = C.METRIC_JensenShannon
)

// Quantizer type
const (
	QuantizerQT_8bit         = C.QT_8bit         ///< 8 bits per component
	QuantizerQT_4bit         = C.QT_4bit         ///< 4 bits per component
	QuantizerQT_8bit_uniform = C.QT_8bit_uniform ///< same, shared range for all dimensions
	QuantizerQT_4bit_uniform = C.QT_4bit_uniform
	QuantizerQT_fp16         = C.QT_fp16
	QuantizerQT_8bit_direct  = C.QT_8bit_direct ///< fast indexing of uint8s
	QuantizerQT_6bit         = C.QT_6bit        ///< 6 bits per component
)
