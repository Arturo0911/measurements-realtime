import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)


const routes = [
    {
        path: '/',
        component:()=> import(/* webpackChunkName: "about" */ '../components/RealTime.vue'),
        name:"realtime"
    },
    {
        path: '/statistics',
        component:()=> import(/* webpackChunkName: "about" */ '../components/Statistic.vue'),
        name:"statistic"
    },
    {
        path: '/charts',
        component:()=> import(/* webpackChunkName: "about" */ '../components/DataChart.vue'),
        name:"charts"
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
  })
  


export default router