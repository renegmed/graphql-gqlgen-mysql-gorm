package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/renegmed/graphql-gqlgen-mysql-gorm/graph"
	"github.com/renegmed/graphql-gqlgen-mysql-gorm/graph/generated"
	"github.com/renegmed/graphql-gqlgen-mysql-gorm/graph/model"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:password123@tcp(localhost:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.LogMode(true)
	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	// db.Exec("CREATE DATABASE IF NOT EXISTS test_db")
	// db.Exec("USE test_db")
	// Migration to create tables for Order and Item schema

	db.AutoMigrate(&model.Order{}, &model.Item{})

	// db.Model(&model.Item{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
