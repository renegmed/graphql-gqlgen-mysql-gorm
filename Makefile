init-project:
	go mod init github.com/renegmed/graphql-gqlgen-mysql-gorm 

init-gqlgen:
	go run github.com/99designs/gqlgen init	 

up:
	docker-compose up --build -d 

