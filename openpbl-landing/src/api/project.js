import request from './request'
import qs from 'qs'

export default {
  getProjectDetailForStudent(id) {
    return request({
      url:'/project/student/' + id,
      method: 'get'
    })
  },
  getProjectDetailForTeacher(id) {
    return request({
      url:'/project/teacher/' + id,
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
  updateProject(data) {
    return request({
      url:'/project/info',
      method: 'post',
      data: qs.stringify(data)
    })
  },
  getProjectChapters(id) {
    return request({
      url:'/project/chapters/' + id,
      method: 'get',
    })
  },
  createProjectChapter(chapter) {
    return request({
      url:'/project/chapter',
      method: 'post',
      data: qs.stringify(chapter)
    })
  },
  getChapterSections(id) {
    return request({
      url:'/project/chapter/sections/' + id,
      method: 'get',
    })
  },
  createChapterSection(section) {
    return request({
      url:'/project/chapter/section',
      method: 'post',
      data: qs.stringify(section)
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

  getSubmitFiles(sid, pid) {
    return request({
      url:'/project/submit-files/' + pid + '/' + sid,
      method: 'get',
    })
  },
  createSubmitFiles(data) {
    return request({
      url:'/project/submit-files',
      method: 'post',
      data: qs.stringify(data)
    })
  },

  getProjectStudents(pid) {
    return request({
      url: '/project/students/' + pid,
      method: 'get',
    })
  }

}
