export GO111MODULE=on
export GOOS=linux

.PHONY: build deploy gomodgen

ifdef profile
  PROFILE=--aws-profile ${profile}
endif

build:
	go mod tidy
	go build -ldflags="-s -w" -o bin/task ./task
	go build -ldflags="-s -w" -o bin/map ./map
	go build -ldflags="-s -w" -o bin/reduce ./reduce

deploy: build
	echo ${PROFILE}
	serverless deploy --verbose ${PROFILE}

dryrun: build
	echo ${PROFILE}
	serverless deploy --noDeploy --verbose ${PROFILE}
