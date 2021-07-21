import request from "./request";
import qs from "qs";

export default {
  publishProject() {
    return request({
      url:'/teacher/publish',
      method: 'post',
    })
  },
}
