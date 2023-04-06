<template>
  <div class="trace-content">
    <div class="trace-bar">
      <div class="trace-input-box">
        <span>Ticket number</span>
        <el-input v-model="ticketID"></el-input>
      </div>
      <div class="trace-btn">
        <el-button plain icon="el-icon-refresh-left">Reset</el-button>
        <el-button type="primary" icon="el-icon-search" @click="handleChange">Search</el-button>
      </div>
    </div>
    <div class="trace-result">
      <el-empty :image-size="200" description="No data available" v-if="!this.hasRes"></el-empty>
      <div class="res-container" v-else>
        <el-descriptions class="margin-top" :column="2" border>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-s-flag"></i>
              Ticket number
            </template>
            {{ticketID}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-coin"></i>
              Ticket amount
            </template>
            {{ infoData_privacy.price }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-user"></i>
              Ticket issuer
            </template>
            {{ infoData_public.fromOrder }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-user-solid"></i>
              Ticket receiver
            </template>
            {{ infoData_public.ownerOrg }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-date"></i>
              Issue/transfer date
            </template>
            {{infoData_public.createTime}}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-shopping-cart-1"></i>
              Expired date
            </template>
            {{ infoData_public.duedate }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-notebook-1"></i>
              Remark
            </template>
            {{ infoData_public.description }}
          </el-descriptions-item>
        </el-descriptions>

        <div class="history-box">
          <el-steps align-center>
            <el-step v-for="(item, index) in this.infoData_blockchain">
              <template slot="description">
                <div class="bill-info-box">
                  <el-card>
                    <div class="item-info">
                      <span class="item-info-title">Ticket number:</span>
                      <span>{{ item.record.ticketID }}</span>
                    </div>
                    <div class="item-info">
                      <span class="item-info-title">Organization MSP:</span>
                      <span>{{ item.record.ownerOrg }}</span>
                    </div>
                    <div class="item-info">
                      <span class="item-info-title">Endorsement time:</span>
                      <span>{{item.timestamp}}</span>
                    </div>
                    <div class="item-info">
                      <span class="item-info-title">Channel ID:</span>
                      <span>{{ item.channelId }}</span>
                    </div>
                    <div class="item-info">
                      <span class="item-info-title">Transaction ID:</span>
                      <span>{{ item.txId }}</span>
                    </div>
                    <div class="item-info">
                      <span class="item-info-title">Block hash:</span>
                      <span>{{ item.hash }}</span>
                    </div>
                  </el-card>
                </div>
              </template>
            </el-step>
          </el-steps>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      infoData_public: {},
      infoData_privacy: {},
      infoData_blockchain: {},
      ticketID: "",
      hasRes: false
    };
  },
  created() {
    console.log("跳转到票据详情");
    console.log(this.infoData_public);
  },
  methods: {
    initData() {
      // console.log(this.$root.id);

      this.$axios
        .get("/console/queryTicketPublicById", { params: { ticketID: this.ticketID } })
        .then((data) => {
          console.log("infoData_public:", data.data.data);
          this.infoData_public = data.data.data;
          // this.hasRes = true
        });
    },
    initData2() {
      // console.log(this.$root.id);
      let ticketID = this.$route.query.id;

      this.$axios
        .get("/console/queryTicketPrivateById", { params: { ticketID: this.ticketID } })
        .then((data) => {
          console.log("infoData_privacy:", data.data.data);
          this.infoData_privacy = data.data.data;
        });
    },
    initData3() {
      // console.log(this.$root.id);
      let ticketID = this.$route.query.id;

      this.$axios
        .get("/console/queryTicketHistoryById", { params: { ticketID: this.ticketID } })
        .then((data) => {
          console.log("infoData_blockchain:", data.data.data);
          if (data.data.data !== null) {
            this.infoData_blockchain = data.data.data;
            this.hasRes = true;
          }
        });
    },
    handleChange() {
      this.initData();
      this.initData2();
      this.initData3();

    },
  },

};
</script>

<style src="../../style/component/traceTicket.css"></style>