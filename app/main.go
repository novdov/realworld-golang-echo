package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/novdov/realworld-golang-echo/article"
	"github.com/novdov/realworld-golang-echo/router"
	"github.com/novdov/realworld-golang-echo/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURIFormat = "mongodb+srv://novdov:%s@cluster0-bubht.mongodb.net/<dbname>?retryWrites=true&w=majority"
)

func main() {
	LoadEnv(".env")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI()))
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}

	defer client.Disconnect(context.TODO())

	db := client.Database("real-world")

	userRepo := user.NewUserRepository(db, "user")
	userService := user.NewUserService(userRepo)
	articleRepo := article.NewArticleRepository(db, "article")
	articleService := article.NewArticleService(articleRepo)

	r := router.NewRouter()
	g := r.Group("/api")

	userHandler := user.NewHandler(userService)
	userHandler.Register(g)
	articleHandler := article.NewHandler(articleService, userService)
	articleHandler.Register(g)

	r.Logger.Fatal(r.Start(":8000"))
}

func mongoURI() string {
	return fmt.Sprintf(mongoURIFormat, os.Getenv("ATALS_PASSWORD"))
}

func getEnvFiles(rootDir string) []string {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func LoadEnv(envRootDir string) {
	envFiles := getEnvFiles(envRootDir)
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Fatal("Failed to load env files")
	}
}
