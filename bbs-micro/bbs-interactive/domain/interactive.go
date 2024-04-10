package domain

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:35

type Interactive struct {
	Biz        string
	BizId      int64
	ReadCnt    int64 `json:"read_cnt"`
	LikeCnt    int64 `json:"like_cnt"`
	CollectCnt int64 `json:"collect_cnt"`
	CommentCnt int64 `json:"comment_cnt"`
	Liked      bool  `json:"liked"`
	Collected  bool  `json:"collected"`
}
