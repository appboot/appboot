// WSS_URL developed locally can be modified to "ws://127.0.0.1:8888/appboot"
// !!!Do not replace default value when build docker
// In start.zh we use the default value to occupy the place, to be replaced by the environment variable WS_URL, in order to achieve the purpose of use the docker environment variable to change the backed url
export const WSS_URL = "ws://ws.appboot.com:8888/appboot";