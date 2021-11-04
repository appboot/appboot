<template>
  <div id="finish">
    <div class="finish-text">
      <a-icon class="icons-list" type="check-circle" theme="twoTone" twoToneColor="#52c41a" />
      Congratulations, the application {{ name }} was created successfully!
    </div>
  </div>
</template>

<script>
export default {
  name: "Success",
  props: {
    name: String,
  },
};
</script>

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