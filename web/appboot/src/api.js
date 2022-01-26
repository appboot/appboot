const axios = require("axios").default;
import { API_URL } from "./config";

axios.defaults.baseURL = API_URL;

export function getTemplates() {
  return new Promise((resolve, reject) => {
    axios
      .get("/templates")
      .then(function(response) {
        resolve(response.data.data);
      })
      .catch(function(error) {
        reject(error.toString());
      });
  });
}

export function getConfigs(template) {
  return new Promise((resolve, reject) => {
    axios
      .get("/configs/" + template)
      .then(function(response) {
        resolve(response.data.data);
      })
      .catch(function(error) {
        reject(error.toString());
      });
  });
}

export function updateTemplates() {
  return new Promise((resolve, reject) => {
    axios
      .put("/templates")
      .then(function(response) {
        resolve(response.data.data);
      })
      .catch(function(error) {
        reject(error.toString());
      });
  });
}

export function getTemplatesGitHash() {
  return new Promise((resolve, reject) => {
    axios
      .get("/templates/git_hash")
      .then(function(response) {
        resolve(response.data.data);
      })
      .catch(function(error) {
        reject(error.toString());
      });
  });
}

export function createApp(name, template, params) {
  return new Promise((resolve, reject) => {
    const form = new FormData();
    form.append("name", name);
    form.append("template", template);
    form.append("params", params);

    axios
      .post("/app", form)
      .then(function(response) {
        resolve(response.data.message);
      })
      .catch(function(error) {
        reject(error.toString());
      });
  });
}
