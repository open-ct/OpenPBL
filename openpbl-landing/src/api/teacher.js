import request from "./request";
import qs from "qs";

export default {
  publishProject(pid) {
    return request({
      url:'/teacher/publish',
      method: 'post',
      data: qs.stringify({pid: pid})
    })
  },
}
