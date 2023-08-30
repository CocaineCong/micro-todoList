DIR = $(shell pwd)/app

CONFIG_PATH = $(shell pwd)/config
IDL_PATH = $(shell pwd)/idl

SERVICES := gateway user task
service = $(word 1, $@)

node = 0

BIN = $(shell pwd)/bin

.PHONY: proto
proto:
	@for file in $(IDL_PATH)/*.proto; do \
		protoc -I $(IDL_PATH) $$file --micro_out=$(IDL_PATH)/pb --go_out=$(IDL_PATH)/pb; \
	done

.PHONY: proto2
proto2:
	@for file in $(shell find $(IDL_PATH)/pb/* -type f); do \
		protoc-go-inject-tag -input=$$file; \
	done


.PHONY: $(SERVICES)
$(SERVICES):
	go build -o $(BIN)/$(service) $(DIR)/$(service)/cmd
	$(BIN)/$(service) -config $(CONFIG_PATH) -srvnum=$(node)

.PHONY: env-up
env-up:
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down

.PHONY: run
run: 
	make -j3 run-all;

.PHONY: run-all
run-all: $(addprefix run-, $(SERVICES))

.PHONY: run-%
run-%:
	go run $(DIR)/$*/cmd/main.go;