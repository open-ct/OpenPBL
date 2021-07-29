import request from "./request";
import qs from 'qs'

const TaskApi = {
  getSectionTasks(sid, pid) {
    return request({
      url: '/project/tasks/' + sid + '/' + pid,
      method: 'get',
    })
  },
  createTask(q) {
    return request({
      url: '/project/task',
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateTask(q) {
    return request({
      url: '/project/task/' + q.id,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  deleteTask(id) {
    return request({
      url: '/project/task/delete/' + id,
      method: 'post'
    })
  },
  exchangeTask(id1, id2) {
    return request({
      url: '/project/task/exchange/' + id1 + '/' + id2,
      method: 'post'
    })
  }
}

export default TaskApi