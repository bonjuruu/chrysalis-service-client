.SILENT:

generate_clients:
	openapi-generator generate -i openapi.yml -g typescript-axios -o typescript-client \
		--git-user-id bonjuruu \
    	--git-repo-id chrysalis-service-client

	openapi-generator generate -i openapi.yml -g go -o go-client \
		--git-user-id bonjuruu \
    	--git-repo-id chrysalis-service-client/go-client \
