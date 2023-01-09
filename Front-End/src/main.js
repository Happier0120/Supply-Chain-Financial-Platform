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

