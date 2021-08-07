import request from './request'

const ProjectListApi = {
  getUserProjectList(mode, q) {
    return request({
      url: '/project-list/' + mode,
      params: q,
      method: 'get'
    })
  },
}

export default ProjectListApi