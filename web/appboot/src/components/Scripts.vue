<template>
  <div>
    <div id="scripts">
      <div class="title">Scripts</div>
    </div>

    <div v-if="beforeScripts.length > 0">
      <div class="script">
        <div class="sub-title">Before</div>
        <a-switch default-checked @change="onBeforeChange" />
      </div>

      <a-list bordered :data-source="beforeScripts">
        <a-list-item slot="renderItem" slot-scope="item"> {{ item }} </a-list-item>
      </a-list>
    </div>

    <div v-if="afterScripts.length > 0">
      <div class="script">
        <div class="sub-title">After</div>
        <a-switch default-checked @change="onAfterChange" />
      </div>

      <a-list bordered :data-source="afterScripts">
        <a-list-item slot="renderItem" slot-scope="item"> {{ item }} </a-list-item>
      </a-list>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  props: {
    beforeScripts: Array,
    afterScripts: Array
  },
  methods: {
    onBeforeChange(checked) {
      this.$emit("onBeforeChange", checked);
    },
    onAfterChange(checked) {
      this.$emit("onAfterChange", checked);
    }
  }
};
</script>

<style>
.script {
  display: flex;
  flex-direction: row;
  align-items: center;
}
</style>
