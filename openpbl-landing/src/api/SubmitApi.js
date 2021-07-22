import request from "./request";
import qs from "qs";

export default {
  getSubmits(sid, pid) {
    return request({
      url:'/project/submit/' + pid + '/' + sid,
      method: 'get',
    })
  },
  createSubmits(data) {
    return request({
      url:'/project/submit',
      method: 'post',
      data: qs.stringify(data)
    })
  },
}
