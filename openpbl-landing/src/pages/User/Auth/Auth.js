import {trim} from "./Util";
import axios from "axios";

export let authConfig = {
  category: "email",
  serverUrl: "http://121.196.149.183:8000", // your Casdoor URL, like the official one: https://door.casbin.com
  clientId: "9145af9a92d8194b5ede",         // your Casdoor OAuth Client ID
  appName: "openpbl",                       // your Casdoor application name, like: "app-built-in"
  organizationName: "openct",               // your Casdoor organization name, like: "built-in"
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
