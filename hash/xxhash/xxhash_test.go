// Copyright 2021 The maryjane Authors
// SPDX-License-Identifier: BSD-3-Clause

package xxhash_bench

import (
	"flag"
	"fmt"
	"testing"

	"github.com/cespare/xxhash"
	xxhashv2 "github.com/cespare/xxhash/v2"
	"github.com/zeebo/xxh3"
)

var (
	cesparev1 bool
	cesparev2 bool
	zeebo     bool
)

func init() {
	flag.BoolVar(&cesparev1, "cesparev1", false, "run cespare/xxhash")
	flag.BoolVar(&cesparev2, "cesparev2", false, "run cespare/xxhash/v2")
	flag.BoolVar(&zeebo, "zeebo", false, "run zeebo/xxh3")
}

var acc uint64

func BenchmarkSum64String(b *testing.B) {
	sizes := []int{
		0, 1, 3, 4, 8, 9, 16, 17, 32,
		33, 64, 65, 96, 97, 128, 129, 240, 241,
		512, 1024, 100 * 1024,
	}

	if cesparev1 {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
				b.SetBytes(int64(size))
				d := string(make([]byte, size))
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					acc += xxhash.Sum64String(d)
				}
			})
		}
	}

	if cesparev2 {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
				b.SetBytes(int64(size))
				d := string(make([]byte, size))
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					acc += xxhashv2.Sum64String(d)
				}
			})
		}
	}

	if zeebo {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
				b.SetBytes(int64(size))
				d := string(make([]byte, size))
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					acc += xxh3.HashString(d)
				}
			})
		}
	}
}
