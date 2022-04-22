import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../views/Home.vue";
import NotFound from "../views/404.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export default router;
