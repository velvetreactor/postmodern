image-js:
	docker build -t nycdavid/postapoc-js:0.0.1 -f Dockerfile.javascripts .
image-postapoc:
	docker build -t nycdavid/postapoc:0.0.1 . 
image-css:
	docker build -t nycdavid/postapoc-css:0.0.1 -f Dockerfile.stylesheets .
image-nightwatch:
	docker build -t nycdavid/nightwatch:0.0.1 -f Dockerfile.nightwatch .

# Testing
test-integ:
	docker-compose -f docker-compose.test.yml run nightwatch
test-unit:
	docker-compose run postapoc go test -v ./...
