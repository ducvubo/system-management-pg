package model

type InternalNoteOutput struct {
	ItnNoteId      string `json:"itn_note_id" binding:"required"`
	ItnNoteTitle   string `json:"itn_note_title" binding:"required"`
	ItnNoteContent string `json:"itn_note_content" binding:"required"`
	ItnNoteType    string `json:"itn_note_type" binding:"required"`
	Isdeleted      int32  `json:"isDeleted"`
}

type CreateInternalNoteDto struct {
	ItnNoteTitle   string `json:"itn_note_title" binding:"required"`
	ItnNoteContent string `json:"itn_note_content" binding:"required"`
	ItnNoteType    string `json:"itn_note_type" binding:"required"`
}

type UpdateInternalNoteDto struct {
	ItnNoteId      string `json:"itn_note_id" binding:"required"`
	ItnNoteTitle   string `json:"itn_note_title" binding:"required"`
	ItnNoteContent string `json:"itn_note_content" binding:"required"`
	ItnNoteType    string `json:"itn_note_type" binding:"required"`
}
