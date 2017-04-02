FIT_PKGS 	:= $(shell go list ./... | grep -v /vendor/)
FIT_FILES	:= $(shell find . -name '*.go' -not -path "*vendor*")
FIT_DIRS 	:= $(shell find . -type d -not -path "*vendor*" -not -path "./.git*" -not -path "*testdata*")

FIT_PKG_PATH 	:= github.com/tormoder/fit
FITGEN_PKG_PATH := $(FIT_PKG_PATH)/cmd/fitgen
GOFUZZ_PKG_PATH := github.com/dvyukov/go-fuzz
LATLONG_PKG_PATH:= github.com/bradfitz/latlong

.PHONY: all
all: deps build test testrace check

.PHONY: deps
deps:
	@echo "go get:"
	go get $(LATLONG_PKG_PATH)

.PHONY: build
build:
	@echo "go build:"
	go build -v -i $(FIT_PKGS)

.PHONY: test
test:
	@echo "go test:"
	go test -v -cpu=2 $(FIT_PKGS)

.PHONY: testrace
testrace:
	@echo "go test -race:"
	go test -v -cpu=1,2,4 -race $(FIT_PKGS)

.PHONY: bench
bench:
	go test -v -run=$$$$ -bench=. $(FIT_PKGS)

.PHONY: fitgen
fitgen:
	go install $(FITGEN_PKG_PATH)

.PHONY: gofuzz
gofuzz:
	go get -u $(GOFUZZ_PKG_PATH)/go-fuzz
	go get -u $(GOFUZZ_PKG_PATH)/go-fuzz-build
	go-fuzz-build $(FIT_PKG_PATH)

.PHONY: gofuzzclean
gofuzzclean: gofuzz
	rm -rf workdir/
	mkdir -p workdir/corpus
	find testdata -name \*.fit -exec cp {} workdir/corpus/ \;

.PHONY: clean
clean:
	go clean -i ./...
	rm -f fit-fuzz.zip
	find . -name '*.prof' -type f -exec rm -f {} \;
	find . -name '*.test' -type f -exec rm -f {} \;
	find . -name '*.current' -type f -exec rm -f {} \;
	find . -name '*.current.gz' -type f -exec rm -f {} \;

.PHONY: gcoprofile 
gcoprofile:
	git checkout types.go messages.go profile.go

.PHONY: profcpu
profcpu:
	go test -run=$$$$ -bench=ActivitySmall -cpuprofile=cpu.prof
	go tool pprof fit.test cpu.prof

.PHONY: profmem
profmem:
	go test -run=$$$$-bench=ActivitySmall -memprofile=allocmem.prof
	go tool pprof -alloc_space fit.test allocmem.prof

.PHONY: profobj
profobj:
	go test -run=NONE -bench=ActivitySmall -memprofile=allocobj.prof
	go tool pprof -alloc_objects fit.test allocobj.prof

.PHONY: mdgen
mdgen:
	godoc2md $(FIT_PKG_PATH) Fit Header CheckIntegrity > MainApiReference.md

.PHONY: getgvt
getgvt:
	go get -u github.com/FiloSottile/gvt

.PHONY: getchecktools
getchecktools:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint
	go get -u github.com/jgautheron/goconst/cmd/goconst
	go get -u github.com/kisielk/errcheck
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/mdempsky/unconvert
	go get -u honnef.co/go/unused/cmd/unused
	go get -u honnef.co/go/simple/cmd/gosimple
	go get -u github.com/mvdan/interfacer/cmd/interfacer
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u honnef.co/go/staticcheck/cmd/staticcheck

.PHONY: check
check:
	@echo "check (basic)":
	@echo "gofmt (simplify)"
	@gofmt -s -l .
	@echo "go vet"
	@go vet ./...

.PHONY: checkfull
checkfull:
	@echo "check (full):"
	@echo "gofmt (simplify)"
	@! gofmt -s -l $(FIT_FILES) | grep -vF 'No Exceptions'
	@echo "goimports"
	@! goimports -l $(FIT_FILES) | grep -vF 'No Exceptions'
	@echo "vet"
	@! go tool vet $(FIT_DIRS) 2>&1 | \
		grep -vF 'vendor/'
	@echo "vet --shadow"
	@! go tool vet --shadow $(FIT_DIRS) 2>&1 | grep -vF 'vendor/'
	@echo "golint"
	@for pkg in $(FIT_PKGS); do \
		! golint $$pkg | \
		grep -vE '(FileId|SegmentId|messages.go|types.*.\go|fitgen/internal|cmd/stringer)' ; \
	done
	@echo "goconst"
	@for dir in $(FIT_DIRS); do \
		goconst $$dir ; \
	done
	@echo "errcheck"
	@errcheck -ignore 'bytes:Write*,archive/zip:Close,io:Close,Write' $(FIT_PKGS)
	@echo "ineffassign"
	@for dir in $(FIT_DIRS); do \
		ineffassign -n $$dir ; \
	done
	@echo "unconvert"
	@! unconvert $(FIT_PKGS) | grep -vF 'messages.go'
	@echo "unused"
	@unused $(FIT_PKGS)
	@echo "gosimple"
	@for pkg in $(FIT_PKGS); do \
		gosimple $$pkg ; \
	done
	@echo "interfacer"
	@interfacer $(FIT_PKGS)
	@echo "misspell"
	@ ! misspell ./**/* | grep -vE '(messages.go|/vendor/|profile/testdata)'
	@echo "staticcheck"
	@staticcheck $(GORUMS_PKGS)
