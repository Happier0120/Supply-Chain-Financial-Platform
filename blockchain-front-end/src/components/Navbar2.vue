<template>
  <nav>
    <v-app-bar flat app dark color="deep-purple">
      <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

      <v-toolbar-title>中交智运“区块链+网络货运”平台</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" absolute temporary>
      <v-list>
        <v-list-item>
          <v-list-item-avatar class="mx-auto">
            <v-img aspect-ration="0.3" src="../assets/ZJZY-logo.png"></v-img>
          </v-list-item-avatar>
        </v-list-item>

        <v-list-item link>
          <v-list-item-content>
            <v-list-item-title class="text-h6"> CCCCIT </v-list-item-title>
            <v-list-item-subtitle>ccccit@admin.com</v-list-item-subtitle>
          </v-list-item-content>

          <v-list-item-action>
            <v-icon>mdi-menu-down</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-list>
      <v-list nav dense>
        <v-list-group
          v-for="item in items"
          :key="item.title"
          :prepend-icon="item.action"
          no-action
        >
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title v-text="item.title"></v-list-item-title>
            </v-list-item-content>
          </template>

          <v-list-item
            v-for="child in item.items"
            :key="child.subtitle"
            :to="child.route"
            router
          >
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
        action: "mdi-silverware-fork-knife",
        route: "/bill",
        items: [
          { subtitle: "运输企业首页", route: "/bill/editBill" },
        ],
        title: "首页",
      },
      {
        action: "mdi-silverware-fork-knife",
        route: "/order",
        items: [
          { subtitle: "票据列表", route: "/order/OrderList" },
          { subtitle: "开立票据", route: "/order/createOrder" },
          // { subtitle: "编辑订单", route: "/order/editOrder" },
          { subtitle: "查询票据", route: "/order/queryOrder" },
        ],
        title: "未到期票据",
      },

      {
        action: "mdi-silverware-fork-knife",
        route: "/trans",
        items: [
          { subtitle: "票据列表", route: "/trans/createTrans" },
          // { subtitle: "", route: "/trans/editTrans" },
          // { subtitle: "查询运单", route: "/trans/queryTrans" },
        ],
        title: "已到期票据",
      },
    ],
  }),
};
</script>
