import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Order from '../views/Order.vue'
import Bill from '../views/Bill.vue'
import Trans from '../views/Trans.vue'

import OrderList from '../components/coreOrder/OrderList.vue'
import CreateOrder from '../components/coreOrder/CreateOrder.vue'
import EditOrder from '../components/coreOrder/EditOrder.vue'
import QueryOrder from '../components/coreOrder/QueryOrder.vue'


import CreateBill from '../components/coreBill/CreateBill.vue'
import EditBill from '../components/coreBill/EditBill.vue'
import QueryBill from '../components/coreBill/QueryBill.vue'

import CreateTrans from '../components/coreTrans/CreateTrans.vue'
import EditTrans from '../components/coreTrans/EditTrans.vue'
import QueryTrans from '../components/coreTrans/QueryTrans.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'dashboard',
    component: Dashboard
  },

  {
    path: '/order',
    name: 'order',
    component: Order,
    children: [
      {
        path: '/order/orderList',
        name: 'orderlist',
        component: OrderList
      },{
      path: '/order/createOrder',
      name: 'createorder',
      component: CreateOrder
    },
    {
      path: '/order/editOrder',
      name: 'editorder',
      component: EditOrder
    },
    {
      path: '/order/queryOrder',
      name: 'queryorder',
      component: QueryOrder
    }]
  },
  {
    path: '/bill',
    name: 'bill',
    component: Bill,
    children: [{
      path: '/bill/createBill',
      name: 'createBill',
      component: CreateBill
    },
    {
      path: '/bill/editBill',
      name: 'editBill',
      component: EditBill
    },
    {
      path: '/bill/queryBill',
      name: 'queryBill',
      component: QueryBill
    }]
  },
  {
    path: '/trans',
    name: 'trans',
    component: Trans,
    children: [{
      path: '/trans/createTrans',
      name: 'createTrans',
      component: CreateTrans
    },
    {
      path: '/trans/editTrans',
      name: 'editTrans',
      component: EditTrans
    },
    {
      path: '/trans/queryTrans',
      name: 'queryTrans',
      component: QueryTrans
    }]
  },
 
  
  
  

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
