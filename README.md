# cypher

_created by Austin Poor_

A Cypher query language parser for Go.

## Notes

* [Cypher Manual](https://neo4j.com/docs/cypher-manual/current)
* [openCypher Resources](https://opencypher.org/resources)
* [Cypher Railroad Diagram](https://s3.amazonaws.com/artifacts.opencypher.org/M18/railroad/Cypher.html#SingleQuery)
* [Reading Railroad Diagrams](https://www.ibm.com/docs/en/integration-bus/10.0?topic=diagrams-how-read-railroad)

## Dev Notes

Cypher grammar file obtained from the [openCypher website](https://opencypher.org/resources).

Code generated with [Antlr](https://github.com/antlr/antlr4) (version 4.7 [docker image](https://hub.docker.com/r/leodido/antlr/)) via:

```bash
$ antlr -Dlanguage=Go -o parser Cypher.g4
# in docker
$ docker run --rm -v $(shell pwd):/build -w /build --user $(shell id -u):$(shell id -g) leodido/antlr:4.7 -Dlanguage=Go -o parser Cypher.g4

```
