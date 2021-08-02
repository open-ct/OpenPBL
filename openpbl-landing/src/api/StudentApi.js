import request from './request'
import qs from 'qs'

const StudentApi = {
  learnProject(pid) {
    return request({
      url: '/student/learn/' + pid,
      method: 'post',
    })
  },
  exitProject(pid) {
    return request({
      url: '/student/exit/' + pid,
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
  },
  getLearnSection(sid) {
    return request({
      url: `/student/section/${sid}`,
      method: 'get'
    })
  },
  updateLearnSection(sid, data) {
    return request({
      url: `/student/section/${sid}`,
      method: 'post',
      data: qs.stringify(data)
    })
  }
}

export default StudentApi