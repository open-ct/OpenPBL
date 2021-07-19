import axios from 'axios'

// 创建axios实例
const service = axios.create({
  baseURL: 'http://localhost:5000/v1', // url = base url + request url
  timeout: 10000 // request timeout
})

// request拦截器
/*
service.interceptors.request.use(
  config => {
    // resolve 401
    if (store.state.token!='') {// 让每个请求携带自定义token
      config.headers['Authorization'] = store.state.token
    }
    return config
  },  error => {
    console.log(error) // for debug
    return Promise.reject(error)
  }
)
*/
// respone拦截器
// service.interceptors.response.use(
//   response => {
//     const res = response.data
//     /**
//      * 如果非200，响应在此被拦截
//      */
//     if (res.code !== 200) {
//       Message({
//         message: res.message || 'Error',
//         type: 'error',
//         duration: 5 * 1000
//       })
//       if (res.code === 403) {
//         MessageBox.confirm('你已被登出，点击确定返回首页', '状态异常', {
//           confirmButtonText: '确定',
//           cancelButtonText: '取消',
//           type: 'warning'
//         }).then(() => {
//           router.push({ path: '/login'})
//           location.reload() //跳转+强制刷新
//         })
//       }
//       if (res.code === 401) {
//         MessageBox.confirm('请先登录', '状态异常', {
//           confirmButtonText: '确定',
//           cancelButtonText: '取消',
//           type: 'warning'
//         }).then(() => {
//           router.push({ path: '/login'})
//           // location.reload() //跳转+强制刷新
//         })
//       }
//       return Promise.reject('Error')
//     } else {
//       return res
//     }
//   },
//   error => {
//     console.log('错误' + error) // for debug
//     Message({
//       message: error.message,
//       type: 'error',
//       duration: 5 * 1000
//     })
//     return Promise.reject(error)
//   }
// )

export default service
