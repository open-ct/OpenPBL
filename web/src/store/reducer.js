import {combineReducers} from 'redux-immutable'
import TypeReducer from './TypeReducer'

const reducer = combineReducers({
  userType: TypeReducer
})

export default reducer