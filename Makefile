CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
MOCKVER=v3.3.6
LINTVER=v1.52.0
MOCKBIN=${BINDIR}/minimock_${GOVER}_${MOCKVER}
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=github.com/Svoevolin/workshop_1_bot/cmd/bot

all: format build test lint

build: bindir
	go build -o ${BINDIR}/bot ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

generate: install-minimock
	cd ${CURDIR}/internal/model/messages/ && ${MOCKBIN} -o ${CURDIR}/internal/mocks/messages/ -s "_mock.go"

install-minimock: bindir
	test -f ${MOCKBIN} || \
		(GOBIN=${BINDIR} go install github.com/gojuno/minimock/v3/cmd/minimock@${MOCKVER} && \
		mv ${BINDIR}/minimock ${MOCKBIN})

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

docker-run:
	sudo docker compose up
