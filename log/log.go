package log

import "sync/atomic"

var Count int64

func Add() {

	atomic.AddInt64(&Count, 1)
}
