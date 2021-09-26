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
import qs from "qs";

const ChapterApi = {
  getProjectChapters(id, studentId) {
    return request({
      url: `/project/${id}/chapters`,
      method: 'get',
      params: {studentId: studentId}
    })
  },
  createProjectChapter(chapter) {
    return request({
      url: `/project/${chapter.projectId}/chapter`,
      method: 'post',
      data: qs.stringify(chapter)
    })
  },
  updateProjectChapter(chapter) {
    return request({
      url: `/project/${chapter.projectId}/chapter/${chapter.id}`,
      method: 'post',
      data: qs.stringify(chapter)
    })
  },

  deleteProjectChapter(c) {
    return request({
      url: `/project/${c.projectId}/chapter/${c.id}/delete`,
      method: 'post',
      data: qs.stringify(c)
    })
  },
  exchangeProjectChapter(pid, id1, id2) {
    return request({
      url: `/project/${pid}/chapters/exchange`,
      method: 'post',
      data: qs.stringify({chapterId1: id1, chapterId2: id2})
    })
  },
}

export default ChapterApi