import request from './request'
import qs from 'qs'

export default {
  createStudent(data) {
    return request({
      url:'/student',
      method: 'post',
      data: qs.stringify(data)
    })
  },
  checkStudentEmail(email) {
    return request({
      url:'/student/checkout/' + email,
      method: 'get'
    })
  },
  getStudent(id) {
    return request({
      url: '/' + id,
      method: 'get'
    })
  },
  learnProject(pid) {
    return request({
      url: '/student/learn/' + pid,
      method: 'post',
    })
  },
  FinishedProject(sid, pid) {
    return request({
      url: '/student/finished',
      method: 'post',
      data: qs.stringify({
        studentId: sid,
        projectId: pid,
        learning: false
      })
    })
  }


}
