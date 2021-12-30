OUTPUT_FILE=calcul.out

all: test build

build:
	go build -o ${OUTPUT_FILE} main.go

test:
	go test ./...

run: build
	./${OUTPUT_FILE}

clean:
	go clean
	rm ${OUTPUT_FILE}

http: all
ifneq (,$(port))
	echo port is $(port)
	./${OUTPUT_FILE} --web-server --port $(port)
else
	./${OUTPUT_FILE} --web-server
endif

cli: all
	./${OUTPUT_FILE}