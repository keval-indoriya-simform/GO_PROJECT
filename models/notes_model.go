package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type Note struct {
	NoteID             int32            `gorm:"primaryKey" json:"note_id,omitempty"`
	Subject            string           `gorm:"type:varchar;not null" json:"subject,omitempty"`
	CustomerLocationID int32            `json:"customer_location_id,omitempty"`
	Attachment         string           `gorm:"type:varchar" json:"attachment,omitempty"`
	Note               string           `gorm:"type:varchar" json:"note,omitempty"`
	AssignedToID       int32            `gorm:"default:1" json:"assigned_to_id,omitempty"`
	CreatedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID    int32            `json:"created_by_user_id,omitempty"`
	UpdatedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID    int32            `json:"updated_by_user_id,omitempty"`
	DeletedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID    int32            `json:"deleted_by_user_id,omitempty"`
	CustomerLocation   CustomerLocation `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User               User             `gorm:"references:UserID;foreignKey:AssignedToID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (noteModel *Note) TableName() string {
	return "notes"
}

func (noteModel *Note) CreateNote() error {
	createQuery := initializers.DB().Create(&noteModel)
	if noteModel.NoteID > 0 {
		CreateLog(noteModel.CreatedByUserID, "Created", noteModel.TableName(), noteModel.NoteID)
		return nil
	}
	return createQuery.Error
}

func (noteModel *Note) UpdateNote() error {
	updateQuery := initializers.DB().Where("deleted_at IS NULL").
		Updates(&noteModel)
	if updateQuery.Error != nil {
		return updateQuery.Error
	} else if updateQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		CreateLog(noteModel.UpdatedByUserID, "Updated", noteModel.TableName(), noteModel.NoteID)
		return nil
	}
}

func (noteModel *Note) DeleteNote(noteIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND note_id IN ?", noteIDs).
		Updates(&noteModel)
	if deleteQuery.RowsAffected > 0 {
		var logs []Log
		for index := range noteIDs {
			logs = append(logs, Log{
				UserID:    noteModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: noteModel.TableName(),
				DataID:    noteIDs[index],
			})
		}
		CreateBatchLogs(logs)
		return nil
	} else if deleteQuery.Error == nil && deleteQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		return deleteQuery.Error
	}
}

func RetrieveNoteRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                     int64
		noteIdQuery, subjectQuery string
	)

	if queries["note_id"] != "" {
		noteIdQuery = "AND notes.note_id =" + queries["note_id"]
	}

	if queries["subject"] != "" {
		subjectQuery = "AND notes.subject ILIKE '%" + queries["subject"] + "%'"
	}
	retrievedRows := initializers.DB().Scopes(pagination).Model(&Note{}).
		Joins("JOIN customer_locations as c ON c.customer_location_id=notes.customer_location_id " +
			"JOIN users as u ON u.user_id=notes.assigned_to_id JOIN customers as ct ON ct.customer_id=notes.created_by_user_id").
		Where("notes.deleted_at IS NULL " + noteIdQuery + subjectQuery).
		Select("notes.note_id,notes.subject, notes.attachment, notes.note, notes.assigned_to_id," +
			"c.name as customer_location, u.name as assigned_to, ct.name as created_by, notes.created_at").
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
