package model

type BlogOutput struct {
	BlID                string `json:"bl_id"`
	CatBlID             string `json:"cat_bl_id"`
	BlTitle             string `json:"bl_title"`
	BlDescription       string `json:"bl_description"`
	BlSlug              string `json:"bl_slug"`
	BlStatus            int32  `json:"bl_status"`
	BlImage             string `json:"bl_image"`
	BlContent           string `json:"bl_content"`
	IsDeleted           int32  `json:"isDeleted"`
	BlType              int32  `json:"bl_type"`
	BlView              int32  `json:"bl_view"`
	BlPublishedTime     string `json:"bl_published_time"`
	BlPublishedSchedule string `json:"bl_published_schedule"`
}

type CreateBlogDto struct {
	CatBlID       string        `json:"cat_bl_id" binding:"required"`
	BlTitle       string        `json:"bl_title" binding:"required"`
	BlDescription string        `json:"bl_description" binding:"required"`
	BlImage       string        `json:"bl_image" binding:"required"`
	BlContent     string        `json:"bl_content" binding:"required"`
	BlType        int32         `json:"bl_type" binding:"required"`
	BlogNote      []BlogNote    `json:"blog_note"`
	BlogRelated   []BlogRelated `json:"blog_related"`
}

type UpdateBlogDto struct {
	BlID          string        `json:"bl_id" binding:"required"`
	CatBlID       string        `json:"cat_bl_id" binding:"required"`
	BlTitle       string        `json:"bl_title" binding:"required"`
	BlDescription string        `json:"bl_description" binding:"required"`
	BlImage       string        `json:"bl_image" binding:"required"`
	BlContent     string        `json:"bl_content" binding:"required"`
	BlogNote      []BlogNote    `json:"blog_note"`
	BlogRelated   []BlogRelated `json:"blog_related"`
}

type UpdateStatusBlogDto struct {
	BlID     string `json:"bl_id" binding:"required"`
	BlStatus int32  `json:"bl_status" binding:"required"`
}

type BlogNote struct {
	BlContent string `json:"bl_content" binding:"required"`
}

type BlogRelated struct {
	BlRltID string `json:"bl_rlt_id" binding:"required"`
}
