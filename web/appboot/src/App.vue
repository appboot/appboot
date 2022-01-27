<template>
  <div id="app">
    <Logo />
    <div id="creator" v-if="!finish">
      <Template @change="onTemplateChange" @onConfigChange="onConfigChange" />
      <TemplateDesc v-if="desc" :desc="desc" />
      <Params v-if="selectedTemplate" @change="onNameChange" :params="params" :paramsLength="paramsLength" />
      <Scripts id="scripts" v-if="showScripts" :beforeScripts="beforeScripts" :afterScripts="afterScripts" @onBeforeChange="onBeforeChange" @onAfterChange="onAfterChange" />
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
import Scripts from "./components/Scripts.vue";
import Success from "./components/Success.vue";
import { decodeParams, encodeParams } from "./params";
import { createApp } from "./api";
import download from "./download";

export default {
  name: "App",
  data() {
    return {
      name: "",
      desc: "",
      selectedTemplate: "",
      paramsLength: 0,
      params: [],
      beforeScripts: [],
      afterScripts: [],
      enableBefore: true,
      enableAfter: true,
      creating: false,
      finish: false
    };
  },
  computed: {
    showScripts: function() {
      return this.beforeScripts.length > 0 || this.afterScripts.length > 0;
    }
  },
  methods: {
    onTemplateChange(template) {
      this.selectedTemplate = template;
    },
    onConfigChange(configs) {
      const params = configs.parameters;
      this.desc = configs.desc;
      this.beforeScripts = configs.scripts.before ?? [];
      this.afterScripts = configs.scripts.after ?? [];
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
    onBeforeChange(value) {
      this.enableBefore = value;
    },
    onAfterChange(value) {
      this.enableAfter = value;
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
      var skipBeforeScripts = this.enableBefore ? "false" : "true";
      var skipAfterScripts = this.enableAfter ? "false" : "true";
      createApp(this.name, this.selectedTemplate, encodeParams(this.params), skipBeforeScripts, skipAfterScripts)
        .then(function(data) {
          that.creating = false;
          if (data.code == 0) {
            that.finish = true;
            if (data.path) {
              download(data.path, that.name + ".zip");
            }
          } else {
            that.$message.error(data.message);
          }
        })
        .catch(function(error) {
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
    }
  },
  components: {
    Logo,
    Template,
    Params,
    Scripts,
    Success,
    TemplateDesc
  }
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
