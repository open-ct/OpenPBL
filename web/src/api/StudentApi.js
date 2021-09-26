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