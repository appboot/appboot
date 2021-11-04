<template>
  <div id="app">
    <Logo />
    <div id="creator" v-if="!finish">
      <Template @change="onTemplateChange" @onConfigChange="onConfigChange" />
      <TemplateDesc v-if="desc" :desc="desc" />
      <Params v-if="selectedTemplate" @change="onNameChange" :params="params" :paramsLength="paramsLength" />
      <a-button v-if="selectedTemplate" class="create-button" type="primary" icon="plus" :loading="creating" @click="onCreate">Create</a-button>
    </div>
    <Success v-if="finish" :name="name" />
  </div>
</template>

<script>
import Logo from "./components/Logo.vue";
import Template from "./components/Template.vue";
import TemplateDesc from "./components/TemplateDesc.vue";
import Params from "./components/Params.vue";
import Success from "./components/Success.vue";
import { decodeParams, encodeParams } from "./params";
import { createApp } from "./api";

export default {
  name: "App",
  data() {
    return {
      name: "",
      desc: "",
      selectedTemplate: "",
      paramsLength: 0,
      params: [],
      creating: false,
      finish: false,
    };
  },
  methods: {
    onTemplateChange(template) {
      this.selectedTemplate = template;
    },
    onConfigChange(configs) {
      const params = configs.parameters;
      this.desc = configs.desc;
      if (params) {
        this.params = decodeParams(params);
        this.paramsLength = this.params.length;
      } else {
        this.params = [];
        this.paramsLength = 0;
      }
    },
    onNameChange(value) {
      this.name = value;
    },
    onCreate() {
      if (this.name.length < 1) {
        this.$message.error("name cannot be empty.");
        return;
      }
      if (this.selectedTemplate.length < 1) {
        this.$message.error("template cannot be empty.");
        return;
      }
      if (!this.checkParams()) {
        this.$message.error("the key and value of all params cannot be empty.");
        return;
      }

      var that = this;
      this.creating = true;
      createApp(this.name, this.selectedTemplate, encodeParams(this.params))
        .then(function () {
          that.creating = false;
          that.finish = true;
        })
        .catch(function (error) {
          that.creating = false;
          that.$message.error(error);
        });
    },
    checkParams() {
      var result = true;
      for (var j = 0, len = this.params.length; j < len; j++) {
        const param = this.params[j];
        if (param.key.length < 1 || param.value.length < 1) {
          result = false;
          break;
        }
      }
      return result;
    },
  },
  components: {
    Logo,
    Template,
    Params,
    Success,
    TemplateDesc,
  },
};
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
}
#creator {
  display: flex;
  flex-direction: column;
  justify-content: left;
  margin-left: 30%;
  margin-right: 30%;
}
.title {
  margin-bottom: 10px;
  margin-top: 10px;
  font-size: larger;
  font-weight: bold;
}
.create-button {
  width: 50%;
  height: 40px;
  align-self: center;
  font-size: medium;
  margin-top: 30px;
}
.action-button {
  margin-left: 10px;
  align-self: center;
}
</style>
