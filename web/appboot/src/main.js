import { createApp } from "vue";
import App from "./App.vue";
import {
  Button,
  Input,
  Form,
  Radio,
  InputNumber,
  message,
  Select,
  Tooltip,
  Switch,
  List,
} from "ant-design-vue";

import "ant-design-vue/es/message/style/css"; //vite只能用 ant-design-vue/es 而非 ant-design-vue/lib

const app = createApp(App);

app.use(Button);
app.use(Input);
app.use(Form);
app.use(Radio);
app.use(InputNumber);
app.use(Select);
app.use(Tooltip);
app.use(Switch);
app.use(List);

app.mount("#app");

app.config.globalProperties.$message = message;
