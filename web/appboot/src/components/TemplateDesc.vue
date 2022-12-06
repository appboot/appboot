<template>
  <div>
    <div class="title">Description</div>
    {{ desc }}
  </div>
</template>

<script setup lang="ts">
defineProps({
  desc: String,
});
</script>
