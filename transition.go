// Copyright (c) 2021 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package transition implements smooth transition.
package transition

import (
	"github.com/hslam/scheduler"
	"sync"
)

const lastsSize = 4

//Transition implements smooth transition.
type Transition struct {
	threshold   int
	lasts       [lastsSize]int
	cursor      int
	count       int
	concurrency func() int
	wg          sync.WaitGroup
	pipeline    scheduler.Scheduler
}

// NewTransition returns a new transition.
func NewTransition(threshold int, concurrency func() int) *Transition {
	return &Transition{
		threshold:   threshold,
		concurrency: concurrency,
		pipeline:    scheduler.New(1, &scheduler.Options{Threshold: 2}),
	}
}

// Close closes the smooth transition.
func (w *Transition) Close() {
	w.pipeline.Close()
}

// Smooth ensures a smooth transition from the low function to the high function.
func (w *Transition) Smooth(low func(), high func()) {
	w.smooth(func() {
		w.wg.Wait()
		low()
	}, func() {
		w.wg.Add(1)
		w.pipeline.Schedule(func() {
			high()
			w.wg.Done()
		})
	})
}

func (w *Transition) batch() (n int) {
	w.cursor++
	w.lasts[w.cursor%lastsSize] = w.concurrency()
	var max int
	for i := 0; i < lastsSize; i++ {
		if w.lasts[i] > max {
			max = w.lasts[i]
		}
	}
	return max
}

func (w *Transition) smooth(low func(), high func()) {
	batch := w.batch()
	w.count++
	if batch <= w.threshold {
		low()
		w.count = 0
	} else if batch <= w.threshold*w.threshold {
		if w.count < w.threshold {
			low()
		} else {
			high()
			if w.count == batch {
				w.count = 0
			}
		}
	} else {
		alpha := w.threshold*2 - (batch-1)/w.threshold
		if alpha > 1 {
			if w.count < alpha {
				low()
			} else {
				high()
				if w.count == batch {
					w.count = 0
				}
			}
		} else {
			high()
			if w.count == batch {
				w.count = 0
			}
		}
	}
}
