import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
import Order from '../views/Order.vue'

import TicketList from '../components/ticket/TicketList.vue'
import IssueTicket from '../components/ticket/IssueTicket.vue'
import TransferTicket from "../components/ticket/TransferTicket.vue"
import TraceTicket from "../components/ticket/TraceTicket.vue"

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login,
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Dashboard,
  },
  {
    path: '/ticket',
    name: 'ticket',
    component: Order,
    children: [
      {
        path: '/ticket/issueticket',
        name: 'issueticket',
        component: IssueTicket,
      },
      {
        path: '/ticket/ticketlist',
        name: 'ticketlist',
        component: TicketList,
      },
      {
        path: '/ticket/transferticket',
        name: 'transferticket',
        component: TransferTicket,
      },
      {
        path: '/ticket/traceticket',
        name: 'traceticket',
        component: TraceTicket,
      },]
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
