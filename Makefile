.SILENT:

generate_clients:
	openapi-generator generate -i openapi.yml -g typescript-axios -o typescript-client
	openapi-generator generate -i openapi.yaml -g go -o go-client
