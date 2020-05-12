import Vue from "vue";
import App from "./App.vue";
import {
  Button,
  Input,
  Form,
  Radio,
  InputNumber,
  message,
  Icon
} from "ant-design-vue";

import hljs from 'highlight.js'
import 'highlight.js/styles/monokai-sublime.css'

Vue.use(Button);
Vue.use(Input);
Vue.use(Form);
Vue.use(Radio);
Vue.use(InputNumber);
Vue.use(Icon);
Vue.prototype.$message = message;

Vue.config.productionTip = false;

Vue.directive('highlight',function (el) {
  let blocks = el.querySelectorAll('pre code');
      blocks.forEach((block)=>{
      hljs.highlightBlock(block)
  })
})

new Vue({
  render: h => h(App)
}).$mount("#app");
