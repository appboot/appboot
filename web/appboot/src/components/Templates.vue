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
    <div v-if="groups.length > 0">
      <a-tabs>
        <a-tab-pane v-for="g in groups" :key="g.id" :tab="g.desc">
          <a-radio-group v-model:value="selectedTemplate" button-style="solid" @change="onChange">
            <a-tooltip v-for="(t, index) in g.templates">
              <template #title>{{ t.desc }}</template>
              <a-radio-button style="margin: 3px" :key="index" :value="t">{{ t.id }}</a-radio-button>
            </a-tooltip>
          </a-radio-group>
        </a-tab-pane>
      </a-tabs>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ReloadOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { onMounted, ref } from "vue";
import { updateTemplates } from "../app/api";
import type { Template, TemplateGroup } from "../app/appboot";
import socket, { SokcetCMD, SokcetEvent } from "../app/ws";

// emit
const emit = defineEmits(["change", "update"]);

// variable
const loading = ref(false);
const groups = ref<TemplateGroup[]>([]);
const gitHash = ref("");
const selectedTemplate = ref<Template>();

onMounted(() => {
  loading.value = true;

  socket.on(SokcetEvent.open, () => {
    socket.getTemplates();
  });

  socket.on(SokcetEvent.message, (data: string) => {
    loading.value = false;

    let obj = JSON.parse(data);
    if (obj.cmd === SokcetCMD.getTemplates) {
      groups.value = obj.groups;
      gitHash.value = obj.hash;
    }
  });
});

function onChange(e: any) {
  let value = e.target.value;
  emit("change", value);
}

function onUpdate() {
  emit("update");

  loading.value = true;
  updateTemplates()
    .then(function (data: any) {
      groups.value = data.groups;
      gitHash.value = data.hash;
    })
    .catch(function (error) {
      message.error(error);
    })
    .finally(() => {
      loading.value = false;
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
