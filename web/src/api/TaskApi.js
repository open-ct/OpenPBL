// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
  getProjectTasksDetail(pid, sid) {
    return request({
      url: `/project/${pid}/tasks-detail`,
      method: 'get',
      params: {
        studentId: sid
      }
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