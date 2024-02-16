package user

import "gorm.io/gorm"

type Repository interface {
	// 获取用户信息
	Save(user User) (User, error)
}
type repository struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}


func (r *repository) Save(user User) (User, error) {
	err:= r.db.Create(&user).Error
	if err != nil{
		return User{}, err
	}
	return user, nil
}