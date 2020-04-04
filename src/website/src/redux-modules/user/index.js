const userInitialState = {
    account: {},
    token: '',
}

function userReducer(state = userInitialState, action) {
    switch(action.type){
        case 'LOGIN':
            const { account, token } = action.payload
            return {
                ...state,
                account,
                token
            }
        case 'CREATE_ACCOUNT':
            const { createdAccount, createdAccountToken } = action.payload
            return {
                ...state,
                account: createdAccount,
                token: createdAccountToken,
            }
        default:
            return state
    }
}

export default userReducer