<template>
  <div>
    <div id="template">
      <div class="title">Template</div>
      <a-button class="action-button" icon="reload" type="link" @click="onUpdate" :loading="loading"></a-button>
    </div>
    <div v-if="templates.length > 0">
      <a-radio-group class="radio" buttonStyle="solid" @change="onChange">
        <a-radio-button v-for="(t, index) in templates" :key="index" :value="t">{{ t }}</a-radio-button>
      </a-radio-group>
    </div>
  </div>
</template>

<script>
import { getTemplates, getConfigs, updateTemplates } from "../api";

export default {
  name: "Template",
  data() {
    return {
      loading: false,
      templates: [],
    };
  },
  props: {},
  mounted() {
    var that = this;
    getTemplates()
      .then(function (templates) {
        that.templates = templates;
      })
      .catch(function (error) {
        that.$message.error(error);
      });
  },
  methods: {
    onChange(e) {
      let value = e.target.value;
      this.$emit("change", value);

      var that = this;
      getConfigs(value)
        .then(function (configs) {
          that.$emit("onConfigChange", configs);
        })
        .catch(function (error) {
          that.$message.error(error);
          that.$emit("onConfigChange", []);
        });
    },
    onUpdate() {
      this.loading = true;
      this.$emit("update");

      var that = this;
      updateTemplates()
        .then(function (templates) {
          that.loading = false;
          that.templates = templates;
        })
        .catch(function (error) {
          that.loading = false;
          that.$message.error(error);
        });
    },
  },
};
</script>

<style>
#template {
  display: flex;
  flex-direction: row;
}
</style>