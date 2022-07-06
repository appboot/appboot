<script setup>
import { PlusOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { computed, ref, watch } from "vue";
import { createApp } from "./api";
import Logo from "./components/Logo.vue";
import Params from "./components/Params.vue";
import Scripts from "./components/Scripts.vue";
import Success from "./components/Success.vue";
import Template from "./components/Template.vue";
import TemplateDesc from "./components/TemplateDesc.vue";
import download from "./download";
import { decodeParams, encodeParams } from "./params";

const current = ref(0);
const desc = ref("");
const name = ref("");
const selectedTemplate = ref("");
const paramsLength = ref(0);
const params = ref([]);
const beforeScripts = ref([]);
const afterScripts = ref([]);
const enableBefore = ref(true);
const enableAfter = ref(true);
const creating = ref(false);
const createErr = ref(false);

const showScripts = computed(() => {
  return beforeScripts.value.length > 0 || afterScripts.value.length > 0;
});

function onTemplateChange(template) {
  selectedTemplate.value = template;
  current.value = 1;
}

watch(current, () => {
  if (current.value === 0) {
    creating.value = false;
    createErr.value = false;
  }
});

function onConfigChange(configs) {
  const ps = configs.parameters;
  desc.value = configs.desc;
  beforeScripts.value = configs.scripts.before ?? [];
  afterScripts.value = configs.scripts.after ?? [];
  if (ps) {
    params.value = decodeParams(ps);
    paramsLength.value = params.length;
  } else {
    params.value = [];
    paramsLength.value = 0;
  }
}

function onNameChange(value) {
  name.value = value;
}

function onBeforeChange(value) {
  enableBefore.value = value;
}

function onAfterChange(value) {
  enableAfter.value = value;
}

function onCreate() {
  if (selectedTemplate.value.length < 1) {
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
  createApp(name.value, selectedTemplate.value, encodeParams(params.value), skipBeforeScripts, skipAfterScripts)
    .then(function (data) {
      creating.value = false;
      if (data.code == 0) {
        current.value = 2;
        if (data.path) {
          download(data.path, name.value + ".zip");
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
    const param = params.value[j];
    if (param.key.length < 1 || param.value.length < 1) {
      result = false;
      break;
    }
  }
  return result;
}

function stepOneDesc() {
  var defaultValue = "Select a template";
  if (current.value === 0) {
    return defaultValue;
  }
  return selectedTemplate.value ? "Selected: " + selectedTemplate.value : defaultValue;
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

    <div id="creator" v-if="current < 2">
      <Template @change="onTemplateChange" @onConfigChange="onConfigChange" v-if="current === 0" />

      <div style="display: flex; flex-direction: column" v-if="current === 1">
        <TemplateDesc v-if="desc" :desc="desc" />
        <Params v-if="selectedTemplate" @change="onNameChange" :params="params" :paramsLength="paramsLength" />
        <Scripts id="scripts" v-if="showScripts" :beforeScripts="beforeScripts" :afterScripts="afterScripts" @onBeforeChange="onBeforeChange" @onAfterChange="onAfterChange" />

        <a-button v-if="selectedTemplate" class="create-button" type="primary" :loading="creating" @click="onCreate">
          <template #icon><PlusOutlined /></template>
          Create
        </a-button>
      </div>
    </div>

    <Success v-if="current === 2" :name="name" />
  </div>
</template>

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
