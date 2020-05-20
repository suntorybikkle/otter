import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Record from '../views/Record.vue'
import Table from '../graphs/Table.vue'
import Liner from '../graphs/Liner.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/record',
    name: 'Record',
    component: Record,
    children: [
      {
        path: 'liner',
        name: 'Liner',
        component: Liner
      },
      {
        path: 'table',
        name: 'Table',
        component: Table
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
