package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetAllPinjaman(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(dbContext).(*gorm.DB)
	if !ok {
		http.Error(w, "no db", http.StatusInternalServerError)
		return
	}

	var allPinjaman []Pinjaman

	if err := db.Find(&allPinjaman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allPinjaman)
}

func CreatePinjaman(w http.ResponseWriter, r *http.Request)  {
	db, ok := r.Context().Value(dbContext).(*gorm.DB)
	if !ok {
		http.Error(w, "no db", http.StatusInternalServerError)
		return
	}

	var pinjaman Pinjaman
	if err := json.NewDecoder(r.Body).Decode(&pinjaman); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pinjaman.IDPinjaman = 0
	if err := db.Create(&pinjaman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pinjaman)
}

func UpdatePinjaman(w http.ResponseWriter, r *http.Request)  {
	db, ok := r.Context().Value(dbContext).(*gorm.DB)
	if !ok {
		http.Error(w, "no db", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pinjaman Pinjaman
	if err := json.NewDecoder(r.Body).Decode(&pinjaman); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pinjaman.IDPinjaman = id
	existPinjaman := Pinjaman{
		IDPinjaman: id,
	}

	if err := db.First(&existPinjaman).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	existPinjaman.IDPeminjam = pinjaman.IDPeminjam
	existPinjaman.PinjamanPokok = pinjaman.PinjamanPokok
	existPinjaman.SukuBunga = pinjaman.SukuBunga
	existPinjaman.TenorPinjaman = pinjaman.TenorPinjaman
	existPinjaman.StatusPinjaman = pinjaman.StatusPinjaman
	existPinjaman.DokumenPinjaman = pinjaman.DokumenPinjaman

	if err := db.Save(&existPinjaman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existPinjaman)
}

func DeletePinjaman(w http.ResponseWriter, r *http.Request)  {
	db, ok := r.Context().Value(dbContext).(*gorm.DB)
	if !ok {
		http.Error(w, "no db", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	existPinjaman := Pinjaman{
		IDPinjaman: id,
	}

	if err = db.First(&existPinjaman).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	
	if err := db.Delete(&existPinjaman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existPinjaman)
}