package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func Connect() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), "postgres://postgres:salam@localhost:5432/sces?pool_max_conns=10")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

const (
	listQuery = "select array_to_json(array_agg(row_to_json(t))) from (select * from %s) t"
	getQuery  = "select row_to_json(%s) from %[1]s where uid=$1"
	putQuery  = "insert into %s values (%s) returning uid"
)

var Ctx = context.Background()

func List(table string, conn *pgxpool.Pool) string {
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

func Retrieve(uid string, table string, conn *pgxpool.Pool) string {
	query := fmt.Sprintf(getQuery, table)
	row := conn.QueryRow(Ctx, query, uid)

	var res string
	if err := row.Scan(&res); err != nil {
		return JSONError(FirstWords(err.Error(), 3))
	}

	return res
}

func Shchema(table string, obj DatabaseModel, conn pgxpool.Conn) error {

	return nil
}
