import {trim} from "./Util";
import axios from "axios";

export let authConfig = {
  serverUrl: process.env.REACT_APP_CASDOOR_ENDPOINT,
  clientId: process.env.REACT_APP_CLIENT_ID,
  appName: process.env.REACT_APP_APP_NAME,
  organizationName: process.env.REACT_APP_CASDOOR_ORGANIZATION,
}

export function initAuthWithConfig(config) {
  authConfig = config;
}
export function getCasdoorService() {
  return axios.create({
    baseURL: authConfig.serverUrl,
    timeout: 10000,
    withCredentials: true
  })
}

export function getUser(teacherId) {
  let request = getCasdoorService()
  return request({
    url: '/api/get-user',
    method: 'get',
    params: {id: authConfig.organizationName + '/' + teacherId}
  })
}

export function getAccount(account) {
  let request = getCasdoorService()
  return request({
    url: `/account?access_token=${account.accessToken}`,
    method: 'get',
  })
}

export function getAuthorizeUrl() {
  const redirectUri = `${window.location.origin}/callback/`;
  const scope = "read";
  const state = authConfig.appName;
  return `${trim(authConfig.serverUrl)}/login/oauth/authorize?client_id=${authConfig.clientId}&response_type=code&redirect_uri=${redirectUri}&scope=${scope}&state=${state}`;
}

export function getUserProfileUrl(userName, account) {
  let param = "";
  if (account !== undefined && account !== null) {
    param = `?access_token=${account.accessToken}`;
  }
  return `${trim(authConfig.serverUrl)}/users/${authConfig.organizationName}/${userName}${param}`;
}

export function getMyProfileUrl(account) {
  let param = "";
  if (account !== undefined && account !== null) {
    param = `?access_token=${account.accessToken}`;
  }
  return `${trim(authConfig.serverUrl)}/account${param}`;
}
