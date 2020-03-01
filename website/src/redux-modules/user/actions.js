export const LoginAction = (token) => {
    return {
        type: 'LOGIN',
        payload: {
            token
        }
    }
}
export const CreateAccountAction = (account,token) => {
    return {
        type: 'CREATE_ACCOUNT',
        payload: {
            createdAccount: account,
            createdAccountToken: token,
        }
    }
}