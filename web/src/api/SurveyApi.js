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