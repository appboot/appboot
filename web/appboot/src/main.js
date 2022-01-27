import Vue from "vue";
import App from "./App.vue";

import { Button, Input, Form, Radio, InputNumber, message, Select, Icon, Tooltip, Switch, List } from "ant-design-vue";

Vue.use(Button);
Vue.use(Input);
Vue.use(Form);
Vue.use(Radio);
Vue.use(InputNumber);
Vue.use(Icon);
Vue.use(Select);
Vue.use(Tooltip);
Vue.use(Switch);
Vue.use(List);
Vue.prototype.$message = message;

Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");
