import request from './request'
import qs from 'qs'

export default {
  getProjectListByTid(id, mode, q) {
    return request({
      url:'/project-list/teacher/' + mode + '/' + id,
      params: q,
      method: 'get'
    })
  },

  getProjectListBySid(id, mode, q) {
    return request({
      url:'/project-list/student/' + mode + '/' + id,
      params: q,
      method: 'get'
    })
  },

  getStudentPublicProjects(id, q) {
    return request({
      url:'/project-list/public/student/' + id,
      params: q,
      method: 'get'
    })
  },

  getTeacherPublicProjects(id, q) {
    return request({
      url:'/project-list/public/teacher/' + id,
      params: q,
      method: 'get'
    })
  },
}
