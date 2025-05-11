
# TODO: download actual Cypher.g4 grammar file
# https://s3.amazonaws.com/artifacts.opencypher.org/M23/Cypher.g4

gen: Cypher.g4 # generata
	# docker run --rm -v $(shell pwd):/build -w /build --user $(shell id -u):$(shell id -g) leodido/antlr:4.7 -Dlanguage=Go -o pkg/cypher_parser Cypher.g4
	docker run --rm -v $(shell pwd):/build -w /build --user $(shell id -u):$(shell id -g) leodido/antlr:4.7 -Dlanguage=Go -o parser Cypher.g4



