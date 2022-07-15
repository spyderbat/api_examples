all: go python javascript

go: openapi.json go-config.yaml
	openapi-generator-cli generate -g go -c go-config.yaml -i openapi.json -o go

python: openapi.json python-config.yaml
	openapi-generator-cli generate -g python -c python-config.yaml -i openapi.json -o python

javascript: openapi.json javascript-config.yaml
	openapi-generator-cli generate -g javascript -c javascript-config.yaml -i openapi.json -o javascript

openapi.json: 
	wget https://api.spyderbat.com/openapi.json

clean:
	rm -f openapi.json openapitools.json

clean_libraries:
	rm -r go javascript python
