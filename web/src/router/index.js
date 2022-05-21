import {createRouter, createWebHistory} from "vue-router"
import Home from "@/views/Home.vue";
import User from "@/views/User.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            component: Home,
        },
        {
            path: "/user",
            component: User,
            redirect: "/user/login",
            children: [
                {
                    path: "login",
                    component: () => import("@/components/Login.vue"),
                },
                {
                    path: "register",
                    component: () => import("@/components/Register.vue"),
                }
            ]
        }
    ]
})

export default router