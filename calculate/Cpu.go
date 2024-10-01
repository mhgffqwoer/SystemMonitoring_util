package calculate

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var (
	CpuCalculateCh = make(chan []float64, 5)
)

func CpuCalculate() {
	go func() {
		for {
			CpuCalculate, _ := cpu.Percent(time.Second, true)
			CpuCalculateCh <- CpuCalculate
		}
	}()
}