import Vue from "vue"
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        navbar_show: false,
        navbar2_show: false,
    },
    mutations: {
        changeNavbar(state) {
            state.navbar_show = !state.navbar_show
        },
        changeNavbar2(state) {
            state.navbar2_show = !state.navbar_show
        }
    }
})