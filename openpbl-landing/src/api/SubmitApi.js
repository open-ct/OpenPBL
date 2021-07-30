import request from "./request";
import qs from "qs";

const SubmitApi = {
  createSubmit(pid, tid, data) {
    return request({
      url: `/project/${pid}/task/${tid}/submit`,
      method: 'post',
      data: qs.stringify(data, {arrayFormat: 'repeat'})
    })
  },
  updateSubmit(pid, tid, sid, data) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
}

export default SubmitApi
