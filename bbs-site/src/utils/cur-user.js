const TOKEN_KEY = "cust-token"
const USER_KEY = "username"

export default {
    getToken () {
        return localStorage.getItem(TOKEN_KEY)
    },
    getUserName () {
        return localStorage.getItem(USER_KEY)
    },
    /**
     * 登录成功后，保存token到localStorage
     */
    setToken (token) {
        localStorage.setItem(TOKEN_KEY, token)
    },
    setUserName (username) {
        localStorage.setItem(USER_KEY, username)
    }
}
