package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/markbates/pkger"
)

func main() {
	godotenv.Load()
	pkger.Stat("github.com/markbates/pkger:/README.md")
	dir := http.FileServer(pkger.Dir("/src"))
	log.Println("Running at http://0.0.0.0:" + os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), dir)
}
