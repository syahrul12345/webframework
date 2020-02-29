export const getLoginUrl = () => {
  let url = ""
  if (process.env.NODE_ENV == "production") {
      url = '/api/v1/login'
  }else{
      url = 'http://localhost:8004/api/v1/login'
  }
  return url
}
export const getCreateAccountUrl = () => {
  let url = ""
  if (process.env.NODE_ENV === "production") {
      url = '/api/v1/createAccount'
  }else{
    url = 'http://localhost:8004/api/v1/createAccount'
  }
  return url
}