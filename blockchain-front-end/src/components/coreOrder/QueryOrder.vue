<template>
  <div class="main-box">
    <div class="bottom-box">
      <!-- 票据详情所在界面 -->
      <div class="wea-info">
        <div class="wea-type-group">
          <div class="type-info">
            <span>票据编号: {{ infoData_public.ticketID }} </span>
            <span> 担保机构: {{ infoData_public.guarantor }} </span>
            <span> 票据金额: {{ infoData_privacy.price }} </span>
          </div>
        </div>

        <div class="wea-type-group">
          <div class="type-info">
            <span>票据接收方: {{ infoData_public.ownerOrg }} </span>
            <span> 开单日期: {{ infoData_public.createTime }} </span>
            <span> 自动兑付日期: {{ infoData_public.duedate }} </span>
          </div>
        </div>
        <div class="wea-type-group">
          <div class="type-info">
            <span>备注: {{ infoData_public.description }} </span>
          </div>
        </div>
        <div class="wea-type-group" style="margin: auto ">
          <div id="app" >
            <div v-for="j in this.infoData_blockchain" :key="j" >
            <div >
              <el-card class="box-card">              
                <div slot="header" class="clearfix">

                  <div>{{ "组织MSP：" }}</div>
                  <div class="text">{{  j.record.ownerOrg }}</div>

                  <div>{{ "背书时间：" }}</div>
                  <div class="text">{{  j.timestamp }}</div>

<!-- 
                  <el-button style="float: right; padding: 3px 0" type="text"
                    >操作按钮</el-button
                  > -->
                </div>
                <div class=" item">
                  {{ "通道:"  }} 
                </div>
                <div class="text item">
                  {{j.channelId}}
                </div>
                <div class=" item">
                  {{ "交易ID："  }}
                </div>
                <div class="text item">
                  {{j.txId}}
                </div>
                 <div class=" item">
                  {{ "区块Hash："  }}
                </div>
                <div class="text item">
                  {{j.hash}}
                </div>
                  <div class=" item">
                  {{ "票据编号："  }}
                </div>
                <div class="text item">
                  {{j.record.ticketID}}
                 </div>

              </el-card>

              </div>
              
              <div >
                  <v-img  src="../../assets/right.png" style="width:100px ;height:100px;float: left;margin-left: 30px;margin-top: 180px">
                  </v-img>
                </div>
            </div>
          </div>
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
      infoData_privacy:{},
      infoData_blockchain:{},
    };
  },
  created() {
    console.log("跳转到票据详情");
  },
  methods: {
    initData() {
      // console.log(this.$root.id);
      let ticketID = this.$route.query.id;


      this.$axios
        .get("/console/queryTicketPublicById", { params: {ticketID:ticketID } })
        .then((data) => {
          console.log("infoData_public:", data.data.data);
          this.infoData_public = data.data.data;
        });
    },
    initData2() {
      // console.log(this.$root.id);
      let ticketID = this.$route.query.id;

      this.$axios
        .get("/console/queryTicketPrivateById", { params:{ticketID:ticketID } })
        .then((data) => {
          console.log("infoData_privacy:", data.data.data);
          this.infoData_privacy = data.data.data;
        });
    },
    initData3() {
      // console.log(this.$root.id);
      let ticketID = this.$route.query.id;

      this.$axios
        .get("/console/queryTicketHistoryById", { params:{ticketID:ticketID } })
        .then((data) => {
          console.log("infoData_blockchain:", data.data.data);
          this.infoData_blockchain = data.data.data;
        });
    },
  },
  mounted() {
    this.initData();
    this.initData2();
    this.initData3();

  },
};
</script>

<style lang="scss" scoped>
.main-box {
  padding-left: 50px;
  .top-box {
    //   padding-left: 40px;
    &::after {
      content: "";
      display: block;
      visibility: hidden;
      clear: both;
    }
    .pic {
      float: left;
      width: 120px;
      height: 80px;
      // border: 1px solid #ccc;
      // display: inline-block;
      background-repeat: no-repeat;
      background-size: 100% 100%;
    }
    .info {
      float: left;
      height: 100px;
      margin-left: 80px;
      // display: inline-block;
      .name {
        margin-top: 10px;
        margin-bottom: 30px;
        font-size: 20px;
        font-weight: 900;
      }
      .info-item {
        margin-right: 100px;
      }
    }
  }
  .bottom-box {
    margin-top: 20px;
    .wea-type-group {
      margin-bottom: 50px;
      .type-name {
        width: 50%;
        height: 40px;
        line-height: 40px;
        padding-left: 20px;
        font-size: 18px;
        font-weight: 900;
        background-image: linear-gradient(to right, #eee, #fff);
      }
      .type-info {
        height: 50px;
        margin-top: 40px;
        display: flex;
        padding-left: 20px;
        flex-direction: row;
        span {
          flex: 1;
        }
      }
      .text {
        font-size: 16px;
        font-weight: bold;
      }

      .item {
        word-break: break-all;
        word-wrap: break-word

      }

      .box-card {
        float: left;
        width: 300px;
        margin-top: 40px;
        margin-left: 100px;
        background-color:#dfd2ff	;
      }
    }
  }
}
</style>