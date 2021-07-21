import request from './request'
import qs from 'qs';

export default {
  login(code, state) {
    return request({
      url: '/auth/login',
      method: 'post',
      data: qs.stringify({code: code, state: state}),
    })
  },
  logout() {
    return request({
      url: '/auth/logout',
      method: 'post'
    })
  },
  getAccount() {
    return request({
      url: '/auth/account',
      method: 'get',
    })
  }
}
