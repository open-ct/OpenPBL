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