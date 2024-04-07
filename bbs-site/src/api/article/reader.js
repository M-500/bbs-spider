import api from '@/utils/ajax'


export const PubArticleListAPI = (params = {}) => {
  return api.$post("/articles/list", params)
}
