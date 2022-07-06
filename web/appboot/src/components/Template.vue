<script setup>
import { ReloadOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { ref } from "vue";
import { getTemplates, getConfigs, updateTemplates, getTemplatesGitHash } from "../api";

// emit
const emit = defineEmits(["change", "onConfigChange", "update"]);

// variable
const loading = ref(false);
const templates = ref([]);
const gitHash = ref("ssss");
const selectedTemplate = ref("");

getTemplates()
  .then(function (ts) {
    templates.value = ts;
  })
  .catch(function (error) {
    message.error(error);
  });

getTemplatesGitHash()
  .then(function (hash) {
    gitHash.value = hash;
  })
  .catch(function (error) {
    message.error(error);
  });

function onChange(e) {
  let value = e.target.value;
  emit("change", value);

  getConfigs(value)
    .then(function (configs) {
      emit("onConfigChange", configs);
    })
    .catch(function (error) {
      message.error(error);
      emit("onConfigChange", []);
    });
}

function onUpdate() {
  loading.value = true;
  emit("update");

  updateTemplates()
    .then(function (ts) {
      loading.value = false;
      templates.value = ts;

      getTemplatesGitHash()
        .then(function (hash) {
          gitHash.value = hash;
        })
        .catch(function (error) {
          message.error(error);
        });
    })
    .catch(function (error) {
      loading.value = false;
      message.error(error);
    });
}
</script>

<template>
  <div>
    <div id="template">
      <a-tooltip placement="top">
        <template #title>
          <span>git commit: {{ gitHash }}</span>
        </template>
        <div class="title">Template Commit: {{ gitHash }}</div>
      </a-tooltip>
      <a-button class="action-button" type="link" @click="onUpdate" :loading="loading">
        <template #icon><ReloadOutlined /></template>
      </a-button>
    </div>
    <div v-if="templates.length > 0">
      <a-radio-group v-model:value="selectedTemplate" button-style="solid" @change="onChange">
        <a-radio-button style="margin: 3px" v-for="(t, index) in templates" :key="index" :value="t">{{ t }}</a-radio-button>
      </a-radio-group>
    </div>
  </div>
</template>

<style>
#template {
  display: flex;
  flex-direction: row;
}
.gitHash {
  margin-bottom: 10px;
  font-size: medium;
}
</style>
