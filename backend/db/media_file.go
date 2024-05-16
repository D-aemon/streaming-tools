package db

import (
	"backend/proto"
	"gorm.io/gorm"
)

type MediaFile struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Type    string `gorm:"column:type"`
	Size    string `gorm:"column:size"`
	Path    string `gorm:"column:path"`
	Comment string `gorm:"column:comment"`
}

func (f MediaFile) DBtoProto() *proto.MediaFile {
	return &proto.MediaFile{
		Uuid:    uint32(f.ID),
		Name:    f.Name,
		Type:    f.Type,
		Size:    f.Type,
		Path:    f.Path,
		Comment: f.Comment,
	}
}

func MediaFileProtoToDB(f *proto.MediaFile) MediaFile {
	return MediaFile{
		Model:   gorm.Model{ID: uint(f.Uuid)},
		Name:    f.Name,
		Type:    f.Type,
		Size:    f.Size,
		Path:    f.Path,
		Comment: f.Comment,
	}
}

func (db DB) GetMediaFile() (res []MediaFile, err error) {
	return
}

func (db DB) UpsertMediaFile() error {
	return nil
}

func (db DB) DeleteMediaFile() error {
	return nil
}
