<template>
  <v-container>
    <div>
      <h1>票据列表->开立票据</h1>
    <div class="input">担保机构： <input  class="input2" v-model="ticket.guarantor" placeholder="请输入" /></div> 
    <div class="input">票据金额： <input  class="input2" v-model="ticket.price" placeholder="请输入数字" /></div> 
    <div class="input">票据接收方： <input  class="input2" v-model="ticket.ownerOrg" placeholder="请输入" /></div> 
    <div class="input">兑付日期： <input  class="input2" v-model="ticket.duedate" placeholder="请输入" /></div> 
    <div class="input">备注： <input  class="input2" v-model="ticket.description" placeholder="请输入" /></div> 

    </div>
    <div class="crate-order">
        <el-button type="primary" @click="clickCreate">确认开立</el-button>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
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
    this.$router.push({ path: `/order/orderList`});

   }).catch(err => {
    console.log("err.data:",err.data);
    this.$message.error('开立失败')  // .....失败信息
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

</style>