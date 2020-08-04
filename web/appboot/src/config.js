// API_HOST develop locally can be modified to "http://127.0.0.1:8000"
// !!!Do not replace default value when build docker
// In start.zh we use the default value to occupy the place, to be replaced by the environment variable API_HOST, in order to achieve the purpose of use the docker environment variable to change the backed url
export const API_HOST = "http://api.appboot.com:8000";
// export const API_HOST = "http://127.0.0.1:8000";