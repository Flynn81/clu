help:
	$(info CLU - help)
	$(info tasks:)
	$(info )
	$(info test - run unit tests)
	$(info benchmark - benchmark the code via tests)
	$(info profile - profile the application while in use against a set test suite)
	$(info bin - build a binary)
	$(info )
test:
	$(info running tests)
	go test ./...
benchmark:
	$(info FILL IN BENCHMARK)
profile:
	$(info FILL IN PROFILING)
bin:
	go build -o clu
