import {createRouter, createWebHistory} from "vue-router"
import Cookie from 'js-cookie'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/home",
            component: () => import("@/views/Home.vue"),
        },
        {
            path: "/login",
            component: () => import("@/views/Login.vue"),
        },
        {
            path: "/register",
            component: () => import("@/views/Register.vue"),
        },
        {
            path: "/404",
            component: () => import("@/views/404.vue"),
        },
        {
            path: "/500",
            component: () => import("@/views/500.vue"),
        },
    ]
})

router.beforeEach((to, from, next) => {
    // if (!to.matched.length) {
    //     next('/404');
    // }
    if (Cookie.get('SESSIONID')) {
        next();
    } else {
       if (to.path === '/login' || to.path === '/register') {
           next();
       }
       next('/login');
    }
})

export default router