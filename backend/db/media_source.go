package db

import (
	"backend/proto"
	"gorm.io/gorm"
)

type MediaSource struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Type string `gorm:"column:type"`
	Url  string `gorm:"column:url"`
}

func (s MediaSource) DBtoProto() *proto.MediaSource {
	return &proto.MediaSource{
		Uuid: uint32(s.ID),
		Name: s.Name,
		Type: s.Type,
		Url:  s.Url,
	}
}

func MediaSourceProtoToDB(m *proto.MediaSource) *MediaSource {
	return &MediaSource{
		Model: gorm.Model{ID: uint(m.Uuid)},
		Name:  m.Name,
		Type:  m.Type,
		Url:   m.Url,
	}
}

func (db DB) GetMediaSource() (res []MediaSource, err error) {
	return
}

func (db DB) UpsertMediaSource() error {
	return nil
}

func (db DB) DeleteMediaSource() error {
	return nil
}
