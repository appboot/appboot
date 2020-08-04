<template>
  <div>
    <div id="params">
      <div class="title">Params</div>
      <a-button class="action-button" icon="plus" type="link" @click="addParam"></a-button>
    </div>

    <a-form layout="inline" style="margin-bottom: 15px">
      <!-- Name -->
      <a-form-item label="key">
        <a-input default-value="Name" disabled="true"></a-input>
      </a-form-item>

      <a-form-item label="value">
        <a-input @change="OnNameChange"></a-input>
      </a-form-item>

      <!-- Other Params -->
      <div v-for="(param, index) in params" :key="index">
        <a-form-item label="key">
          <a-input v-model="param.key"></a-input>
        </a-form-item>

        <a-form-item label="value" v-if="param.type == 'string'">
          <a-input v-model="param.value"></a-input>
        </a-form-item>

        <a-form-item label="value" v-if="param.type == 'int'">
          <a-input-number
            :min="param.min"
            :max="param.max"
            :placeholder="param.value"
            v-model="param.value"
          />
        </a-form-item>

        <a-form-item label="value" v-if="param.type == 'float'">
          <a-input-number
            :min="param.min"
            :max="param.max"
            :placeholder="param.value"
            v-model="param.value"
          />
        </a-form-item>

        <a-form-item label="value" v-if="param.type == 'select'">
          <a-select v-model="param.value" @change="handleChange" style="width: 150px">
            <a-select-option v-for="option in param.options" :key="option">
              {{ option }}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-form-item v-if="index >= paramsLength">
          <a-icon
            class="dynamic-delete-button"
            type="minus-circle-o"
            @click="deleteParam(param, index)"
          />
        </a-form-item>
      </div>
    </a-form>
  </div>
</template>

<script>
export default {
  name: "Params",
  data() {
    return {};
  },
  props: {
    params: Array,
    paramsLength: Number,
  },
  mounted() {},
  methods: {
    OnNameChange(e) {
      let value = e.target.value;
      this.$emit("change", value);
    },
    deleteParam(param, index) {
      this.params.splice(index, 1);
    },
    addParam() {
      this.params.push({
        key: "",
        type: "string",
        value: "",
      });
    },
  },
};
</script>

<style>
#params {
  display: flex;
  flex-direction: row;
}
</style>