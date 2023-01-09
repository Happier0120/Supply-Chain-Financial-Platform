import Vue from 'vue'
import store from './store'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
//element-ui
import Element from 'element-ui'
import "element-ui/lib/theme-chalk/index.css"
//axios
import axios from 'axios'
Vue.prototype.$axios = axios
//baseURL
axios.defaults.baseURL = "http://192.168.1.61:14449"
Vue.use(Element)

Vue.config.productionTip = false

Vue.prototype.$Navbar_show = false  //控制侧边栏的展示与否，这边设置全局变量，然后APP.vue界面进行控制
Vue.prototype.$Navbar2_show = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App),
  data() {
    return {
      id: '',
    }
  }
}).$mount('#app')

