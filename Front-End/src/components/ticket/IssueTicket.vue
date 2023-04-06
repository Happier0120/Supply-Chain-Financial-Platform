<template>
  <div class="issue-content">
    <div class="crumb-bar">
      <v-breadcrumbs :items="items" small></v-breadcrumbs>
    </div>

    <div class="issue-workbench">
      <div class="workbench-container">
        <div class="issue-bill-header-area">
          <span class="op-note">Issue ticket</span>
          <el-button type="primary" @click="clickCreate">Issue</el-button>
        </div>
        <el-form :label-position="labelPosition" label-width="100px">
          <el-row :gutter="10">
            <el-col :span="8">
              <el-form-item label="Ticket issuer">
                <el-input v-model="ticket.fromOrder" :disabled="true"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="8">
              <el-form-item label="Ticket amount">
                <el-input v-model="ticket.price"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="8">
              <el-form-item label="Ticket receiver">
                <el-input v-model="ticket.ownerOrg"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="8">
              <el-form-item label="Expired date">
                <el-date-picker :value-format="'yyyyMMdd'" class="query-input" v-model="ticket.duedate" type="date"
                  placeholder="select date">
                </el-date-picker>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="8">
              <el-form-item label="Remark">
                <el-input v-model="ticket.description"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
    </div>

  </div>
</template>


<script>
import TransferTicketVue from './TransferTicket.vue';
export default {
  data() {
    return {
      items: [
        {
          text: 'Unexpired tickets',
          disabled: true,
          href: 'breadcrumbs_link_1',
        },
        {
          text: 'Issue ticket',
          disabled: true,
          href: 'breadcrumbs_link_2',
        },
      ],
      labelPosition: 'right',
      ticket: {
        createTime: "",
        description: "",
        duedate: "",
        fromOrder: this.global.orgName,
        guarantor: "",
        objectType: "",
        ownerOrg: "",
        price: "",
        ticketID: ""
      },
    };
  },
  methods: {
    clickCreate() {
      //跳转到详情
      console.log("票据信息", this.ticket);

      this.$axios({
        method: 'post',
        url: '/console/createTicket',
        headers: {
          'Content-Type': 'application/json'//设置Content-Type为jsonss
        },
        data: this.ticket
      }).then(res => {
        console.log("res.data:", res.data);
        this.$message.success('开立成功')  // ....成功信息
        this.$router.push({ path: `/ticket/ticketList/notExpired` });
      }).catch(err => {
        console.log("err.data:", err.data);
        this.$message.error('开立失败')  // .....失败信息
      })
    },
  },
}
</script>

<style src="../../style/component/issueTicket.css"></style>