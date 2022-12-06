<template>
  <div>
    <div id="template">
      <a-tooltip placement="top">
        <template #title>
          <span>git commit: {{ gitHash }}</span>
        </template>
        <div class="title">Templates</div>
      </a-tooltip>
      <a-tooltip placement="right">
        <template #title>
          <span>update templates</span>
        </template>
        <a-button class="action-button" type="link" @click="onUpdate" :loading="loading">
          <template #icon><ReloadOutlined /></template>
        </a-button>
      </a-tooltip>
    </div>
    <div v-if="templates.length > 0">
      <a-radio-group v-model:value="selectedTemplate" button-style="solid" @change="onChange">
        <a-tooltip v-for="(t, index) in templates">
          <template #title>{{t.desc}}</template>
            <a-radio-button style="margin: 3px" :key="index" :value="t">{{ t.name }}</a-radio-button>
          </a-tooltip>
      </a-radio-group>
    </div>
  </div>
</template>

<script setup>
import { ReloadOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { ref } from "vue";
import { getTemplates, updateTemplates } from "../api";

// emit
const emit = defineEmits(["change", "update"]);

// variable
const loading = ref(false);
const templates = ref([]);
const gitHash = ref("");
const selectedTemplate = ref("");

getTemplates()
  .then(function (ts) {
    templates.value = ts.templates;
    gitHash.value = ts.hash;
  })
  .catch(function (error) {
    message.error(error);
  });

function onChange(e) {
  let value = e.target.value;
  emit("change", value);
}

function onUpdate() {
  loading.value = true;
  emit("update");

  updateTemplates()
    .then(function (ts) {
      loading.value = false;

      templates.value = ts.templates;
      gitHash.value = ts.hash;
    })
    .catch(function (error) {
      loading.value = false;
      message.error(error);
    });
}
</script>

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
