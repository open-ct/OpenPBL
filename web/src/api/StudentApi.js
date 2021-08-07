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
  getLearnSection(pid, sid) {
    return request({
      url: `/student/project/${pid}/section/${sid}`,
      method: 'get'
    })
  },
  updateLearnSection(pid, sid, data) {
    return request({
      url: `/student/project/${pid}/section/${sid}`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  getLastLearnSection(pid) {
    return request({
      url: `/student/last-learn/project/${pid}`,
      method: 'get'
    })
  }
}

export default StudentApi