// 非同步uint64計次
//
// 必須要先宣告才能使用
//
// var counter.Counter
package counter

import (
	"sync/atomic"
)

// The Structure for Counter
type Counter struct {
	time uint64
}

// 計次+1
//
// Counter time +1
func (this *Counter) Add() uint64 {
	return atomic.AddUint64(&this.time, 1)
}

// 重設計次
//
// Reset Counter
func (this *Counter) Reset() {
	atomic.StoreUint64(&this.time, 0)
}

// 回傳當前計次
//
// Reponse the Counter time
func (this *Counter) Time() uint64 {
	return this.time
}