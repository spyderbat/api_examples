all: go python javascript rust

go: openapi.json go-config.yaml
	openapi-generator-cli generate -g go -c go-config.yaml -i openapi.json -o go

python: openapi.json python-config.yaml
	openapi-generator-cli generate -g python -c python-config.yaml -i openapi.json -o python

javascript: openapi.json javascript-config.yaml
	openapi-generator-cli generate -g javascript -c javascript-config.yaml -i openapi.json -o javascript

rust: openapi.json rust-config.yaml
	openapi-generator-cli generate -g rust -c rust-config.yaml -i openapi.json -o rust

openapi.json: 
	wget https://api.spyderbat.com/openapi.json

clean:
	rm -f openapi.json openapitools.json

clean_libraries:
	rm -rf go javascript python rust/src
