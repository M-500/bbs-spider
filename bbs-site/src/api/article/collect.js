import api from '@/utils/ajax'

export const createCollectAPI = (params = {}) => {
  return api.$post("/collect/create", params)
}
