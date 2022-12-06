<template>
  <div>
    <div class="title">Description</div>
    {{ desc }}
  </div>
</template>

<script setup>
defineProps({
  desc: String,
});
</script>