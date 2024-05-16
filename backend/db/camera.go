package db

import (
	"backend/proto"
	"gorm.io/gorm"
)

type Camera struct {
	gorm.Model
	DeviceName string `gorm:"column:device_name"`
	Username   string `gorm:"column:username"`
	Passwd     string `gorm:"column:passwd"`
	Url        string `gorm:"column:url"`
}

func (c Camera) DBtoProto() *proto.Camera {
	return &proto.Camera{
		Uuid:       uint32(c.ID),
		DeviceName: c.DeviceName,
		Username:   c.Username,
		Passwd:     c.Passwd,
		Url:        c.Url,
	}
}

func CameraProtoToDB(c *proto.Camera) *Camera {
	return &Camera{
		Model:      gorm.Model{ID: uint(c.Uuid)},
		DeviceName: c.DeviceName,
		Username:   c.Username,
		Passwd:     c.Passwd,
		Url:        c.Url,
	}
}

func (db DB) GetCamera() (res []Camera, err error) {
	return
}

func (db DB) UpsertCamera() error {
	return nil
}

func (db DB) DeleteCamera() error {
	return nil
}
