package models

import "github.com/jinzhu/gorm"

type Machine struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(50);unique_index"`
	Host     string `json:"host" gorm:"type:varchar(50)"`
	Ip       string `json:"host" gorm:"type:varchar(80)"`
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
	query := db
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
	return db.Create(ins).Error
}
