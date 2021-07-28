import request from "./request";
import qs from "qs";

export default {
  createSubmit(tid, data) {
    return request({
      url: `/project/task/${tid}/submit`,
      method: 'post',
      data: qs.stringify(data, {arrayFormat: 'repeat'})
    })
  },
  updateSubmit(tid, sid, data) {
    return request({
      url: `/project/task/${tid}/submit/${sid}`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
}
