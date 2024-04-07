/**
 * Ajax方法封装
 * @Author linlin 2022-7-27
 *
 * */
import qs from 'qs'
import axios from 'axios'
import Vue from 'vue'
import { host } from '@/config'
import curUser from './cur-user'
import { Message } from "element-ui";

Vue.prototype.$axios = axios

const service = axios.create({
    baseURL: host,
    timeout: 180 * 1000,
    validateStatus: x => x === 200
})

const done = function (h) {
    return this.then(h).catch()
}
/**
 * 请求发送前处理
 * @param {any} o 配置前
 * @returns {object} 配置後
 */
const onsend = o => {
    const token = curUser.getToken()
    if (token) o.headers.common.Authorization = token
    if (typeof o.data === 'string' || typeof o.data === 'object') {
        o.headers[o.method]['Content-Type'] = 'application/json'

    } else if (Array.isArray(o.data)) {
        const form = new FormData()
        Object.entries(o.data[0]).forEach(([k, v]) => {
            if (Array.isArray(v)) {
                v.forEach(x => form.append(k, x))
            } else {
                form.append(k, v)
            }
        })
        o.data = form
    } else {
        o.headers[o.method]['Content-Type'] = 'application/json'
        o.data = qs.stringify(o.data, { indices: false, allowDots: true })
    }
    o.url += o.url.indexOf('?') > -1 ? '&_t=' + new Date().getTime() : '?_t=' + new Date().getTime();
    // o.headers.common.version = version
    return o
}

/**
 * 请求成功回调
 * @param {object} o 请求对象
 * @returns {object} 响应
 */
const onsuccess = o => {
    if (o.status !== 200) throw o
    else if (o.config.responseType === 'blob') {
        return {
            key: o.headers.key,
            blob: o.data
        }
    }

    const { code, msg, data } = o.data
    if (code === 0 || code === 200) {
        return data
    } else {
        throw { code, msg, response: o }
    }
}

/**
 * 请求失败回调方法
 */
const onerror = e => {
    const resp = e.response
    console.log("diudiudiu", resp)
    if (resp) {
        if(resp.status == 403){
          window.localStorage.clear()
          window.sessionStorage.clear()
          this.$router.push("/login")
          throw resp.statusText
        }
        if (resp.data) {
            throw resp.data
        }
    }
    throw {
        code: resp ? resp.status : 0,
        msg: e.toString(),
        e
    }
}

Promise.prototype.done = done
window.Promise.prototype.done = done
service.interceptors.request.use(onsend)
service.interceptors.response.use(onsuccess, onerror)
const paramsSerializer = p => qs.stringify(p, { indices: false, allowDots: true })
export default {
    $get: (url, params, cfg = {}) => service.get(url, Object.assign({ params, paramsSerializer }, cfg)),
    $pop: (url, params, cfg) => service.delete(url, Object.assign({ params, paramsSerializer }, cfg)),
    $put: (url, o, cfg) => service.put(url, o, cfg),
    $post: (url, o, cfg) => service.post(url, o, cfg),
    $patch: (url, o, cfg) => service.patch(url, o, cfg),
    $form: (url, o, config = {}) => service.post(url, [o], config),
    $auth: url => service.get(url, { responseType: 'blob' }),
    $getBlob: (url, params, cfg) => service.get(url, Object.assign({ params, paramsSerializer, responseType: 'blob' }, cfg))
}
