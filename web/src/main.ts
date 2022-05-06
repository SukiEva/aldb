import {createApp} from "vue";
import App from "./App.vue";
import router from "./router";

import "~/styles/index.scss";
import "element-plus/dist/index.css";
import "element-plus/theme-chalk/src/message.scss";


const app = createApp(App);
// app.use(ElementPlus);

app.use(router);
app.mount("#app");
