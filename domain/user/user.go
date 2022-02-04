package user

// User - 유저 관리 테이블
type User struct {
	Idx      uint   `gorm:"primary_key; auto_increment:true" json:"idx"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
}
