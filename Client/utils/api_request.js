/*
此页面为发起request请求页面
*/
const shema = 'http://'
const domain='api.ydq6.com'
const urlPrefix= shema + domain+'/DietRegimen/client'

const userLogin = urlPrefix + '/user/userLogin'

module.exports = {
  domain:domain,
  urlPrefix:urlPrefix,
  userLogin:userLogin
}