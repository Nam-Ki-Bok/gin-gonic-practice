package post

// Post - 게시글 관리 테이블
type Post struct {
	Idx     uint   `gorm:"primary_key; auto_increment:true" json:"idx"`
	Email   string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Content string `gorm:"type:varchar(100);not null" json:"content"`
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
}
