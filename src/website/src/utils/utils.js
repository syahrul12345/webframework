export const getLoginUrl = () => {
  let url = ""
  if (process.env.NODE_ENV === "production") {
      url = '/api/v1/loginAccount'
  }else{
      url = 'http://localhost:8005/api/v1/loginAccount'
  }
  return url
}
export const getCreateAccountUrl = () => {
  let url = ""
  if (process.env.NODE_ENV === "production") {
      url = '/api/v1/createAccount'
  }else{
    url = 'http://localhost:8005/api/v1/createAccount'
  }
  return url
}
export const getVerifyTokenUrl = () => {
  let url = ""
  if (process.env.NODE_ENV === "production") {
      url = '/api/v1/verify'
  } else{
    url = 'http://localhost:8005/api/v1/verify'
  }
  return url
}
export const getResetPasswordUrl = () => {
  let url = ""
  if (process.env.NODE_ENV === "production") {
    url = '/api/v1/forgetPassword'
  } else {
    url = 'http://localhost:8005/api/v1/forgetPassword'
  }
  return url
}