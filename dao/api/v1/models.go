package v1

type Student struct {
	StudentId      string `json:"id" gorm:"column:studentId"`
	StudentName    string `json:"name" gorm:"column:studentName"`
	StudentSex     string `json:"sex" gorm:"column:studentSex"`
	StudentAddress string `json:"address"  gorm:"column:address"`
	City           string `json:"city" gorm:"city"`
}

func (Student) TableName() string {
	return "student"
}
