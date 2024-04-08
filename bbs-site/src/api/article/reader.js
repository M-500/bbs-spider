import api from '@/utils/ajax'


export const PubArticleListAPI = (params = {}) => {
  return api.$post("/articles/list", params)
}

export const PubArticleDetailAPI = (id,params = {}) => {
  return api.$get("/pub/"+id, params)
}

export const ArticleLikeDetailAPI = (params = {}) => {
  return api.$post("/pub/like", params)
}
