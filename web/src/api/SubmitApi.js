import request from "./request";
import qs from "qs";

const SubmitApi = {
  createSubmit(pid, tid, data) {
    return request({
      url: `/project/${pid}/task/${tid}/submit`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  updateSubmit(pid, tid, sid, data) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  getSubmitFiles(pid, tid, sid) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/files`,
      method: 'get',
    })
  },
  uploadSubmitFile(pid, tid, sid, data) {
    if (sid === "") {
      sid = "0"
    }
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/file`,
      method: 'post',
      data: qs.stringify(data)
    })
  },
  updateSubmitFile(pid, tid, sid, file) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/file/${file.id}/update`,
      method: 'post',
      data: qs.stringify(file)
    })
  },
  deleteSubmitFile(pid, tid, sid, fid) {
    return request({
      url: `/project/${pid}/task/${tid}/submit/${sid}/file/${fid}/delete`,
      method: 'post'
    })
  }
}

export default SubmitApi
