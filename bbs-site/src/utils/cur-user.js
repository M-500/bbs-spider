const TOKEN_KEY = "cust-token"
const USER_KEY = "username"
const USER_ID_KEY = "uid"
const USER_Avatar = "avatar"

export default {
    getToken () {
        return localStorage.getItem(TOKEN_KEY)
    },
    getUserName () {
        return localStorage.getItem(USER_KEY)
    },
  getUserId(){
    return localStorage.getItem(USER_ID_KEY)
  },
  getUserAvatar(){
    return localStorage.getItem(USER_Avatar)
  },
    /**
     * 登录成功后，保存token到localStorage
     */
    setToken (token) {
        localStorage.setItem(TOKEN_KEY, token)
    },
    setUserName (username) {
        localStorage.setItem(USER_KEY, username)
    },
  setUserId (username) {
    localStorage.setItem(USER_ID_KEY, username)
  },
  setUserAvatar (username) {
    localStorage.setItem(USER_Avatar, username)
  }
}
