import request from "./request";
import qs from 'qs'

const ResourceApi = {
  getResource(id) {
    return request({
      url:'/resource/' + id,
      method: 'get',
    })
  },
  createResource(q) {
    return request({
      url:'/resource',
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateResource(q) {
    return request({
      url:'/resource/' + q.id,
      method: 'post',
      data: qs.stringify(q)
    })
  },
  updateResourceContent(id, content) {
    return request({
      url:'/resource/' + id,
      method: 'post',
      data: qs.stringify({content: content})
    })
  },

}

export default ResourceApi