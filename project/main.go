package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/zhangpeihao/gortmp"
)

type VideoMetadata struct {
	gorm.Model
	Title       string  `gorm:"not null"`
	Format      string  `gorm:"not null"`
	Duration    float64 `gorm:"not null"`
	Bitrate     string  `gorm:"not null"`
	Width       int     `gorm:"not null"`
	Height      int     `gorm:"not null"`
	Codec       string  `gorm:"not null"`
	FrameRate   string  `gorm:"not null"`
}

func extractMetadata(videoPath string) (VideoMetadata, error) {
	var metadata VideoMetadata

	client := gortmp.NewClient()
	if err := client.Connect(videoPath); err != nil {
		return metadata, err
	}
	defer client.Close()

	streamInfo, err := client.GetStreamInfo()
	if err != nil {
		return metadata, err
	}

	metadata.Title = videoPath
	metadata.Format = streamInfo.Format
	metadata.Duration = streamInfo.Duration
	metadata.Bitrate = streamInfo.Bitrate
	metadata.Width = streamInfo.Width
	metadata.Height = streamInfo.Height
	metadata.Codec = streamInfo.Codec
	metadata.FrameRate = streamInfo.FrameRate

	return metadata, nil
}

func saveMetadata(db *gorm.DB, metadata VideoMetadata) error {
	return db.Create(&metadata).Error
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	videoPath := r.URL.Query().Get("video")
	if videoPath == "" {
		http.Error(w, "Video path is required", http.StatusBadRequest)
		return
	}

	metadata, err := extractMetadata(videoPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Сохранение метаданных в базу данных
	if err := saveMetadata(db, metadata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metadata)
}

var db *gorm.DB

func main() {
	var err error
	dsn := "root:&7Tgd75T@tcp(127.0.0.1:3306)/video_meta?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	db.AutoMigrate(&VideoMetadata{})

	http.HandleFunc("/metadata", metadataHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}