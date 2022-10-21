ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build-single-container:
	@echo $(ROOT_DIR)
	docker build --progress=plain --rm -t "notion-task-integrator:1.0.0" -f "$(ROOT_DIR)/Dockerfile" "$(ROOT_DIR)"

run-server:
	@echo $(ROOT_DIR)
	cd "$(ROOT_DIR)/app-backend" && go run main.go