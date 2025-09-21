package models

import "time"

type Skripsi struct {
	ID            uint      `gorm:"primaryKey;column:id" json:"id"`
	NIM           string    `gorm:"column:nim;size:20;not null" json:"nim"`
	NamaMahasiswa string    `gorm:"column:nama_mahasiswa;size:100" json:"nama_mahasiswa"`
	Prodi         string    `gorm:"column:prodi;size:100" json:"prodi"`
	Judul         string    `gorm:"column:judul;size:255" json:"judul"`
	Abstrak       string    `gorm:"column:abstrak;type:text" json:"abstrak"`
	Pembimbing    string    `gorm:"column:pembimbing;size:100" json:"pembimbing"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Skripsi) TableName() string {
	return "tbdataskripsi"
}
