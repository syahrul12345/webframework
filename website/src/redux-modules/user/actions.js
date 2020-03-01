export const LoginAction = (token) => {
    return {
        type: 'LOGIN',
        payload: {
            token
        }
    }
} 