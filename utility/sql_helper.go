package utility

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDbClient() *sql.DB {
	godotenv.Load(".env")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		os.Getenv("host"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbPassword"), os.Getenv("dbName"))
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("Verbindung mit Datenbank kann nicht hergestellt werden: %v", err))
	}
	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(20)
	return conn
}

func GetStringValue(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
func GetBoolValue(s sql.NullBool) bool {
	if s.Valid {
		return s.Bool
	}
	return false
}

func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
