import api from '@/utils/ajax'


export const imageCaptchaAPI = (params = {}) => {
	return api.$get("/code", params)
}

export const registerAPI = (params = {}) => {
	return api.$post("/sign-up", params)
}


export const pwdLoginAPI = (params = {}) => {
	return api.$post("/pwd-login", params)
}
