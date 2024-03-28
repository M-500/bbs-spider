import api from '@/utils/ajax'


export const imageCaptchaAPI = (params = {}) => {
	return api.$get("/code", params)
}
