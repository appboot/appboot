import { createApp } from "vue";
import App from "./App.vue";
import { Button, Input, Form, Radio, InputNumber, message, Select, Tooltip, Switch, List, Steps } from "ant-design-vue";

import "ant-design-vue/es/message/style/css"; //vite只能用 ant-design-vue/es 而非 ant-design-vue/lib

const app = createApp(App);

app.use(Button).use(Input).use(Form).use(Radio).use(InputNumber).use(Select).use(Tooltip).use(Switch).use(List).use(Steps);

app.mount("#app");

app.config.globalProperties.$message = message;
