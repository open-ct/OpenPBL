import request from "./request";
import qs from 'qs'

const SurveyApi = {
  getSurveyDetailByTaskId(tid) {
    return request({
      url: '/project/task/survey/' + tid,
      method: 'get'
    })
  },
  updateSurvey(q) {
    return request({
      url: '/project/task/survey/' + q.id,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  createQuestion(q) {
    return request({
      url: '/project/task/survey/question/',
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateQuestion(q) {
    return request({
      url: '/project/task/survey/question/' + q.id,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  deleteQuestion(qid) {
    return request({
      url: '/project/task/survey/question/delete/' + qid,
      method: 'post',
    })
  },
  exchangeQuestion(id1, id2) {
    return request({
      url: '/project/task/survey/question/exchange/' + id1 + '/' + id2,
      method: 'post',
    })
  }
}

export default SurveyApi