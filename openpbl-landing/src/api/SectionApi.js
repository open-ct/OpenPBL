import request from "./request";
import qs from "qs";


export default {
  getChapterSections(id) {
    return request({
      url:'/project/chapter/sections/' + id,
      method: 'get',
    })
  },
  createChapterSection(section) {
    return request({
      url:'/project/chapter/section',
      method: 'post',
      data: qs.stringify(section)
    })
  },
  updateChapterSection(section) {
    return request({
      url:'/project/chapter/section/' + section.id,
      method: 'post',
      data: qs.stringify(section)
    })
  },
  deleteChapterSection(id) {
    return request({
      url:'/project/chapter/section/delete/' + id,
      method: 'post',
    })
  },
  exchangeChapterSection(id1, id2) {
    return request({
      url:'/project/chapter/section/exchange/' + id1 + '/' + id2,
      method: 'post',
    })
  },
  getSectionDetail(id) {
    return request({
      url:'/project/chapter/section/' + id,
      method: 'get',
    })
  },
}
