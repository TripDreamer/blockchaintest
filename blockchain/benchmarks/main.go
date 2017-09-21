package main

import (
	"fmt"
	"time"

	"blockchain/core"
)

func main() {

	fmt.Println("Benchmarking...")

	t1 := benchmark(func() {
		t := core.NewTransaction(nil, nil, nil)
		t.GenerateNonce(core.TRANSACTION_POW) //产生随机数
		t.Sign(core.GenerateNewKeypair())
	})
	fmt.Println("Transaction took", t1)

	t2 := benchmark(func() {
		b := core.NewBlock(nil)
		b.GenerateMerkelRoot()
		b.GenerateNonce(core.BLOCK_POW)
		b.Sign(core.GenerateNewKeypair())
	})
	fmt.Println("Block took", t2)
}

func benchmark(f func()) time.Duration {

	t0 := time.Now()

	f()

	return time.Since(t0)
}
