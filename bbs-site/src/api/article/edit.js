import api from '@/utils/ajax'


export const editArticleAPI = (params = {}) => {
	return api.$post("/articles/edit", params)
}
