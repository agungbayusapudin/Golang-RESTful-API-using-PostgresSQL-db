package main

import (
	"database/sql"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type jenis_barang struct {
	ID         int    `json:"id"`
	Nama_jenis string `json:"nama_jenis"`
	Deskripsi  string `json:"deskripsi"`
}

type jenis_barang_2 struct {
	Nama string `json:"nama_jenis"`
}

func CnectionDb() *sql.DB {
	con, err := sql.Open("postgres", "postgres://postgres:sikucing@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return con
}

func GetAllUser(db *sql.DB) ([]jenis_barang, error) {
	data, err := db.Query("SELECT * FROM jenis_barang")
	if err != nil {
		log.Fatal(err)
	}

	var ListData []jenis_barang
	for data.Next() {
		var ListD jenis_barang
		if err := data.Scan(ListD.ID, ListD.Nama_jenis, ListD.Deskripsi); err != nil {
			log.Fatal(err)
		}

		ListData = append(ListData, ListD)
	}

	return ListData, nil

}

func main() {
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", "postgres://postgres:XXXXXXX@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a Gin router
	r := gin.Default()

	// Define API routes
	r.GET("/users", func(c *gin.Context) {
		// Fetch all users from database
		rows, err := db.Query("SELECT id,nama_jenis,deskripsi FROM jenis_barang")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []jenis_barang
		for rows.Next() {
			var user jenis_barang
			if err := rows.Scan(&user.ID, &user.Nama_jenis, &user.Deskripsi); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, user)
		}

		c.JSON(http.StatusOK, users)
	})

	r.GET("/nama", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT nama_jenis FROM jenis_barang")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"eror": err.Error()})
			return
		}

		defer rows.Close()

		var users []jenis_barang_2

		for rows.Next() {
			var userr jenis_barang_2

			if err := rows.Scan(&userr.Nama); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"eror": err.Error()})
			}

			users = append(users, userr)
		}

		ctx.JSON(http.StatusOK, users)
	})

	r.PUT("/ubah", func(ctx *gin.Context) {
		_, err := db.Query("UPDATE nama_jenis = 'sampo' ,deskripsi = 'barang bagus' WHERE id = 2")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"eror": err.Error()})
		}

		ctx.JSON(http.StatusOK, "berhasil merubah data")
	})

	r.Run() // Listen and serve on 0.0.0.0:8080
}
