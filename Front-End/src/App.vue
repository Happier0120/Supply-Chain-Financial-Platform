<template>
  <v-app>
    <!-- <Navbar v-show= "$store.state.navbar_show">
    </Navbar>
    <Navbar2 v-show= "$store.state.navbar2_show">
    </Navbar2> -->
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>


<style>
.v-main {
  padding: 0px 0px 0px 0px !important;
}
</style>

<script>
// import Navbar from "@/components/Navbar";
// import Navbar2 from "@/components/Navbar2"

export default {
  name: "App",
  // components: { Navbar, Navbar2 },

  data() {
    return {
    };
  },
  created() {
    // if(sessionStorage.getItem('navbar_show')==null && sessionStorage.getItem('navbar2_show')==null){
    //   console.log("跳转到票据详情");
    // }else{
    //   this.$store.commit("changeNavbar",sessionStorage.getItem('navbar_show'));
    //   this.$store.commit("changeNavbar2",sessionStorage.getItem('navbar2_show'));
    //   console.log(" Navbar_show:", this.$store.state.navbar_show);
    //   console.log(" Navbar2_show:",this.$store.state.navbar2_show);
    // }
  },
};
</script>


