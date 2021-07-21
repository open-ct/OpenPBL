import request from './request'
import qs from 'qs'

export default {
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
