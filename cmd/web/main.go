package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"snippet-box.mms.io/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	models   *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	dsn := flag.String("dsn", "web:@/snippet-box?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog,
		errorLog,
		&models.SnippetModel{db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err = server.ListenAndServe()
	infoLog.Printf("Started server on port %s\n", *addr)
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
