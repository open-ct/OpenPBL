import axios from 'axios'
import {message} from 'antd'

function getServerUrl() {
  const hostname = window.location.hostname
  if (hostname === 'localhost') {
    return `http://${hostname}:5000/api`
  }
  return '/api'
}

const request = axios.create({
  baseURL: getServerUrl(),
  timeout: 10000,
  withCredentials: true
})

request.interceptors.response.use(res=>{
  if (res.data.code === 401) {
    message.error(res.data.msg)
    setTimeout(()=>{
      window.location.href = '/'
    }, 1000)
  } else if (res.data.code === 403) {
    message.error(res.data.msg)
    setTimeout(()=>{
      window.location.href = '/'
    }, 1000)
  } else {
    return res
  }
})

export default request
