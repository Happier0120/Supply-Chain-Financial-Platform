<template>
  <v-app>
    <v-container>
      <Navbar v-show= "$store.state.navbar_show" >
      </Navbar>
      <Navbar2 v-show= "$store.state.navbar2_show" >
      </Navbar2>
    </v-container>

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>

  </v-app>
</template>

<script>
import Navbar from "@/components/Navbar";
import Navbar2 from "@/components/Navbar2"
import NavMenu from "./components/NavMenu.vue"

export default {
  name: "App",
  components: { Navbar,Navbar2},

  data() {
    return {
    };
  },
};
</script>


