package entity

type User struct {
	FirstName string `json:"firstName" gorm:"type:varchar(100);not null"`
	LastName  string `json:"lastName" gorm:"type:varchar(100);not null"`
	Email     string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string `json:"-" gorm:"type:varchar(255);not null"` // không lưu trữ password dưới dạng văn bản
}
