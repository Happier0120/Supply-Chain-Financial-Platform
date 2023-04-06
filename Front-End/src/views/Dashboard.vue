<template>
  <div class="board-content">
    <div class="topbar">
      <div class="top-container">
        <div class="left">
          <img src="../images/供应链-copy.png" alt="供应链icon">
          <span class="prodIcon">Supply-Chain Financial Platform</span>
        </div>
        <div class="right">
          <button class="profile-container" @click="noticeClick">
            <div class="avatar-box">
              <img class="avatar" src="../images/avatar14.png" alt="avatar">
            </div>
            <span>{{this.global.orgName}}</span>
          </button>
          <button class="notice-container" @click="noticeClick">
            <img src="../images/notice (1).png" alt="notice">
            <span>Message</span>
          </button>
          <button class="account-op-container" @click="noticeClick">
            <img src="../images/md-log-out (2).png" alt="Log out">
            <span>Log out</span>
          </button>
        </div>
      </div>
    </div>
    <div class="body">
      <div class="side-bar">
        <CoreNavbar v-if="this.global.type == '核心企业'"></CoreNavbar>
        <TransNavbar v-else></TransNavbar>
      </div>
      <div class="info-content">
        <div class="info-container">
          <!-- <div class="top-content">
            
          </div>
          <div class="bottom-content">
          </div> -->
          <router-view/>
        </div>
      </div>
    </div>
  </div>
  
</template>

<script>
import TransNavbar from '@/components/transNavbar.vue';
import CoreNavbar from '@/components/coreNavbar.vue';
import Test from '../components/test.vue'

export default {
  components: {
    "block-index": Test,
    CoreNavbar,
    TransNavbar
},
  data() {
    return {
      infoData: {},
      infoData2: {},
    };
  },
  created() {
    console.log("跳转到票据详情");
  },
  methods: {
    initData() {
      // 返回数据可用
      this.$axios
        .get("/console/queryChaincodeInfo")
        .then((data) => {
          this.infoData = data.data.data;
          console.log("data:",data.data.data)
          this.$forceUpdate()
        });
    },

    initData2() {
      /// 返回数据为空
      this.$axios
        .get("/console/queryTicketsInfo")
        .then((data) => {
          this.infoData2 = data.data.data;
          console.log("data2:",data.data.data)
          this.$forceUpdate()

        });
    },

    noticeClick() {
      this.$notify.info({
          title: 'Notice',
          message: 'This feature is not yet available.'
        });
    }
    
  },
  mounted() {
    this.initData();
    this.initData2();

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
    margin-top: 0px;
    margin-left: 200px;
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
        font-size: 14px;
      }

      .item {
        padding: 18px 0;
      }

      .box-card {
        float: left;
        width: 300px;
        margin-top: 40px;
        margin-left: 100px;
        background-color: #dfd2ff;
        font-size: 25px;
        color: rgb(65, 65, 65);
      }
      .box-card2 {
        float: left;
        width: 300px;
        margin-top: 40px;
        margin-left: 100px;
        background-color: #e3e1e7;
        font-size: 25px;
        color: rgb(65, 65, 65);
      }
    }
  }
}
.block-index {
  padding-top: 700px;
}
</style>


<style src="../style/dashboard.css"></style>