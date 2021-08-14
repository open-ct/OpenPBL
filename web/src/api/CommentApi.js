import request from "./request";
import qs from 'qs'

const CommentApi = {
  getProjectComments(pid, q) {
    return request({
      url: `/project/${pid}/comments`,
      method: 'get',
      params: q
    })
  },
  createComment(pid, data) {
    return request({
      url: `/project/${pid}/comment`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  deleteComment(pid, cid) {
    return request({
      url: `/project/${pid}/comment/${cid}/delete`,
      method: 'post',
    })
  }
}

export default CommentApi