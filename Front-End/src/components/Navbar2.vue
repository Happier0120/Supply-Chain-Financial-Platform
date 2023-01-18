<template>
  <nav>
    <v-app-bar flat app dark color="deep-purple">
      <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

      <v-toolbar-title>中交智运“区块链+网络货运”平台</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" absolute permanent>
      <v-list>
        <v-list-item>
          <v-list-item-avatar class="mx-auto">
            <v-img aspect-ration="0.3" src="../assets/ZJZY-logo.png"></v-img>
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
    items: [
      {
        action: "el-icon-s-home",
        route: "/dashboard",
        items: [
          { subtitle: "运输企业首页", route: "/dashboard" },
          { subtitle: "票据溯源", route: "/ticket/traceticket" },
        ],
        title: "首页",
      },
      {
        action: "el-icon-s-ticket",
        route: "/ticket",
        items: [
          { subtitle: "票据列表", route: "/ticket/ticketList/notTransferred" },  //在index.js路由配置配置了可传参，这边传参控制list按钮的展示

        ],
        title: "未转让票据",
      },
      {
        action: "el-icon-s-ticket",
        route: "/ticket",
        items: [
          { subtitle: "票据列表", route: "/ticket/ticketList/transferred" },
        ],
        title: "已转让票据",
      },
      {
        action: "el-icon-s-ticket",
        route: "/ticket",
        items: [
          { subtitle: "票据列表", route: "/ticket/ticketList/expired" },
          // { subtitle: "", route: "/trans/editTrans" },
          // { subtitle: "查询运单", route: "/trans/queryTrans" },
        ],
        title: "已兑付票据",
      },
    ],
  }),
};
</script>
