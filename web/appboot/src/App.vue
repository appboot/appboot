<template>
  <div class="container">
    <Logo class="logo" />

    <div class="steps">
      <a-steps v-model:current="current">
        <a-step title="Step 1" :description="stepOneDesc()" />
        <a-step title="Step 2" description="Create an application" :status="stepTwoStatus()" disabled />
        <a-step title="Step 3" description="See the result" disabled />
      </a-steps>
    </div>

    <div id="creator" v-show="current < 2">
      <Templates @change="onTemplateChange" v-show="current === 0" />

      <div style="display: flex; flex-direction: column" v-show="current === 1 && selectedTemplate">
        <TemplateDesc :desc="selectedTemplate ? selectedTemplate.desc : ''" />
        <Params @change="onNameChange" :params="params" :paramsLength="paramsLength" />
        <Scripts id="scripts" v-if="showScripts" :beforeScripts="beforeScripts" :afterScripts="afterScripts" @onBeforeChange="onBeforeChange" @onAfterChange="onAfterChange" />

        <a-button class="create-button" type="primary" :loading="creating" @click="onCreate">
          <template #icon><PlusOutlined /></template>
          Create
        </a-button>
      </div>
    </div>

    <Success v-show="current === 2" :name="name" />
  </div>
</template>

<script setup lang="ts">
import { PlusOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { computed, ref, watch } from "vue";
import { createApp, getConfigs } from "./app/api";
import Logo from "./components/Logo.vue";
import Params from "./components/Params.vue";
import Scripts from "./components/Scripts.vue";
import Success from "./components/Success.vue";
import Templates from "./components/Templates.vue";
import TemplateDesc from "./components/TemplateDesc.vue";
import download from "./app/download";
import { decodeParams, encodeParams } from "./app/params";
import type { Parameter, Template } from "./app/appboot";
import { API_URL } from "./app/config";

const current = ref(0);
const name = ref("");
const selectedTemplate = ref<Template>();
const paramsLength = ref(0);
const params = ref<Parameter[]>([]);
const beforeScripts = ref<string[]>([]);
const afterScripts = ref<string[]>([]);
const enableBefore = ref(true);
const enableAfter = ref(true);
const creating = ref(false);
const createErr = ref(false);

const showScripts = computed(() => {
  return (beforeScripts.value && beforeScripts.value.length > 0) || (afterScripts.value && afterScripts.value.length > 0);
});

async function onTemplateChange(template: Template) {
  selectedTemplate.value = template;
  current.value = 1;

  try {
    const configs = await getConfigs(template.id);
    const ps = configs.parameters;
    beforeScripts.value = configs.scripts.before ?? [];
    afterScripts.value = configs.scripts.after ?? [];
    if (ps) {
      params.value = decodeParams(ps);
      paramsLength.value = params.value.length;
    } else {
      params.value = [];
      paramsLength.value = 0;
    }
  } catch (error) {
    message.error("get template config failed." + error);
  }
}

watch(current, () => {
  if (current.value === 0) {
    creating.value = false;
    createErr.value = false;
  }
});

function onNameChange(value: string) {
  name.value = value;
}

function onBeforeChange(value: boolean) {
  enableBefore.value = value;
}

function onAfterChange(value: boolean) {
  enableAfter.value = value;
}

function onCreate() {
  if (!selectedTemplate.value || selectedTemplate.value.id.length < 1) {
    message.error("template cannot be empty.");
    return;
  }
  if (name.value.length < 1) {
    message.error("name cannot be empty.");
    return;
  }
  if (!checkParams()) {
    message.error("the key and value of all params cannot be empty.");
    return;
  }

  creating.value = true;
  createErr.value = false;
  var skipBeforeScripts = enableBefore.value ? "false" : "true";
  var skipAfterScripts = enableAfter.value ? "false" : "true";
  createApp(name.value, selectedTemplate.value.id, encodeParams(params.value), skipBeforeScripts, skipAfterScripts)
    .then(function (data: any) {
      creating.value = false;
      if (data.code == 0) {
        current.value = 2;
        if (data.path) {
          download(API_URL + data.path, name.value + ".zip");
        }
      } else {
        createErr.value = true;
        message.error(data.message);
      }
    })
    .catch(function (error) {
      creating.value = false;
      createErr.value = true;
      message.error(error);
    });
}

function checkParams() {
  var result = true;
  for (var j = 0, len = params.value.length; j < len; j++) {
    const p = params.value[j];
    if (!isValidParameter(p)) {
      result = false;
      break;
    }
  }
  return result;
}

function isValidParameter(p: Parameter) {
  return p.key.length > 0 && p.value && p.value.toString().length > 0;
}

function stepOneDesc() {
  var defaultValue = "Select a template";
  if (current.value === 0) {
    return defaultValue;
  }
  return selectedTemplate.value ? "Selected: " + selectedTemplate.value.id : defaultValue;
}

function stepTwoStatus() {
  if (current.value < 1) {
    return "wait";
  } else if (current.value === 1) {
    if (createErr.value) {
      return "error";
    }
    return "process";
  } else if (current.value > 1) {
    return "finish";
  }
  return "wait";
}
</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
}

#app {
  display: flex;
  flex-direction: column;
}

.logo {
  width: 100%;
}

.steps {
  width: 70%;
}

#creator {
  display: flex;
  flex-direction: column;
  justify-content: left;
  padding: 10px;
  width: 70%;
}

.title {
  margin-bottom: 10px;
  margin-top: 10px;
  font-size: larger;
  font-weight: bold;
}
.sub-title {
  margin-bottom: 10px;
  margin-top: 10px;
  margin-right: 10px;
  font-size: me;
  font-weight: bold;
}
.create-button {
  width: 50%;
  height: 40px;
  align-self: center;
  font-size: medium;
  margin-top: 30px;
  margin-bottom: 40px;
}
.action-button {
  margin-left: 10px;
  align-self: center;
}
</style>
