package utils

import (
	"fmt"
	"runtime"
)
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v Kb", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v Kb", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
  return b / 1024
}
