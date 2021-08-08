import request from "./request";
import qs from 'qs'

const MessageApi = {
  getUserMessages(query) {
    return request({
      url: `/message`,
      method: 'get',
      params: query
    })
  },
  readMessage(m) {
    return request({
      url: `/message/${m.id}/read`,
      method: 'post',
      data: qs.stringify(m)
    })
  },
  readAllMessage() {
    return request({
      url: `/message/read-all`,
      method: 'post',
    })
  },
  deleteMessage(m) {
    return request({
      url: `/message/${m.id}/delete`,
      method: 'post',
    })
  }
}

export default MessageApi