<template>
  <v-app>
    <Navbar v-show= "this.$Navbar_show" >
    </Navbar>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import Navbar from "@/components/Navbar";
import Navbar2 from "@/components/Navbar"


export default {
  name: "App",
  components: { Navbar },

  data() {
    return {
      Navbar_show:this.$Navbar_show
    };
  },
};


</script>
