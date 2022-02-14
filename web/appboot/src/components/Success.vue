<script setup>
import { CheckCircleTwoTone } from "@ant-design/icons-vue";

defineProps({
  name: String,
});
</script>

<template>
  <div id="finish">
    <div class="finish-text">
      <check-circle-two-tone two-tone-color="#52c41a" style="font-size: large; margin-right: 8px" />
      Congratulations, {{ name }} was created successfully!
    </div>
  </div>
</template>

<style>
#finish {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 2%;
}
.finish-text {
  margin-bottom: 10px;
  font-size: larger;
  font-weight: bold;
  margin-top: 15px;
}
</style>
