import api from '@/utils/ajax'

export const createCollectAPI = (params = {}) => {
  return api.$post("/collect/create", params)
}

export const getCollectListAPI = (id,params = {}) => {
  return api.$get(`/collect/${id}/list`, params)
}
export const collectEntityAPI = (cid,bid,params = {}) => {
  return api.$post(`/collect/${cid}/entity/${bid}`, params)
}