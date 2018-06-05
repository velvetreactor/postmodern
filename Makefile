image-js:
	docker build -t nycdavid/postapoc-js:0.0.1 -f Dockerfile.javascripts .
image-postapoc:
	docker build -t nycdavid/postapoc:0.0.1 .
image-css:
	docker build -t nycdavid/postapoc-css:0.0.1 -f Dockerfile.stylesheets .
image-nightwatch:
	docker build -t nycdavid/nightwatch:0.0.1 -f Dockerfile.nightwatch .

# Package management
yarn-add:
	docker-compose run \
	-v $(shell pwd)/docker.javascripts.src:/app \
	javascripts \
	yarn add ${PACKAGE}

dep-ensure:
	docker-compose run \
	-v $(shell pwd)/docker.postapoc.src:/go/src/github.com/velvetreactor/postapocalypse/ \
	postapoc \
	dep ensure

# Testing
test-integ:
	docker-compose stop postgres && \
	docker-compose rm -f postgres && \
	docker-compose \
	-f docker-compose.test.yml \
	up \
	--force-recreate \
	--abort-on-container-exit \
	| grep -E "(nightwatch|postapoc)"

func?=
ifdef func
	runCmd=-run $(func)
endif
test-unit:
	docker-compose stop postgres && \
	docker-compose rm -f postgres && \
	docker-compose run --rm postapoc go test -v $(runCmd) ./...

test-js:
	docker-compose run --rm javascripts npm test
