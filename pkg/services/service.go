package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/datlin-org/sigzag/pkg/helpers"
	//"github.com/datlin-org/sigzag/pkg/models/postgresql"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
	//SigZag        *postgresql.SigZagModel
}

type Server struct {
	//SigZag *postgresql.SigZagModel
	UnimplementedPipelineServer
}

func (s Server) CreatePipeline(ctx context.Context, in *Config) (*Service, error) {

	serviceInfo := Service{
		ServiceID:    helpers.Sha256Digest(256),
		PipelineType: in.PipelineType,
		PipelineId:   in.PipelineID,
	}
	return &serviceInfo, nil
}

// LogTransaction logs transaction
func (s Server) LogTransaction(ctx context.Context, in *Transaction) (*Log, error) {
	var tr []*Transaction
	// need to persist and the retrieve from persistence layer (SQLite).
	tr = append(tr, in)
	transactionLog := Log{
		LogID:       helpers.Sha256Digest(256),
		Transaction: tr,
	}
	return &transactionLog, nil
}

func RunService(addr string) error {
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	///dbConn := os.Getenv("DB_CONN")
	//if dbConn == "" {
	//	panic("database connection env var not set")
	//}
	// Define a new command-line flag for the MySQL DSN string. Change from public to private string when IP address has been added to the digital ocean platform
	//dsn := flag.String("dsn", dbConn, "Postgresql data source name")

	//db, err := openDB(*dsn)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		//SigZag:   &postgresql.SigZagModel{DB: db},
	}

	server := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	go grpcService(*app)
	err := server.ListenAndServe()
	if err != nil {
		println(err)
	}
	infoLog.Printf("Running server on: %s\n", addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
	return nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	fmt.Printf("opening Db connection")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func grpcService(app application) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	listening, err := net.Listen("tcp", ":5885")
	fmt.Println("Starting GRPC")
	if err != nil {
		errorLog.Fatal(err)
	}
	s := grpc.NewServer()
	RegisterPipelineServer(s, &Server{})
	_ = s.Serve(listening)
}
