package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"sync"

	"database/sql"
	"net/http"
)

type Filter struct {
	mu sync.Mutex
	db *sql.DB
}

func dbdbdb() {

	var driverName, dataSourceName string
	db, _ := sql.Open(driverName, dataSourceName)
	fmt.Print(db)
}

func main() {
	var trrr []string
	fmt.Printf("%v", trrr)
	_, err := http.Get("http://sina.cn/")
	if err != nil {
		return
	}

	dbdbdb()
	var unBytes []byte
	var contentEncoding string
	Zip(unBytes, contentEncoding)
	return
}

func Zip(unBytes []byte, contentEncoding string) (zipBytes []byte, err error) {

	switch contentEncoding {
	case "":
		zipBytes = unBytes
		break

	case "gzip":
		var bf bytes.Buffer
		gw := gzip.NewWriter(&bf)
		gw.Write(unBytes)
		gw.Close()

		zipBytes = bf.Bytes()

	case "deflate":
		var bf bytes.Buffer
		var gw *flate.Writer
		gw, err = flate.NewWriter(&bf, -1)
		if err != nil {
			return
		}
		gw.Write(unBytes)
		gw.Close()

		zipBytes = bf.Bytes()

	default:
		err = fmt.Errorf("Zip Unkown Content-Encoding: %v", contentEncoding)
	}

	return
}
