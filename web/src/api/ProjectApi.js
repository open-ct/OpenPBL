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

const ProjectApi = {
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
      url: `/project/${pid}/publish`,
      method: 'post',
    })
  },
  closeProject(pid) {
    return request({
      url: `/project/${pid}/close`,
      method: 'post',
    })
  },
  deleteProject(pid) {
    return request({
      url: `/project/${pid}/delete`,
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

  getProjectStudents(pid) {
    return request({
      url: `/project/${pid}/students`,
      method: 'get',
    })
  },
  removeStudent(pid, sid) {
    return request({
      url: `/project/${pid}/remove/${sid}`,
      method: 'post'
    })
  },
  updateWeight(pid, data) {
    return request({
      url: `/project/${pid}/weight`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  addFavourite(pid) {
    return request({
      url: `/project/${pid}/favourite/add`,
      method: 'post',
    })
  },
  removeFavourite(pid) {
    return request({
      url: `/project/${pid}/favourite/remove`,
      method: 'post',
    })
  },
  getSubjectsAndSkills(pid) {
    return request({
      url: `/project/${pid}/subjects-skills`,
      method: 'get'
    })
  },
  viewProject(pid) {
    return request({
      url: `/project/${pid}/view`,
      method: 'post'
    })
  },
  cloneProject(pid) {
    return request({
      url: `/project/${pid}/clone`,
      method: 'post'
    })
  }
}

export default ProjectApi