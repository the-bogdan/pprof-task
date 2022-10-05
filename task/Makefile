all: bench

.PHONY: bench
bench:
	@go test -bench .  -benchmem

.PHONY: bench-profile
bench-profile:
	@go test -bench BenchmarkFast -cpuprofile pprof-profiles/cpu.out -memprofile pprof-profiles/mem.out

.PHONY: cpu-pprof
cpu-pprof: bench-profile
	@go tool pprof pprof-task.test pprof-profiles/cpu.out

.PHONY: mem-pprof
mem-pprof: bench-profile
	@go tool pprof pprof-task.test pprof-profiles/mem.out