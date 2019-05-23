package models

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var _ = time.Thursday

//User
type User struct {
	Id        uint       `gorm:"column:id" form:"id" json:"id" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:"PRI"`
	Username  string     `gorm:"column:username" form:"username" json:"username" comment:"昵称/登陆用户名" columnType:"varchar(50)" dataType:"varchar" columnKey:"UNI"`
	FullName  string     `gorm:"column:full_name" form:"full_name" json:"full_name" comment:"真实姓名" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	Email     string     `gorm:"column:email" form:"email" json:"email" comment:"邮箱" columnType:"varchar(255)" dataType:"varchar" columnKey:"UNI"`
	Mobile    string     `gorm:"column:mobile" form:"mobile" json:"mobile" comment:"手机号码" columnType:"varchar(11)" dataType:"varchar" columnKey:"UNI"`
	Password  string     `gorm:"column:password" form:"password" json:"password" comment:"密码" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	RoleId    uint       `gorm:"column:role_id" form:"role_id" json:"role_id" comment:"角色ID:2-超级用户,4-普通用户" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Status    uint       `gorm:"column:status" form:"status" json:"status" comment:"状态: 1-正常,2-禁用/删除" columnType:"int(10) unsigned" dataType:"int" columnKey:""`
	Avatar    string     `gorm:"column:avatar" form:"avatar" json:"avatar" comment:"用户头像" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	Remark    string     `gorm:"column:remark" form:"remark" json:"remark" comment:"备注" columnType:"varchar(255)" dataType:"varchar" columnKey:""`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at" comment:"" columnType:"timestamp" dataType:"timestamp" columnKey:""`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at" comment:"" columnType:"timestamp" dataType:"timestamp" columnKey:""`
}

//TableName
func (m *User) TableName() string {
	return "users"
}

//One
func (m *User) One() (one *User, err error) {
	one = &User{}
	err = crudOne(m, one)
	return
}

//All
func (m *User) All(q *PaginationQuery) (list *[]User, total uint, err error) {
	list = &[]User{}
	total, err = crudAll(m, q, list)
	return
}

//Update
func (m *User) Update() (err error) {
	where := User{Id: m.Id}
	m.Id = 0
	m.makePassword()

	return crudUpdate(m, where)
}

//Create
func (m *User) Create() (err error) {
	m.Id = 0
	m.makePassword()

	return db.Create(m).Error
}

//Delete
func (m *User) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

//Login
func (m *User) Login(ip string) (*jwtObj, error) {
	m.Id = 0
	if m.Password == "" {
		return nil, errors.New("password is required")
	}
	inputPassword := m.Password
	m.Password = ""
	loginTryKey := "login:" + ip
	loginRetries, _ := mem.GetUint(loginTryKey)
	if loginRetries > 10 {
		memExpire := 30
		return nil, fmt.Errorf("for too many wrong login retries the %s will ban for login in %d minitues", ip, memExpire)
	}
	//you can implement more detailed login retry rule
	//for i don't know what your login username i can't implement the ip+username rule in my boilerplate project
	// about username and ip retry rule

	err := db.Where(m).First(&m).Error
	if err != nil {
		//username fail ip retries add 5
		loginRetries = loginRetries + 5
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
	//password is set to bcrypt check
	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(inputPassword)); err != nil {
		// when password failed reties will add 1
		loginRetries = loginRetries + 1
		mem.Set(loginTryKey, loginRetries)
		return nil, err
	}
	m.Password = ""
	key := fmt.Sprintf("login:%d", m.Id)

	//save login user  into the memory store

	data, err := jwtGenerateToken(m)
	mem.Set(key, data)
	return data, err
}

func (m *User) makePassword() {
	if m.Password != "" {
		if bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost); err != nil {
			logrus.WithError(err).Error("bcrypt making password is failed")
		} else {
			m.Password = string(bytes)
		}
	}
}

func (m *User) CreateInitUser() error {
	m.Id = 0
	m.makePassword()
	return db.Where("username = ?", m.Username).FirstOrCreate(m).Error
}
