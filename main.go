package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pinjaman struct {
	IDPinjaman int64 `json:"idPinjaman" gorm:"primaryKey; autoIncrement"`
	IDPeminjam int64 `json:"idPeminjam"`
	PinjamanPokok int `json:"pinjamanPokok"`
	SukuBunga int `json:"sukuBunga"`
	TenorPinjaman int `json:"tenorPinjaman"`
	StatusPinjaman string `json:"statusPinjaman"`
	DokumenPinjaman bool `json:"dokumenPinjaman"`
	TglDibuat time.Time `json:"tglDibuat" `
	TglDiupdate time.Time `json:"tglDiupdate" `
}

type key int 

const (
	dbContext key = iota
)

func DBMiddleware(next http.Handler, db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), dbContext, db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	dsn := "host=localhost user=postgres password=root dbname=pengajuan-pinjaman port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Pinjaman{})
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("/pinjaman", GetAllPinjaman)
	router.HandleFunc("/pinjaman/create", CreatePinjaman)
	router.HandleFunc("/pinjaman/update", UpdatePinjaman)
	router.HandleFunc("/pinjaman/delete", DeletePinjaman)

	routerMiddleware := DBMiddleware(router, db)

	fmt.Println("Server is started at localhost:4000")

	http.ListenAndServe("localhost:4000", routerMiddleware)

}