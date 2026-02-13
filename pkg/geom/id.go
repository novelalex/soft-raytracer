package geom

import "sync/atomic"

type id uint64

var nextID uint64

func newId() uint64 {
	return atomic.AddUint64(&nextID, 1)
}
