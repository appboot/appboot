<template>
  <div>
    <div class="title">Description</div>
    {{ desc }}
  </div>
</template>

<script>
export default {
  props: {
    desc: String,
  },
};
</script>