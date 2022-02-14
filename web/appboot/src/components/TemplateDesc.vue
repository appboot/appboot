<script setup>
defineProps({
  desc: String,
});
</script>

<template>
  <div>
    <div class="title">Description</div>
    {{ desc }}
  </div>
</template>
