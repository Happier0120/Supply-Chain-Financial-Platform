import global from '../global'


<template>
  <nav>
    <v-app-bar flat app dark color="deep-purple">
      <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

      <v-toolbar-title>中交智运基于电子票据的SCF供应链金融平台</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" absolute permanent>
      <v-list>
        <v-list-item>
          <v-list-item-avatar class="mx-auto">
            <v-img aspect-ration="0.6" src="../assets/ZJZY-logo.png">
            </v-img>
          </v-list-item-avatar>
        </v-list-item>

        <v-list-item link>
          <v-list-item-content>
            <v-list-item-title class="text-h6"> CTCCCC </v-list-item-title>
            <v-list-item-subtitle>中交智运</v-list-item-subtitle>
          </v-list-item-content>

          <v-list-item-action>
            <v-icon>mdi-menu-down</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-list>

      <v-list nav dense>
        <v-list-group v-for="item in items" :key="item.title" :prepend-icon="item.action" no-action>
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title v-text="item.title"></v-list-item-title>
            </v-list-item-content>
          </template>

          <v-list-item v-for="child in item.items" :key="child.subtitle" :to="child.route" router>
            <v-list-item-content>
              <v-list-item-title v-text="child.subtitle"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-group>
      </v-list>
    </v-navigation-drawer>
  </nav>
</template>

<script>
export default {
  data: () => ({
    drawer: false,
    account: global.orgName,
    type: "",
    items: [
      {
        action: "el-icon-s-home",
        route: "/ticket",
        items: [
          { subtitle: "核心企业首页", route: "/dashboard" },
          { subtitle: "票据溯源", route: "/ticket/traceticket" },
        ],
        title: "首页",
      },
      {
        action: "el-icon-s-ticket",
        route: "/ticket",
        items: [
          { subtitle: "票据列表", route: "/ticket/ticketList/notExpired" },
          { subtitle: "开立票据", route: "/ticket/issueticket" },
          // { subtitle: "票据溯源", route: "/ticket/traceticket" },
        ],
        title: "未兑付票据",
      },

      {
        action: "el-icon-s-ticket",
        route: "/ticket",
        items: [
          { subtitle: "票据列表", route: "/ticket/ticketList/expired" },
        ],
        title: "已兑付票据",
      },
    ],
  }),
  // mounted() {
  //   let account = this.$route.query.account;
  //   console.log("account:", account)
  // }
};
</script>

<style lang="scss" scoped>
.mx-auto {
  margin-top: 80px;
}
</style>