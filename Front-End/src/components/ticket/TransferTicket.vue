<template>
  <v-container>
    <h1>转让票据</h1>
    <el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign">
      <el-row :gutter="10">
        <el-col :span="8">
          <el-form-item label="票据编号">
            <el-input v-model="ticketID"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row :gutter="10">
        <el-col :span="8">
          <el-form-item label="接收组织">
            <el-input v-model="toOrgMSPID"></el-input>
          </el-form-item>
        </el-col>  
      </el-row>
    </el-form>
    <v-container fluid class="pa-0">
      <div class="text-center my-2">
        <v-btn color="green" dark large @click="clickTransfer"> 确认转让 </v-btn>
      </div>
    </v-container>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      ticketID: "",
      toOrgMSPID: "", 
      items: [
        "待转让票据ID",
        "接收组织MSP",
      ],
    };
  },
  methods:{
    clickTransfer(){
       console.log("转让票据");
       
       this.$axios({
        methoad:'post',
        url:'/console/transferTicket',  ///开放接口
        headers: {
           'Content-Type': 'application/json'  //设置Content-Type为jsonss
        },
        data:{ticketID, toOrgMSPID},
       }).then(res => {
          console.log("res.data:",res.data);
          this.$message.success('转让成功')  // ....成功信息
          this.$router.push({ path: `/order/orderList`});
        }).catch(err => {
            console.log("err.data:",err.data);
            this.$message.error('转让失败')  // .....失败信息
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