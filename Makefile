GO		:= go
FIT_PKGS 	:= $(shell $(GO) list ./... | grep -v /vendor/)
FIT_FILES	:= $(shell find . -name '*.go' -not -path "*vendor*")
FIT_DIRS 	:= $(shell find . -type f -not -path "*vendor*" -not -path "./.git*" -not -path "*testdata*" -name "*.go" -printf "%h\n" | sort -u)

FIT_PKG_PATH 	:= github.com/tormoder/fit
FITGEN_PKG_PATH := $(FIT_PKG_PATH)/cmd/fitgen

GOFUZZ_PKG_PATH := github.com/dvyukov/go-fuzz
LATLONG_PKG_PATH:= github.com/bradfitz/latlong
UTTER_PKG_PATH	:= github.com/kortschak/utter
XXHASH_PKG_PATH := github.com/cespare/xxhash

DECODE_BENCH_NAME := DecodeActivity$$/Small
DECODE_BENCH_TIME := 5s

CHECK_TOOLS :=	golang.org/x/tools/cmd/goimports \
		golang.org/x/lint/golint \
		github.com/jgautheron/goconst/cmd/goconst \
		github.com/kisielk/errcheck \
		github.com/gordonklaus/ineffassign \
		github.com/mdempsky/unconvert \
		mvdan.cc/interfacer \
		github.com/client9/misspell/cmd/misspell \
		honnef.co/go/tools/cmd/staticcheck/ \
		golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow \

.PHONY: all
all: deps testdeps build test testrace checkfull

.PHONY: deps
deps:
	@echo "$(GO) get:"
	$(GO) get -u $(LATLONG_PKG_PATH)

.PHONY: testdeps
testdeps:
	@echo "$(GO) get -u:"
	$(GO) get -u $(UTTER_PKG_PATH)
	$(GO) get -u $(XXHASH_PKG_PATH)

.PHONY: build
build:
	@echo "$(GO) build:"
	$(GO) build -v -i $(FIT_PKGS)

.PHONY: test
test:
	@echo "$(GO) test:"
	$(GO) test -v -cpu=2 $(FIT_PKGS)

.PHONY: testrace
testrace:
	@echo "$(GO) test -race:"
	$(GO) test -v -cpu=1,2,4 -race $(FIT_PKGS)

.PHONY: bench
bench:
	$(GO) test -v -run=^$$ -bench=. $(FIT_PKGS) -benchtime=5s

.PHONY: fitgen
fitgen:
	$(GO) install $(FITGEN_PKG_PATH)

.PHONY: gofuzz
gofuzz:
	$(GO) get -u $(GOFUZZ_PKG_PATH)/go-fuzz
	$(GO) get -u $(GOFUZZ_PKG_PATH)/go-fuzz-build
	go-fuzz-build $(FIT_PKG_PATH)

.PHONY: gofuzzclean
gofuzzclean: gofuzz
	rm -rf workdir/
	mkdir -p workdir/corpus
	find testdata -name \*.fit -exec cp {} workdir/corpus/ \;

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f fit-fuzz.zip
	find . -name '*.prof' -type f -exec rm -f {} \;
	find . -name '*.test' -type f -exec rm -f {} \;
	find . -name '*.current' -type f -exec rm -f {} \;
	find . -name '*.current.gz' -type f -exec rm -f {} \;

.PHONY: gcoprofile 
gcoprofile:
	git checkout types.go messages.go profile.go types_string.go

.PHONY: profcpu
profcpu:
	$(GO) test -run=^$$ -cpuprofile=cpu.prof -bench=$(DECODE_BENCH_NAME) -benchtime=$(DECODE_BENCH_TIME)
	$(GO) tool pprof fit.test cpu.prof

.PHONY: profmem
profmem:
	$(GO) test -run^$$ =-memprofile=allocmem.prof -bench=$(DECODE_BENCH_NAME) -benchtime=$(DECODE_BENCH_TIME)
	$(GO) tool pprof -alloc_space fit.test allocmem.prof

.PHONY: profobj
profobj:
	$(GO) test -run=^$$ -memprofile=allocobj.prof -bench=$(DECODE_BENCH_NAME) -benchtime=$(DECODE_BENCH_TIME)
	$(GO) tool pprof -alloc_objects fit.test allocobj.prof

.PHONY: mdgen
mdgen:
	godoc2md $(FIT_PKG_PATH) Fit Header CheckIntegrity > MainApiReference.md

.PHONY: getdep
getdep:
	$(GO) get -u github.com/golang/dep

.PHONY: getchecktools
getchecktools:
	@echo "$(GO) get'ing (-u) tools for static analysis"
	@$(GO) get -u $(CHECK_TOOLS)

.PHONY: check
check:
	@echo "check (basic)":
	@echo "gofmt (simplify)"
	@gofmt -s -l .
	@echo "$(GO) vet"
	@$(GO) vet ./...

.PHONY: checkfull
checkfull: getchecktools
	@echo "check (full):"
	@echo "gofmt (simplify)"
	@! gofmt -s -l $(FIT_FILES) | grep -vF 'No Exceptions'
	@echo "goimports"
	@! goimports -l $(FIT_FILES) | grep -vF 'No Exceptions'
	@echo "vet"
	@! $(GO) vet $(FIT_DIRS) 2>&1 | \
		grep -vF 'vendor/'
	@echo "vet --shadow"
	@! $(GO) vet -vettool=$(which shadow) $(FIT_DIRS) 2>&1 | grep -vF 'vendor/'
	@echo "golint"
	@! golint $(FIT_PKGS) | grep -vE '(FileId|SegmentId|messages.go|types.*.\go|fitgen/internal|cmd/stringer)'
	@echo "goconst"
	@for dir in $(FIT_DIRS); do \
		goconst $$dir ; \
	done
	@echo "errcheck"
	@errcheck -ignore 'fmt:Fprinf*,bytes:Write*,archive/zip:Close,io:Close,Write' $(FIT_PKGS)
	@echo "ineffassign"
	@for dir in $(FIT_DIRS); do \
		ineffassign -n $$dir ; \
	done
	@echo "unconvert"
	@! unconvert $(FIT_PKGS) | grep -vF 'messages.go'
	@echo "interfacer"
	@interfacer $(FIT_PKGS)
	@echo "misspell"
	@! misspell ./**/* | grep -vE '(messages.go|/vendor/|profile/testdata)'
	@echo "staticcheck"
	@! staticcheck $(FIT_PKGS) | grep -vE '(tdoStderrLogger)'
