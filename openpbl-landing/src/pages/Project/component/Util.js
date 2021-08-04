import moment from 'moment'

function FilterTime(time) {
  return moment(time).format('YYYY-MM-DD HH:mm:ss')
}

function FilterMoment(time) {
  return moment(time).fromNow()
}

function FormatSectionName(sectionName, chapterNumber, sectionNumber) {
  return `${chapterNumber + 1}.${sectionNumber + 1} ${sectionName}`
}

function FormatChapterName(chapterName, chapterNumber) {
  return `${chapterNumber + 1}. ${chapterName}`
}

export default {
  FilterTime,
  FilterMoment,
  FormatSectionName,
  FormatChapterName
}