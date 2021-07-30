import request from "./request";
import qs from "qs";

const ChapterApi = {
  getProjectChapters(id) {
    return request({
      url: `/project/${id}/chapters`,
      method: 'get',
    })
  },
  createProjectChapter(chapter) {
    return request({
      url: `/project/${chapter.projectId}/chapter`,
      method: 'post',
      data: qs.stringify(chapter)
    })
  },
  updateProjectChapter(chapter) {
    return request({
      url: `/project/${chapter.projectId}/chapter/${chapter.id}`,
      method: 'post',
      data: qs.stringify(chapter)
    })
  },

  deleteProjectChapter(c) {
    return request({
      url: `/project/${c.projectId}/chapter/${c.id}/delete`,
      method: 'post',
    })
  },
  exchangeProjectChapter(pid, id1, id2) {
    return request({
      url: `/project/${pid}/chapters/exchange`,
      method: 'post',
      data: qs.stringify({chapterId1: id1, chapterId2: id2})
    })
  },
}

export default ChapterApi