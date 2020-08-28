package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func Connect() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

const (
	listQuery   = "select array_to_json(array_agg(row_to_json(t))) from (select * from %s) t"
	getQuery    = "select row_to_json(%s) from (select %s from %[1]s) where uid=$1"
	deleteQuery = `remove from %s where uid=$1`
)

const (
	SukukOrderInsertQuery = `
		insert into sukuk_order 
			(firm_id, sukuk, price, quantity, side, order_type) 
			values ($1, $2, $3, $4, $5, $6) 
			returning uid`
	salamOrderInsertQuery = "insert into salam_order values ($1, $2, $3, $4, $5, $6, $7) returning uid"
	sukukTransInsertQuery = "insert into sukuk_transaction values ($1, $2, $3, $4, $5) returning uid"
	salamTransInsertQuery = "insert into salam_transaction values ($1, $2, $3, $4, $5, $6) returning uid"
)

var Ctx = context.Background()

func list(table string, conn *pgxpool.Pool) string {
	rows, err := conn.Query(Ctx, fmt.Sprintf(listQuery, table))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var b bytes.Buffer

	for rows.Next() {
		var str string
		_ = rows.Scan(&str)
		b.WriteString(str)
	}

	return b.String()
}

func retrieve(uid string, table string, fields string, conn *pgxpool.Pool) (string, error) {
	query := fmt.Sprintf(getQuery, table, fields)
	row := conn.QueryRow(Ctx, query, uid)

	var res string
	if err := row.Scan(&res); err != nil {
		return "", err
	}

	return res, nil
}

type UIDReturn struct {
	uid string
}

func insert(query string, conn *pgxpool.Pool, values ...interface{}) (string, error) {
	row := conn.QueryRow(Ctx, query, values...)

	var res string
	if err := row.Scan(&res); err != nil {
		return "", err
	}

	return res, nil
}

func remove(table string, uid string, conn *pgxpool.Pool) error {
	if _, err := conn.Exec(Ctx, fmt.Sprintf(deleteQuery, table), uid); err != nil {
		return err
	}

	return nil
}

type DBSukukOrderService struct {
	Conn *pgxpool.Pool
}

// TODO: Add error switch for SQL error codes

func (s *DBSukukOrderService) Get(uid string) (string, error) {
	return retrieve(uid, DB_SUKUKORDER, "firm_id, sukuk, price, quantity, side, order_type", s.Conn)
}

func (s *DBSukukOrderService) Put(values ...interface{}) (string, error) {

	if res, err := insert(SukukOrderInsertQuery, s.Conn, values); err != nil {
		return "", JSONError{Msg: err.Error()}
	} else {
		return res, nil
	}
}

func (s *DBSukukOrderService) Delete(uid string) error {
	if err := remove(uid, DB_SUKUKORDER, s.Conn); err != nil {
		return JSONError{Msg: "delete failed"}
	}

	return nil
}
