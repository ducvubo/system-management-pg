package model

type CategoryBlogOutput struct {
	CatBlId     string `json:"cat_bl_id"`
	CatBlName   string `json:"cat_bl_name"`
	CatBlSlug   string `json:"cat_bl_slug"`
	CatBlDesc   string `json:"cat_bl_description"`
	CatBlOrder  int32  `json:"cat_bl_order"`
	CatBlStatus int32  `json:"cat_bl_status"`
	IsDeleted   int32  `json:"isDeleted"`
}

type CreateCategoryBlogDto struct {
	CatBlName  string `json:"cat_bl_name" binding:"required"`
	CatBlDesc  string `json:"cat_bl_description" binding:"required"`
	CatBlOrder int32  `json:"cat_bl_order" binding:"required"`
}

type UpdateCategoryBlogDto struct {
	CatBlId    string `json:"cat_bl_id" binding:"required"`
	CatBlName  string `json:"cat_bl_name" binding:"required"`
	CatBlDesc  string `json:"cat_bl_description" binding:"required"`
	CatBlOrder int32  `json:"cat_bl_order" binding:"required"`
}

type UpdateStatusCategoryBlogDto struct {
	CatBlId     string `json:"cat_bl_id" binding:"required"`
	CatBlStatus int32  `json:"cat_bl_status" binding:"required"`
}