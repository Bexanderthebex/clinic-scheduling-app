package models

type Hospital struct {
	Id      string `gorm:"primaryKey"`
	Name    string
	City    string
	Address string
	Lat     float64
	Long    float64
	Logo    string
}
