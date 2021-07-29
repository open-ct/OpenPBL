import request from "./request";
import qs from "qs";

const ChapterApi = {
  getProjectChapters(id) {
    return request({
      url: '/project/chapters/' + id,
      method: 'get',
    })
  },
  createProjectChapter(chapter) {
    return request({
      url: '/project/chapter',
      method: 'post',
      data: qs.stringify(chapter)
    })
  },
  updateProjectChapter(chapter) {
    return request({
      url: '/project/chapter/' + chapter.id,
      method: 'post',
      data: qs.stringify(chapter)
    })
  },

  deleteProjectChapter(id) {
    return request({
      url: '/project/chapter/delete/' + id,
      method: 'post',
    })
  },
  exchangeProjectChapter(id1, id2) {
    return request({
      url: '/project/chapter/exchange/' + id1 + '/' + id2,
      method: 'post',
    })
  },
}

export default ChapterApi