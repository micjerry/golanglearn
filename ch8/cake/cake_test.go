package cake

import (
	"testing"
	"time"
	"flag"
)

var defaults = Shop {
	Verbose:     true,
	Cakes:       20,
	BakeTime:    10 * time.Millisecond,
	NumIcers:    1,
	IceTime:     10 * time.Millisecond,
	IncribeTime: 10 * time.Millisecond,
}

func init() {
	testing.Init()
	if !flag.Parsed() {
		flag.Parse()
	}
	defaults.Verbose = testing.Verbose()
}

func Benchmark(b *testing.B) {
	cakeshop := defaults
	cakeshop.Work(b.N)
}

func BenchmarkBuffers(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) 
}

func BenchmarkVariable(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.IncribeStdDev = cakeshop.IncribeTime / 4
	cakeshop.Work(b.N) 
}

func BenchmarkVariableBuffers(b *testing.B) {
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.IncribeStdDev = cakeshop.IncribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) 
}

func BenchmarkSlowIcing(b *testing.B) {
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) 
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) 
}