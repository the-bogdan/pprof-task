package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"pprof-task/fast"
	"pprof-task/slow"
)

// запускаем перед основными функциями по разу чтобы файл остался в памяти в файловом кеше
// ioutil.Discard - это ioutil.Writer который никуда не пишет
func init() {
	slow.Search(filePath, ioutil.Discard)
	fast.Search(filePath, ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	slow.Search(filePath, slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	fast.Search(filePath, fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slow.Search(filePath, ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fast.Search(filePath, ioutil.Discard)
	}
}
