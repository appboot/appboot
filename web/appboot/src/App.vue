<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import { ref, computed } from "vue";
import { message } from "ant-design-vue";
import Logo from "./components/Logo.vue";
import Template from "./components/Template.vue";
import TemplateDesc from "./components/TemplateDesc.vue";
import Params from "./components/Params.vue";
import Scripts from "./components/Scripts.vue";
import Success from "./components/Success.vue";
import { decodeParams, encodeParams } from "./params";
import { PlusOutlined } from "@ant-design/icons-vue";
import download from "./download";
import { createApp } from "./api";

var desc = ref("");
var name = ref("");
var selectedTemplate = ref("");
var paramsLength = ref(0);
var params = ref([]);
var beforeScripts = ref([]);
var afterScripts = ref([]);
var enableBefore = ref(true);
var enableAfter = ref(true);
var creating = ref(false);
var finish = ref(false);

const showScripts = computed(() => {
  return beforeScripts.value.length > 0 || afterScripts.value.length > 0;
});

function onTemplateChange(template) {
  selectedTemplate.value = template;
}

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
  var skipBeforeScripts = enableBefore.value ? "false" : "true";
  var skipAfterScripts = enableAfter.value ? "false" : "true";
  createApp(name.value, selectedTemplate.value, encodeParams(params.value), skipBeforeScripts, skipAfterScripts)
    .then(function (data) {
      creating.value = false;
      if (data.code == 0) {
        finish.value = true;
        if (data.path) {
          download(data.path, name.value + ".zip");
        }
      } else {
        message.error(data.message);
      }
    })
    .catch(function (error) {
      creating.value = false;
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
</script>

<template>
  <Logo />
  <div id="creator" v-if="!finish">
    <Template @change="onTemplateChange" @onConfigChange="onConfigChange" />
    <TemplateDesc v-if="desc" :desc="desc" />
    <Params v-if="selectedTemplate" @change="onNameChange" :params="params" :paramsLength="paramsLength" />
    <Scripts id="scripts" v-if="showScripts" :beforeScripts="beforeScripts" :afterScripts="afterScripts" @onBeforeChange="onBeforeChange" @onAfterChange="onAfterChange" />

    <a-button v-if="selectedTemplate" class="create-button" type="primary" :loading="creating" @click="onCreate">
      <template #icon><PlusOutlined /></template>
      Create
    </a-button>
  </div>
  <Success v-if="finish" :name="name" />
</template>

<style>
#app {
  display: flex;
  flex-direction: column;
}
#creator {
  display: flex;
  flex-direction: column;
  justify-content: left;
  margin-left: 20%;
  margin-right: 20%;
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
