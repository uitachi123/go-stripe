LEVEL?="INFO"
TOKEN?=test
REACT_TOKEN?=token

.PHONY: all test clean build api ui serve test
all: clean test deps build

clean:
	rm -rf ./go-stripe
build:
	go build -a -ldflags 'main.buildTime=$(date)' .
deps:
	npm i
api:
	./go-stripe --logging $(LEVEL) --port "8080" --token $(TOKEN)
ui:
	npm run build && REACT_APP_TOKEN=$(REACT_TOKEN) npm start
serve:
	make -j api ui
test:
	go test ./... -test.v
