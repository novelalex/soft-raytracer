package geom

import "sync/atomic"

var nextID uint64

func newId() uint64 {
	return atomic.AddUint64(&nextID, 1)
}
