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
    ]
})

router.beforeEach((to, from, next) => {
    // if (!to.matched.length) {
    //     next('/404');
    // }
    if (Cookie.get('SESSIONID')) {
        next();
    } else {
       if (to.path === '/login') {
           next();
       }
       next('/login');
    }
})

export default router