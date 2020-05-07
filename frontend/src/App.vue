<template>
  <div id="app">
    <div id="logo">
      <img alt="logo" src="./assets/logo.png" />
    </div>
    <div id="detial" v-if="!finish">
      <div class="title">Application Name</div>
      <a-input
        class="input"
        placeholder="your application name"
        v-model="name"
        @change="onNameChange"
      />
      
      <div class="title">Template</div>
      <div v-if="templates.length > 0">
        <a-radio-group class="radio" v-model="template" buttonStyle="solid">
          <a-radio-button v-for="(t, index) in templates" :key="index" :value=t>{{t}}</a-radio-button>
        </a-radio-group>
      </div>
      <a-input class="input" placeholder="" v-model="template" @change="onTemplateChange" />

      <div class="title">Params</div>
      <a-form layout="inline" style="margin-bottom: 15px">
        <div v-for="(param, index) in form.params" :key="index">
          <a-form-item label="key">
            <a-input v-model="param.key"></a-input>
          </a-form-item>

          <a-form-item label="value">
            <a-input v-model="param.value"></a-input>
          </a-form-item>

          <a-form-item>
            <a-icon
              class="dynamic-delete-button"
              type="minus-circle-o"
              @click="deleteParam(param, index)"
            />
          </a-form-item>
        </div>
      </a-form>
      <a-button type="dashed" style="width: 100%" @click="addParam">
        <a-icon type="plus" />Add param
      </a-button>

      <a-button
        class="createButton"
        type="primary"
        icon="plus"
        :loading="creating"
        @click="onCreate"
      >Create</a-button>
    </div>
    <div id="finish" v-if="finish">
      <div class="finish-text">
        <a-icon class="icons-list" type="check-circle" theme="twoTone" twoToneColor="#52c41a" />
        Congratulations, the application {{name}} was created successfully!
      </div>
    </div>
  </div>
</template>

<script>
import {} from "./string";
import { websocket, createSocket, sendGetTemplates, sendCreateApp, jsonParams} from "./websocket";

export default {
  name: "app",
  data() {
    return {
      creating: false,
      finish: false,
      name: "",
      template: "",
      templates: [],
      form: {
        params: []
      }
    };
  },
  methods: {
    init: function() {
      createSocket()
      websocket.onmessage = this.onmessage;
      websocket.onopen = this.onopen;
      websocket.onerror = this.onerror;
      websocket.onclose = this.onclose;
    },
    onNameChange() {
      this.name = this.name.trim();
    },
    onTemplateChange() {
      this.template = this.template.trim();
    },
    onCreate() {
      if (this.name.length < 1) {
        this.$message.error("application name cannot be empty.");
        return;
      }
      if (this.template.length < 1) {
        this.$message.error("template cannot be empty.");
        return;
      }

      if (!this.checkParams()) {
        this.$message.error("The key and value of all params cannot be empty.");
        return;
      }

      const params = jsonParams(this.form.params);

      this.creating = true;
      this.init();
      var that = this;
      setTimeout(function() {
        sendCreateApp(that.name, that.template, params)
      }, 1 * 1000);
    },
    onclose: function() {
      if (this.finish) {
        return;
      }
      this.creating = false;
    },
    onmessage: function(e) {
      const json = JSON.parse(e.data);
      if (json.method == "GetTemplates") {
        this.templates = json.data
      } else if (json.method == "CreateApp") {
        if (json.code < 500) {
          if (json.code == 0) {
            this.creating = false;
            this.finish = true;
          } else {
            this.$message.info(json.msg);
          }
        } else {
          this.$message.error(json.msg);
        }       
      }
    },
    onerror: function() {
      if (this.finish) {
        return;
      }
      this.$message.error(
        "websocket connection error, please check backend service"
      );
      this.creating = false;
      var that = this;
      setTimeout(function() {
        that.init();
      }, 5 * 1000);
    },
    checkParams() {
      var result = true;
      for (var j = 0, len = this.form.params.length; j < len; j++) {
        const param = this.form.params[j];
        if (param.key.length < 1 || param.value.length < 1) {
          result = false;
          break;
        }
      }
      return result;
    },
    addParam() {
      this.form.params.push({
        key: "",
        value: ""
      });
    },
    deleteParam(param, index) {
      this.form.params.splice(index, 1);
    }
  },
  mounted: function() {
    this.init();
    setTimeout(function() {
      sendGetTemplates()
    }, 100);
  },
  components: {}
};
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
}
#logo {
  display: flex;
  justify-content: center;
  margin-top: 2%;
  margin-bottom: 2%;
}
#finish {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 2%;
}
#detial {
  display: flex;
  flex-direction: column;
  justify-content: left;
  margin-left: 30%;
  margin-right: 30%;
}
.title {
  margin-bottom: 10px;
  font-size: larger;
  font-weight: bold;
}
.input {
  margin-bottom: 15px;
  height: 40px;
}
.radio {
  margin-bottom: 15px;
}
.createButton {
  width: 50%;
  height: 40px;
  align-self: center;
  font-size: medium;
  margin-top: 30px;
}
.addButton {
  width: 50px;
  height: 30px;
  align-self: center;
  font-size: medium;
  margin-top: 15px;
}
.icons-list {
  margin-right: 6px;
  font-size: 24px;
}
.finish-text {
  margin-bottom: 10px;
  font-size: larger;
  font-weight: bold;
  margin-top: 15px;
}
</style>
