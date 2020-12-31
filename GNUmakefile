TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
default: build

build:
	go build .

test:
	 go test $(TEST) -v $(TESTARGS)

testacc:
	GODEBUG=x509ignoreCN=0 \
	KAFKA_BOOTSTRAP_SERVERS=localhost:9092 \
	TF_LOG=DEBUG \
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 9m -count=1

.PHONY: build test testacc
