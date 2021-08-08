import request from "./request";
import qs from "qs";


const SectionApi = {
  getChapterSections(item) {
    return request({
      url: `/project/${item.projectId}/chapter/${item.id}/sections`,
      method: 'get',
    })
  },
  createChapterSection(pid, section) {
    return request({
      url: `/project/${pid}/chapter/${section.chapterId}/section`,
      method: 'post',
      data: qs.stringify(section)
    })
  },
  updateChapterSection(pid, section) {
    return request({
      url: `/project/${pid}/chapter/${section.chapterId}/section/${section.id}`,
      method: 'post',
      data: qs.stringify(section)
    })
  },
  updateSectionsMinute(pid, sections) {
    return request({
      url: `/project/${pid}/sections-minute`,
      method: 'post',
      data: qs.stringify(sections)
    })
  },
  deleteChapterSection(pid, s) {
    return request({
      url: `/project/${pid}/chapter/${s.chapterId}/section/${s.id}/delete`,
      method: 'post',
      data: qs.stringify(s)
    })
  },
  exchangeChapterSection(chapter, id1, id2) {
    return request({
      url: `/project/${chapter.projectId}/chapter/${chapter.id}/sections/exchange`,
      method: 'post',
      data: qs.stringify({
        sectionId1: id1,
        sectionId2: id2
      })
    })
  },
  getSectionDetail(id, pid) {
    return request({
      url: `/project/${pid}/section/${id}`,
      method: 'get',
    })
  },
}

export default SectionApi