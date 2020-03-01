const userInitialState = {
    email: '',
    token: 'dummyToken',
}

function userReducer(state = userInitialState, action) {
    switch(action.type){
        case 'LOGIN':
            const { token } = action.payload
            return {
                ...state,
                token: token
            }
        default:
            return state
    }
}

export default userReducer