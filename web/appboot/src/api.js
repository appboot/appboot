import { API_URL } from "./config";
import axios from "axios";

axios.defaults.baseURL = API_URL;

export function getTemplates() {
  return new Promise((resolve, reject) => {
    axios
      .get("/templates")
      .then(function (response) {
        resolve(response.data.data ?? []);
      })
      .catch(function (error) {
        reject(error.toString());
      });
  });
}

export function getConfigs(template) {
  return new Promise((resolve, reject) => {
    axios
      .get("/configs/" + template)
      .then(function (response) {
        resolve(response.data.data);
      })
      .catch(function (error) {
        reject(error.toString());
      });
  });
}

export function updateTemplates() {
  return new Promise((resolve, reject) => {
    axios
      .put("/templates")
      .then(function (response) {
        resolve(response.data.data ?? []);
      })
      .catch(function (error) {
        reject(error.toString());
      });
  });
}

export function getTemplatesGitHash() {
  return new Promise((resolve, reject) => {
    axios
      .get("/templates/git_hash")
      .then(function (response) {
        resolve(response.data.data);
      })
      .catch(function (error) {
        reject(error.toString());
      });
  });
}

export function createApp(name, template, params, skipBeforeScripts, skipAfterScripts) {
  return new Promise((resolve, reject) => {
    const form = new FormData();
    form.append("name", name);
    form.append("template", template);
    form.append("params", params);
    form.append("skipBeforeScripts", skipBeforeScripts);
    form.append("skipAfterScripts", skipAfterScripts);

    axios
      .post("/app", form)
      .then(function (response) {
        if (response.data.data && response.data.data.length > 0) {
          const path = axios.defaults.baseURL + response.data.data;
          response.data.path = path;
        }
        resolve(response.data);
      })
      .catch(function (error) {
        reject(error.toString());
      });
  });
}
