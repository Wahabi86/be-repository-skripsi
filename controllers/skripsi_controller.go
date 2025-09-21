package controllers

import (
	"net/http"
	"pratikum-repository-skripsi/config"
	"pratikum-repository-skripsi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Untuk Create Data Skripsi Mahasiswa
func CreateSkripsi(c *gin.Context) {
	var input models.Skripsi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// memastikan ID tetap auto increment
	input.ID = 0
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// Untuk Melihat Semua Data Skripsi Mahasiwa
func GetAllSkripsi(c *gin.Context) {
	var data []models.Skripsi
	if err := config.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// Untuk Melihat Data Skripsi Mahasiswa Berdasarkan ID
func GetSkripsiByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var skripsi models.Skripsi
	if err := config.DB.First(&skripsi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}
	c.JSON(http.StatusOK, skripsi)
}

// Untuk Update Data Skripsi Mahasiswa
func PatchSkripsi(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return
    }

    var skripsi models.Skripsi
    if err := config.DB.First(&skripsi, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
        return
    }

    // pakai map[string]interface{} untuk hanya field yang dikirim yang akan diupdate
    var input map[string]interface{}
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&skripsi).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, skripsi)
}

// Untuk Delete Data Skripsi
func DeleteSkripsi(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err := config.DB.Delete(&models.Skripsi{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
