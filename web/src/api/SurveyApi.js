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

const SurveyApi = {
  getSurveyDetailByTaskId(pid, tid) {
    return request({
      url: `/project/${pid}/task/${tid}/survey`,
      method: 'get'
    })
  },
  createQuestion(pid, tid, q) {
    return request({
      url: `/project/${pid}/task/${tid}/survey/${q.surveyId}/question`,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateQuestion(pid, tid, q) {
    return request({
      url: `/project/${pid}/task/${tid}/survey/${q.surveyId}/question/${q.id}`,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  deleteQuestion(pid, tid, sid, qid) {
    return request({
      url: `/project/${pid}/task/${tid}/survey/${sid}/question/${qid}/delete`,
      method: 'post',
    })
  },
  exchangeQuestion(pid, tid, suid, id1, id2) {
    return request({
      url: `/project/${pid}/task/${tid}/survey/${suid}/questions/exchange`,
      method: 'post',
      data: qs.stringify({
        questionId1: id1,
        questionId2: id2
      })
    })
  }
}

export default SurveyApi