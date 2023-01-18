import Vue from "vue"
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        navbar_show: false,
        navbar2_show: false,

    },
    mutations: {
        changeNavbar(state, bool) {
            state.navbar_show = bool
        },
        changeNavbar2(state, bool) {
            state.navbar2_show = bool
        },

    }
})