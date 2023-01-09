<template>
  <v-container>
    <div  class="page">
      <div>
        <v-breadcrumbs
          :items="items"
          large
        ></v-breadcrumbs>
      </div>
     
      <el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign">
        <el-row :gutter="10">
          <el-col :span="8">
            <el-form-item label="开票编号">
              <el-input v-model="ticket.ticketID"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="8">
            <el-form-item label="票据金额">
              <el-input v-model="ticket.price"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="8">
            <el-form-item label="接收方">
              <el-input v-model="ticket.ownerOrg"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="8">
            <el-form-item label="担保机构">
              <el-input v-model="ticket.guarantor"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="8">
            <el-form-item label="兑付日期">
              <el-input v-model="ticket.duedate"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="10">
          <el-col :span="8">
            <el-form-item label="备注">
              <el-input v-model="ticket.description"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div class="text-center my-2">
          <el-button type="primary" dark large @click="clickCreate">确认开立</el-button>
      </div>
    </div>
  </v-container>
</template>


<script>
import TransferTicketVue from './TransferTicket.vue';
export default {
  data() {
    return {
      items: [
        {
          text: '首页',
          disabled: true,
          href: 'dashboard',
        },
        {
          text: '未到期票据',
          disabled: true,
          href: 'breadcrumbs_link_1',
        },
        {
          text: '开立票据',
          disabled: true,
          href: 'breadcrumbs_link_2',
        },
      ],
      labelPosition: 'right',
      ticket:{
         createTime: "", 
         description: "",
         duedate: "",
         fromOrder: "",
         guarantor: "", 
         objectType: "",
         ownerOrg: "",
         price: "",
         ticketID: ""
      },
    };
  },
  methods:{
    clickCreate() {
      //跳转到详情
      console.log("票据信息",this.ticket);

      this.$axios({
         method: 'post',
         url: '/console/createTicket',
           headers: {
           'Content-Type': 'application/json'//设置Content-Type为jsonss
         },
         data: this.ticket
   }).then(res => {
    console.log("res.data:",res.data);
    this.$message.success('开立成功')  // ....成功信息
    // this.$router.push({ path: `/order/orderList`});
   }).catch(err => {
    console.log("err.data:",err.data);
    // this.$message.error('开立失败')  // .....失败信息
})
  },
},
}
</script>

<style lang="scss" scoped>
.input{
  margin-top: 30px;
}
.input2{
  border: 1px solid;

}
.crate-order{
  margin-top: 50px;
}
.page {
  margin-left: 280px;
}
</style>