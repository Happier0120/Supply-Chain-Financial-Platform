import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
import Order from '../views/Order.vue'
// import Navbar from '../components/Navbar.vue'
// import Navbar2 from '../components/Navbar2.vue'

import Home from '../components/home.vue'
import TicketList from '../components/ticket/TicketList.vue'
import IssueTicket from '../components/ticket/IssueTicket.vue'
import TransferTicket from "../components/ticket/TransferTicket.vue"
import TraceTicket from "../components/ticket/TraceTicket.vue"
import TicketDetail from "../components/ticket/TicketDetail.vue"

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
    children: [
      {
        path: '/home',
        name: 'home',
        component: Home,
      },
      {
        path: '/ticket/traceticket',
        name: 'traceticket',
        component: TraceTicket,
      },
      {
        path: '/ticket/ticketlist/:id',
        name: 'ticketlist',
        component: TicketList,
      },
      {
        path: '/ticket/issueticket',
        name: 'issueticket',
        component: IssueTicket,
      },
      {
        path: '/ticket/ticketdetail',
        name: 'ticketdetail',
        component: TicketDetail
      }
    ]
  },
  // {
  //   path: '/Navbar',
  //   name: 'Navbar',
  //   component: Navbar,
  // },
  // {
  //   path: '/ticket',
  //   name: 'ticket',
  //   component: Order,
  //   children: [
  //     {
  //       path: '/ticket/ticketdetail',
  //       name: 'ticketdetail',
  //       component: TicketDetail,
  //     },
  //     {
  //       path: '/ticket/issueticket',
  //       name: 'issueticket',
  //       component: IssueTicket,
  //     },
  //     // {
  //     //   path: '/ticket/ticketlist',
  //     //   name: 'ticketlist',
  //     //   component: TicketList,
  //     // },
  //     {
  //       path: '/ticket/ticketlist/:id',  // 这个用于传参
  //       name: 'ticketlist',
  //       component: TicketList,
  //     },
  //     {
  //       path: '/ticket/transferticket',
  //       name: 'transferticket',
  //       component: TransferTicket,
  //     },
  //     {
  //       path: '/ticket/traceticket',
  //       name: 'traceticket',
  //       component: TraceTicket,
  //     },

  //   ]
  // },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
