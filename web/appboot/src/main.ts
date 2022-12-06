import { message, Button, Input, Form, Radio, InputNumber, Select, Tooltip, Switch, List, Steps } from "ant-design-vue";
// https://www.antdv.com/docs/vue/getting-started-cn
import 'ant-design-vue/es/message/style/css';
import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);

app.use(Button).use(Input).use(Form).use(Radio).use(InputNumber).use(Select).use(Tooltip).use(Switch).use(List).use(Steps);
app.mount('#app');

app.config.globalProperties.$message = message;
