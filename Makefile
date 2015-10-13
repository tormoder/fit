.PHONY: \
	all \
	test \
	testrace \
	bench \
	fmt \
	vet \
	lint \
	style \
	fitgenprofile \
	gofuzz \
	gofuzzclean \
	clean \
	profcpu \
	profmem \
	profobj \
	mdgen \
	get \

all: get fmt vet test testrace

test:
	go test -v ./...

testrace: test
	go test -v -race ./...

bench:
	go test -v -run NONE -bench .

fmt:
	gofmt -l -s . dyncrc16 cmd/fitgen cmd/fitgen/internal/profile

vet:
	go tool vet -shift=false .

lint:
	golint . | \
		grep -v types.go | \
		grep -v types_man.go | \
		grep -v types_string.go | \
		grep -v messages.go | \
		grep -v FileId
	golint dyncrc16
	golint cmd/fitgen

style: fmt vet lint

fitgen:
	go install github.com/tormoder/fit/cmd/fitgen

gofuzz:
	go get -u github.com/dvyukov/go-fuzz/go-fuzz
	go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
	go-fuzz-build github.com/tormoder/fit

gofuzzclean: gofuzz
	rm -rf workdir/
	mkdir -p workdir/corpus
	find testdata -name \*.fit -exec cp {} workdir/corpus/ \;

clean:
	rm -rf profile/*.csv
	rm -rf profile/*.xlsx
	rm -f fit-fuzz.zip
	rm -f *.prof
	rm -f *.test

profcpu:
	go test -run=NONE -bench=ActivitySmall -cpuprofile cpu.prof
	go tool pprof fit.test cpu.prof

profmem:
	go test -run=NONE -bench=ActivitySmall -memprofile allocmem.prof
	go tool pprof -alloc_space fit.test allocmem.prof

profobj:
	go test -run=NONE -bench=ActivitySmall -memprofile allocobj.prof
	go tool pprof -alloc_objects fit.test allocobj.prof

mdgen:
	godoc2md github.com/tormoder/fit Fit Header CheckIntegrity > MainApiReference.md

get:
	go get -v ./...
