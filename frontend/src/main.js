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

Vue.use(Button);
Vue.use(Input);
Vue.use(Form);
Vue.use(Radio);
Vue.use(InputNumber);
Vue.use(Icon);
Vue.prototype.$message = message;

Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");
