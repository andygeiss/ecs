package plugins

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/_examples/engine"
	tm "github.com/buger/goterm"
	"runtime"
	"time"
)

// ShowEngineStats ...
func ShowEngineStats(em *ecs.EntityManager) ecs.Plugin {
	frameTime := time.Now()
	updateTime := time.Now()
	// Return a plugin which will be called by the renderer.
	return func(entityManager *ecs.EntityManager) (state int) {
		dt := time.Since(frameTime)
		frameTime = time.Now()
		// Statistics will be updateTime every 2 seconds.
		if time.Since(updateTime) >= time.Second*2 {

			t0 := time.Now()
			em.Get("worst_case_lookup")
			lookupTime := time.Since(t0)

			t1 := time.Now()
			em.FilterByMask(engine.MaskPosition | engine.MaskSize)
			filterTime := time.Since(t1)

			tm.Clear()
			tm.MoveCursor(0, 0)
			_, _ = tm.Println(dash(47))
			_, _ = tm.Println(format("Current Time:", time.Now().Format(time.Stamp)))
			_, _ = tm.Println(dash(47))
			_, _ = tm.Println(format("Runtime Statistics:", ""))
			_, _ = tm.Println(format("GOOS GOARCH", fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)))
			_, _ = tm.Println(format("NumCPU()", fmt.Sprintf("%d", runtime.NumCPU())))
			_, _ = tm.Println(format("NumCgoCall()", fmt.Sprintf("%d", runtime.NumCgoCall())))
			_, _ = tm.Println(format("NumGoroutine()", fmt.Sprintf("%d", runtime.NumGoroutine())))
			_, _ = tm.Println(format("Version()", runtime.Version()))
			_, _ = tm.Println(dash(47))
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			_, _ = tm.Println(format("Memory Statistics:", ""))
			_, _ = tm.Println(format("MemStats Sys", fmt.Sprintf("%d", r.Sys)))
			_, _ = tm.Println(format("Heap Allocation", fmt.Sprintf("%d", r.HeapAlloc)))
			_, _ = tm.Println(format("Heap Idle", fmt.Sprintf("%d", r.HeapIdle)))
			_, _ = tm.Println(format("Head In Use", fmt.Sprintf("%d", r.HeapInuse)))
			_, _ = tm.Println(format("Heap HeapObjects", fmt.Sprintf("%d", r.HeapObjects)))
			_, _ = tm.Println(format("Heap Released", fmt.Sprintf("%d", r.HeapReleased)))
			_, _ = tm.Println(dash(47))
			_, _ = tm.Println(format("Engine Statistics:", ""))
			_, _ = tm.Println(format("Entities:", fmt.Sprintf("%d", len(em.Entities()))))
			_, _ = tm.Println(format("FilterTime:", fmt.Sprintf("%v", filterTime)))
			_, _ = tm.Println(format("FrameTime:", fmt.Sprintf("%v", dt)))
			_, _ = tm.Println(format("LookupTime:", fmt.Sprintf("%v", lookupTime)))
			_, _ = tm.Println(format("Version:", ecs.Version))
			_, _ = tm.Println(dash(47))
			_, _ = tm.Println()
			tm.Flush()
			updateTime = time.Now()
		}
		return ecs.StateEngineContinue
	}
}

func dash(num int) (out string) {
	for i := 0; i < num; i++ {
		out += "-"
	}
	return out
}

func format(key, val string) string {
	return fmt.Sprintf("| %-20s | %-20s |", key, val)
}
