<script setup>
defineProps({
  params: Array,
  paramsLength: Number,
});

const emit = defineEmits(["change"]);
const labelCol = { span: 4 };
const wrapperCol = { span: 14 };

function onNameChange(e) {
  let value = e.target.value;
  emit("change", value);
}
</script>

<template>
  <div>
    <div id="params">
      <div class="title">Params</div>
    </div>

    <a-form :label-col="labelCol" :wrapper-col="wrapperCol" labelAlign="left">
      <a-tooltip placement="topLeft">
        <template #title>
          <span>Project Name</span>
        </template>
        <a-form-item label="Name" :colon="false" style="margin-bottom: 8px">
          <a-input class="input-string" @change="onNameChange"></a-input>
        </a-form-item>
      </a-tooltip>

      <!-- Other Params -->
      <div v-for="(param, index) in params" :key="index">
        <a-tooltip placement="topLeft" v-if="param.tip">
          <template #title>
            <span>{{ param.tip }}</span>
          </template>
          <a-form-item :label="param.key" :colon="false" style="margin-bottom: 8px">
            <a-input class="input-string" v-model:value="param.value" v-if="param.type == 'string'"></a-input>
            <a-input-number :min="param.min" :max="param.max" v-model:value="param.value" v-if="param.type == 'int'" />
            <a-input-number :min="param.min" :max="param.max" v-model:value="param.value" v-if="param.type == 'float'" />
            <a-select v-model:value="param.value" style="width: 240px" v-if="param.type == 'select'">
              <a-select-option v-for="option in param.options" :key="option">
                {{ option }}
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-tooltip>
      </div>
    </a-form>
  </div>
</template>

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
