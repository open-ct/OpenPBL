import request from "./request";
import qs from 'qs'

const TaskApi = {
  getSectionTasks(sid, pid) {
    return request({
      url: `/project/${pid}/section/${sid}/tasks`,
      method: 'get',
    })
  },
  getProjectTasks(pid) {
    return request({
      url: `/project/${pid}/tasks`,
      method: 'get',
    })
  },
  getProjectTasksAndSubmits(pid) {
    return request({
      url: `/project/${pid}/tasks-submits`,
      method: 'get',
    })
  },
  createTask(pid, q) {
    return request({
      url: `/project/${pid}/task`,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateTask(pid, q) {
    return request({
      url: `/project/${pid}/task/${q.id}`,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  deleteTask(pid, id) {
    return request({
      url: `/project/${pid}/task/${id}/delete`,
      method: 'post'
    })
  },
  exchangeTask(pid, id1, id2) {
    return request({
      url: `/project/${pid}/tasks/exchange`,
      method: 'post',
      data: qs.stringify({
        taskId1: id1,
        taskId2: id2
      })
    })
  }
}

export default TaskApi