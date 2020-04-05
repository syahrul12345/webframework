import { combineReducers } from 'redux'

// import all required reducers
import userReducer from './user'

// RootReducer which combines all reducers
const rootReducer = combineReducers({
  user: userReducer,
})


export default rootReducer

