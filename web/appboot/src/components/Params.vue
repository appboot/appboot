<template>
  <div>
    <div id="params">
      <div class="title">Params</div>
    </div>

    <a-form layout="inline" style="margin-bottom: 15px">
      <!-- Name -->
      <a-form-item class="param-key">
        <a-tooltip placement="top">
          <template #title>
            <span>Project Name</span>
          </template>
          <label>Name</label>
        </a-tooltip>
      </a-form-item>

      <a-form-item>
        <a-input class="input-string" @change="OnNameChange"></a-input>
      </a-form-item>

      <!-- Other Params -->
      <div v-for="(param, index) in params" :key="index">
        <!-- Key -->
        <a-form-item label="" class="param-key">
          <a-tooltip placement="top" v-if="param.tip">
            <template #title>
              <span>{{ param.tip }}</span>
            </template>
            <label class="input-string">{{ param.key }}</label>
          </a-tooltip>
          <div v-if="!param.tip">
            <label class="input-string">{{ param.key }}</label>
          </div>
        </a-form-item>

        <!-- Value -->
        <a-form-item label="" v-if="param.type == 'string'">
          <a-input class="input-string" v-model="param.value"></a-input>
        </a-form-item>

        <a-form-item v-if="param.type == 'int'">
          <a-input-number :min="param.min" :max="param.max" v-model="param.value" />
        </a-form-item>

        <a-form-item v-if="param.type == 'float'">
          <a-input-number :min="param.min" :max="param.max" v-model="param.value" />
        </a-form-item>

        <a-form-item v-if="param.type == 'select'">
          <a-select v-model="param.value" style="width: 240px">
            <a-select-option v-for="option in param.options" :key="option">
              {{ option }}
            </a-select-option>
          </a-select>
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
    paramsLength: Number
  },
  mounted() {},
  methods: {
    OnNameChange(e) {
      let value = e.target.value;
      this.$emit("change", value);
    }
  }
};
</script>

<style>
#params {
  display: flex;
  flex-direction: row;
}

.input-string {
  width: 400px;
}

.param-key {
  width: fit-content;
  min-width: 110px;
}
</style>
