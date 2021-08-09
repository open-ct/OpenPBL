import request from "./request";

const HomeApi = {
  getLatestProjects(num) {
    return request({
      url: '/home/projects/latest',
      method: 'get',
      params: {num: num}
    })
  }
}

export default HomeApi