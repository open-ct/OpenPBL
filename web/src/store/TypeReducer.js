import {fromJS} from 'immutable'

const defaultState = fromJS({
  userType: ""
})

export default (state = defaultState, action) => {
  switch (action.type) {
    case 'set':
      return fromJS({userType: action.userType})
    default:
      return state
  }
}