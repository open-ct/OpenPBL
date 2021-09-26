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

const SubmitApi = {
  createSubmit(pid, tid, data) {
    return request({
      url: `/project/${pid}/task/${tid}/submit`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  updateSubmit(pid, tid, sid, data) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  getSubmitFiles(pid, tid, sid) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/files`,
      method: 'get',
    })
  },
  uploadSubmitFile(pid, tid, sid, data) {
    if (sid === "") {
      sid = "0"
    }
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/file`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  updateSubmitFile(pid, tid, sid, file) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/file/${file.id}/update`,
      method: 'post',
      data: qs.stringify(file)
    })
  },
  deleteSubmitFile(pid, tid, sid, fid) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/file/${fid}/delete`,
      method: 'post'
    })
  }
}

export default SubmitApi
