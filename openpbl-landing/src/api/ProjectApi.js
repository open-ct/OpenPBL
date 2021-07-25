import request from './request'
import qs from 'qs'

export default {
  getProjectDetail(id) {
    return request({
      url:'/project/' + id,
      method: 'get'
    })
  },
  createProject(data) {
    return request({
      url:'/project',
      method: 'post',
      data: qs.stringify(data)
    })
  },
  updateProject(data, id) {
    return request({
      url:'/project/' + id,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  publishProject(pid) {
    return request({
      url:'/project/publish/' + pid,
      method: 'post',
    })
  },
  closeProject(pid) {
    return request({
      url:'/project/close/' + pid,
      method: 'post',
    })
  },
  deleteProject(pid) {
    return request({
      url:'/project/delete/' + pid,
      method: 'post',
    })
  },

  getSectionFiles(id) {
    return request({
      url:'/project/chapter/section/files/' + id,
      method: 'get',
    })
  },
  createSectionFile(f) {
    return request({
      url:'/project/chapter/section/file',
      method: 'post',
      data: qs.stringify(f)
    })
  },

  getProjectStudents(pid) {
    return request({
      url: '/project/students/' + pid,
      method: 'get',
    })
  }

}
