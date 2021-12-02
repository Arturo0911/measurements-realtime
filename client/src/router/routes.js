//import Realtime from '@/components/Realtime.vue'

import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)


const routes = [
    {
        path: '/',
        component:()=> import(/* webpackChunkName: "about" */ '../components/RealTime.vue'),
        name:"search"
    }
]

export default routes