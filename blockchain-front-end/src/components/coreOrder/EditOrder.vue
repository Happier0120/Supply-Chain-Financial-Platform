<template>
  <v-container>
    <div>
      <h1>订单->修改订单</h1>
      <v-form>
        <v-container>
          <v-row>
            <v-col cols="12" sm="6" v-for="item in items" :key="item">
              <v-text-field :label="item" outlined></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-form>
    </div>
    <v-container fluid class="pa-0">
      <div class="text-center my-2">
        <v-btn color="green" dark large> 确认修改 </v-btn>
      </div>
    </v-container>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      items: [
        "委托项目",
        "委托单价",
        "客户机构",
        "承运机构",
        "发货单位",
        "收货单位",
        "发货地",
        "卸货地",
        "预计发货时间",
        "最晚到达时间",
        "完成时间",
        "运输距离",
        "发订单状态",
        "订单类型",
      ],
    };
  },
};
</script>
