package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
	"math/rand"
)

type UserMysqlRepo struct {
	db *gorm.DB
}

func NewUserMysqlRepo(db *gorm.DB) *UserMysqlRepo {
	return &UserMysqlRepo{db}
}

func (u *UserMysqlRepo) Add(update *repositories.User) error {
	if err := u.db.Create(update).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserMysqlRepo) GetRandomID() uint64 {
	id := rand.Intn(10000000)
	for u.IsExistByID(uint64(id)) {
		id = rand.Intn(10000000)
	}
	return uint64(id)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *UserMysqlRepo) IsExistByID(userID uint64) bool {
	user := repositories.User{}
	notFound := u.db.Unscoped().Where("id = ?", userID).First(&user).RecordNotFound()
	return !notFound
}
func (u *UserMysqlRepo) IsExistByEmail(email string) bool {
	user := repositories.User{}
	notFound := u.db.Unscoped().Where("email = ?", email).First(&user).RecordNotFound()
	return !notFound
}

func (u *UserMysqlRepo) CheckAuth(user, pass string) (repositories.User, bool) {
	costumer := repositories.User{}
	notFound := u.db.Where("email = ? and password = ? ", user, pass).First(&costumer).RecordNotFound()
	return costumer, !notFound
}

func (u *UserMysqlRepo) Get(userID int) (repositories.User, bool) {
	user := repositories.User{}
	notFound := u.db.Where("id = ?", userID).First(&user).RecordNotFound()
	return user, !notFound
}

func (u *UserMysqlRepo) UpdatePicture(userID uint64, picture *string) error {
	return u.db.Model(&repositories.User{}).Select("picture").Where("id = ?", userID).Update(map[string]interface{}{"picture": picture}).Error
}
