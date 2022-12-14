
import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import { PromiseDialog } from "vue3-promise-dialog";
import setupInterceptors from "./services/setupInterceptors";

import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "./assets/main.css";

setupInterceptors();



const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(PromiseDialog);

app.mount("#app");
