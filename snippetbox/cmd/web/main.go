package main

import (
	"database/sql"
	"net/http"
	"flag"
	"log/slog"
	"os"
	"snippetbox.jordisalazar.net/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct{
	logger *slog.Logger
	snippets *models.SnippetModel
	
}


func main(){
   loggerHandler := slog.NewJSONHandler(os.Stdout, nil) //second arg is the config options of the handler
   logger := slog.New(loggerHandler)

	addr := flag.String("addr",":4000","HTTP network adress")
	dsn := flag.String("dsn","web:pass@/snippetbox?parseTime=true","Mysql data source name")
	flag.Parse()

	db,err := openDB(*dsn)

    app := &application{
		logger:logger,
		snippets:&models.SnippetModel{DB:db},
	 }

	 if err != nil {
	      logger.Error(err.Error())
		  os.Exit(1)
	 }

	 defer db.Close()

	 logger.Info("server started",slog.String("port",*addr))

	 err = http.ListenAndServe(*addr,app.routes())
	 logger.Error(err.Error())
	 os.Exit(1)
}

func openDB(dsn string)(*sql.DB,error){
	db,err:= sql.Open("mysql",dsn)
	if err != nil {
		return nil,err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil,err
	}

	return db,nil
}