package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Machine struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(50);unique_index"`
	Host     string `json:"host" gorm:"type:varchar(50)"`
	Ip       string `json:"ip" gorm:"type:varchar(80)"`
	Port     uint   `json:"port" gorm:"type:int(6)"`
	User     string `json:"user" gorm:"type:varchar(20)"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Type     string `json:"type" gorm:"type:varchar(20)"`
}

func MachineAdd(name, addr, ip, user, password, key, auth string, port uint) error {
	ins := &Machine{Name: name, Ip: ip, Host: addr, User: user, Password: password, Key: key, Type: auth, Port: port}
	return db.Create(ins).Error
}

func MachineAll(search string) ([]Machine, error) {
	var hs []Machine
	query := db.Order("updated_at")
	if search != "" {
		query = db.Where("name like ?", "%"+search+"%")
	}
	err := query.Find(&hs).Error
	return hs, err
}

func MachineFind(idx uint) (*Machine, error) {
	ins := &Machine{}
	ins.ID = idx
	return ins, db.First(ins).Error
}

func MachineDelete(idx uint) error {
	ins := Machine{}
	ins.ID = idx
	return db.Where("id = ?", idx).Delete(&ins).Error
}
func MachineDeleteAll() error {
	ins := Machine{}
	return db.Delete(&ins).Error
}

func MachineUpdate(name, addr, user, password, pkey, t string, id, port uint) error {
	ins := Machine{Name: name, Host: addr, User: user, Password: password, Key: pkey, Type: t, Port: port}
	wh := &Machine{}
	wh.ID = id
	return db.Model(wh).Updates(ins).Error
}

func MachineDuplicate(idx uint) error {
	ins := &Machine{}
	ins.ID = idx
	err := db.First(ins).Error
	if err != nil {
		return err
	}
	ins.ID = 0
	ins.Name = fmt.Sprintf("%s_du", ins.Name)
	return db.Create(ins).Error
}

func (m *Machine) One() (one *Machine, err error) {
	one = &Machine{}
	err = crudOne(m, one)
	return
}

//All get all for pagination
func (m *Machine) All(q *PaginationQuery) (list *[]Machine, total uint, err error) {
	list = &[]Machine{}
	total, err = crudAll(m, q, list)
	return
}

//Update a row
func (m *Machine) Update() (err error) {
	where := Machine{Model: gorm.Model{ID: m.ID}}
	return crudUpdate(m, where)
}

//Create insert a row
func (m *Machine) Create() (err error) {
	m.ID = 0
	return db.Create(m).Error
}

//Delete destroy a row
func (m *Machine) Delete() (err error) {
	if m.ID == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}

func (m *Machine) ChangeUpdateTime() (err error) {
	m.UpdatedAt = time.Now()
	return db.Save(m).Error
}
