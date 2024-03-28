/**
 * 扩展原来Date的方法，提供格式化、解析等方法
 */

/**
 * 解析后台传过来的"yyyy-MM-ddTHH:mm:ss"格式的时间，返回Date对象
 */
Date.from = function (v) {
    if (!v) return null
    if (v instanceof Date) return v
    if (Number(v)) return new Date(v)
    //不可直接使用new Date(str)，某些浏览器下时区不正确
    return v && new Date(v.replace(/-/g, '/').replace('T', ' ').split('.')[0])
}

/**
 * 日期格式化方法：new Date().format('yyyy-M-d hh:mm')
 * @author Linzongxue 2019-10-17
 */
if (!Date.prototype.format) {
    const weekDays = ['日', '一', '二', '三', '四', '五', '六']
    const zeroize = num => (num > 9 ? '' : '0') + num
    Date.prototype.format = function (fmt) {
        return fmt ? fmt.replace(/(y+|M+|d+|w+|H+|h+|m+|s+)/g, v => {
            switch (v) {
                case 'yy': return String(this.getFullYear()).substring(2)
                case 'yyyy': return this.getFullYear()
                case 'MM': return zeroize(this.getMonth() + 1)
                case 'M': return this.getMonth() + 1
                case 'dd': return zeroize(this.getDate())
                case 'd': return this.getDate()
                case 'w': return weekDays[this.getDay()]
                case 'HH': //同hh
                case 'hh': return zeroize(this.getHours())
                case 'H': //同h
                case 'h': return this.getHours()
                case 'mm': return zeroize(this.getMinutes())
                case 'm': return this.getMinutes()
                case 'ss': return zeroize(this.getSeconds())
                case 's': return this.getSeconds()
                default: return ''
            }
        }) : ''
    }
}

Date.prototype.toString = function () {
    return this.format('yyyy-MM-dd hh:mm:ss');
}

/**
 * 处理时区问题，方便发请求时直接使用Date对象
 */
Date.prototype.toISOString = function () {
    return this.format('yyyy-MM-ddThh:mm:ss')
}