init-project:
	go mod init github.com/renegmed/graphql-gqlgen-mysql-gorm 

init-gqlgen:
	# go run github.com/99designs/gqlgen init	 
	go run github.com/99designs/gqlgen generate	 

up:
	docker-compose up --build -d 

