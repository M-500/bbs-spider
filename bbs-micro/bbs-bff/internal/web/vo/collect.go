package vo

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 11:25

type CreateCollectReq struct {
	CollectName string `json:"collect_name"`
	Desc        string `json:"desc"`
	IsPublic    bool   `json:"is_public"`
}
