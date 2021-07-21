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
  updateProjectChapter(chapter) {
    return request({
      url:'/project/chapter/' + chapter.id,
      method: 'post',
      data: qs.stringify(chapter)
    })
  },
  deleteProjectChapter(id) {
    return request({
      url:'/project/chapter/delete/' + id,
      method: 'post',
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
  updateChapterSection(section) {
    return request({
      url:'/project/chapter/section/' + section.id,
      method: 'post',
      data: qs.stringify(section)
    })
  },
  deleteChapterSection(id) {
    return request({
      url:'/project/chapter/section/delete/' + id,
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
