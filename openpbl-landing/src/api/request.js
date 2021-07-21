import axios from 'axios'

const request = axios.create({
  baseURL: 'http://localhost:5000/api', // url = base url + request url
  timeout: 10000,
  withCredentials: true
})

export default request
