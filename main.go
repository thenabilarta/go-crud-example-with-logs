package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"crudgo/config"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Type  string `json:"type"`
}

func main() {
	var logFile *os.File

	config.LoadConfig(logFile)

	defer logFile.Close()

	log.Infof("Running on version %s, build date : %s", "1", "10")

	db, err := sql.Open("mysql", "root:root@/gocrud")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/product", ProductHandler).Methods("GET")
	r.HandleFunc("/product", CreateProductHandler).Methods("POST")
	r.HandleFunc("/product/{id}/update", UpdateProductHandler).Methods("POST")
	r.HandleFunc("/product/{id}/delete", DeleteProductHandler).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:1905",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	serverError := srv.ListenAndServe()

	if serverError != nil {
		log.Errorf("terminated %s", err)
		os.Exit(1)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("StubService.ListStub Request : %+v", "test")

	values := map[string]string{"username": "username", "password": "password"}

	jsonValue, _ := json.Marshal(values)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("StubService.ListStub Request : %+v", "test")

	// values := map[string]string{"username": "username", "password": "password"}

	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name, price, type FROM product")

	if err != nil {
		panic(err)
	}

	var products []Product

	for rows.Next() {
		var product Product

		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Type)

		products = append(products, product)
	}

	body, err := json.Marshal(&products)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "close")

	w.WriteHeader(http.StatusOK)

	w.Write(body)

	// jsonValue, _ := json.Marshal(values)
	// w.WriteHeader(http.StatusOK)
	// w.Write(jsonValue)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("StubService.ListStub Request : %+v", "test")

	var productRequest Product

	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(r.Body)

	if len(body) > 0 {
		if err = json.Unmarshal(body, &productRequest); err != nil {
			panic(err)
		}
	}

	name := []byte(productRequest.Name)
	price := (productRequest.Price)
	product_type := []byte(productRequest.Type)

	sqlStatement := `
		INSERT INTO product (id, name, price, type)
		VALUES (?, ?, ?, ?)
	`

	row_id := 0

	_, err = db.Exec(sqlStatement, 0, name, price, product_type)

	if err != nil {
		panic(err)
	}

	values := map[string]string{"id": strconv.Itoa(row_id), "password": "password"}

	jsonValue, _ := json.Marshal(values)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("StubService.ListStub Request : %+v", "test")

	var productRequest Product

	product_id := mux.Vars(r)["id"]

	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(r.Body)

	if len(body) > 0 {
		if err = json.Unmarshal(body, &productRequest); err != nil {
			panic(err)
		}
	}

	name := []byte(productRequest.Name)
	price := (productRequest.Price)
	product_type := []byte(productRequest.Type)

	sqlStatement :=
		"UPDATE product SET name = ?, price = ?, type = ? WHERE id = ?"

	row_id := 0

	_, err = db.Query(sqlStatement, name, price, product_type, product_id)

	if err != nil {
		panic(err)
	}

	values := map[string]string{"id": strconv.Itoa(row_id), "password": "password"}

	jsonValue, _ := json.Marshal(values)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("StubService.ListStub Request : %+v", "test")

	product_id := mux.Vars(r)["id"]

	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err)
	}

	sqlStatement :=
		"DELETE FROM product WHERE id = ?"

	row_id := 0

	_, err = db.Query(sqlStatement, product_id)

	if err != nil {
		panic(err)
	}

	values := map[string]string{"id": strconv.Itoa(row_id), "password": "password"}

	jsonValue, _ := json.Marshal(values)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
}
