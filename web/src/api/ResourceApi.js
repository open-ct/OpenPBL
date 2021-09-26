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

const ResourceApi = {
  getResource(pid, id) {
    return request({
      url: `project/${pid}/resource/${id}`,
      method: 'get',
    })
  },
  createResource(pid, q) {
    return request({
      url: `/project/${pid}/resource`,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateResource(pid, q) {
    return request({
      url: `/project/${pid}/resource/${q.id}`,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateResourceContent(pid, id, content) {
    return request({
      url: `/project/${pid}/resource/${id}/content`,
      method: 'post',
      data: qs.stringify({content: content})
    })
  },

}

export default ResourceApi