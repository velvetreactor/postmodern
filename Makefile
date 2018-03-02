image-js:
	docker build -t nycdavid/postapoc-js:0.0.1 -f Dockerfile.javascripts .
image-postapoc:
	docker build -t nycdavid/postapoc:0.0.1 .
image-css:
	docker build -t nycdavid/postapoc-css:0.0.1 -f Dockerfile.stylesheets .
