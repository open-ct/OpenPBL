import request from './request'

export default {
  getUserProjectList(mode, q) {
    return request({
      url: '/project-list/' + mode,
      params: q,
      method: 'get'
    })
  },
}
