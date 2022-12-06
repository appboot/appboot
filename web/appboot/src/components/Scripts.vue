<template>
  <div>
    <div id="scripts">
      <div class="title">Scripts</div>
    </div>

    <div v-if="beforeScripts && beforeScripts.length > 0">
      <div class="script">
        <div class="sub-title">Before</div>
        <a-switch v-model:checked="beforeChecked" @change="onBeforeChange" />
      </div>

      <a-list bordered :data-source="beforeScripts">
        <template #renderItem="{ item }">
          <a-list-item>{{ item }}</a-list-item>
        </template>
      </a-list>
    </div>

    <div v-if="afterScripts && afterScripts.length > 0">
      <div class="script">
        <div class="sub-title">After</div>
        <a-switch v-model:checked="afterChecked" @change="onAfterChange" />
      </div>

      <a-list bordered :data-source="afterScripts">
        <template #renderItem="{ item }">
          <a-list-item>{{ item }}</a-list-item>
        </template>
      </a-list>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

defineProps({
  beforeScripts: Array<string>,
  afterScripts: Array<string>,
});

const beforeChecked = ref(true);
const afterChecked = ref(true);

const emit = defineEmits(["onBeforeChange", "onAfterChange"]);

function onBeforeChange(checked: boolean) {
  emit("onBeforeChange", checked);
}

function onAfterChange(checked: boolean) {
  emit("onAfterChange", checked);
}
</script>

<style>
.script {
  display: flex;
  flex-direction: row;
  align-items: center;
}
</style>
