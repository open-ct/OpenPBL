import moment from 'moment'

function FilterTime(time) {
  return moment(time).format('YYYY-MM-DD HH:mm:ss')
}

function FilterMoment(time) {
  return moment(time).fromNow()
}

export default {
  FilterTime,
  FilterMoment
}