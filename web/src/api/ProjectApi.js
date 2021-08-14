import request from './request'
import qs from 'qs'

const ProjectApi = {
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
      url: `/project/${pid}/publish`,
      method: 'post',
    })
  },
  closeProject(pid) {
    return request({
      url: `/project/${pid}/close`,
      method: 'post',
    })
  },
  deleteProject(pid) {
    return request({
      url: `/project/${pid}/delete`,
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
      url: `/project/${pid}/students`,
      method: 'get',
    })
  },
  removeStudent(pid, sid) {
    return request({
      url: `/project/${pid}/remove/${sid}`,
      method: 'post'
    })
  },
  updateWeight(pid, data) {
    return request({
      url: `/project/${pid}/weight`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  addFavourite(pid) {
    return request({
      url: `/project/${pid}/favourite/add`,
      method: 'post',
    })
  },
  removeFavourite(pid) {
    return request({
      url: `/project/${pid}/favourite/remove`,
      method: 'post',
    })
  },
  getSubjectsAndSkills(pid) {
    return request({
      url: `/project/${pid}/subjects-skills`,
      method: 'get'
    })
  }
}

export default ProjectApi