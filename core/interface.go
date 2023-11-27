package core

import (
	"context"

	"github.com/Video-Miner/vmp-lpms/segmenter"
	"github.com/Video-Miner/vmp-lpms/stream"
)

// RTMPSegmenter describes an interface for a segmenter
type RTMPSegmenter interface {
	SegmentRTMPToHLS(ctx context.Context, rs stream.RTMPVideoStream, hs stream.HLSVideoStream, segOptions segmenter.SegmenterOptions) error
}
