<template>
  <div class="main-box">
    <div class="bottom-box">
      <div class="wea-info">
        <div class="wea-type-group" style="margin: auto">
          <div id="app">
            <el-card class="box-card">
              <div>{{ "区块数" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
              >
                {{this.infoData.blockNumber}}
              </div>

              <div>{{ "BLOCKNumber" }}</div>
              <!-- 
                  <el-button style="float: right; padding: 3px 0" type="text"
                    >操作按钮</el-button
                  > -->
            </el-card>
            <el-card class="box-card">
              <div>{{ "通道数" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
              >
              1
                <!-- {{ this.infoData.blockNumber}} --> 
              </div>

              <div>{{ "CHANNELNUMBER" }}</div>
              <!-- 
                  <el-button style="float: right; padding: 3px 0" type="text"
                    >操作按钮</el-button
                  > -->
            </el-card>
            <el-card class="box-card">
              <div>{{ "节点" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
              >
                {{this.infoData.nodesNumber}}
              </div>

              <div>{{ "NODES" }}</div>
              <!-- 
                  <el-button style="float: right; padding: 3px 0" type="text"
                    >操作按钮</el-button
                  > -->
            </el-card>
            <el-card class="box-card">
              <div>{{ "交易" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
              >
              {{this.infoData.transactionCount}}
              </div>

              <div>{{ "TRANSACTIONS" }}</div>
              <!-- 
                  <el-button style="float: right; padding: 3px 0" type="text"
                    >操作按钮</el-button
                  > -->
            </el-card>
            <el-card class="box-card2">
              <div>{{ "今日开立票据" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
                {{this.infoData2.todayTicketsNumber}}
              </div>
            </el-card>
            <el-card class="box-card2">
              <div>{{ "今日开立金额" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.todayPrice}}
              </div>
            </el-card>
            <el-card class="box-card2">
              <div>{{ "总开立票据" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.todayTicketsNumber}}
              </div>
            </el-card>
            <el-card class="box-card2">
              <div>{{ "总开立金额" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.price}}
              </div>
            </el-card>
            <el-card class="box-card2">
              <div>{{ "今日到期票据" }}</div>
              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.todayExpiredTicketsNumber}}
              </div>


            </el-card>
            <el-card class="box-card2">
              <div>{{ "今日兑付金额" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.todayPrice}}
              </div>

            </el-card>
            <el-card class="box-card2">
              <div>{{ "总到期票据" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.expiredTicketsNumber}}
              </div>


            </el-card>
            <el-card class="box-card2">
              <div>{{ "总兑付金额" }}</div>

              <div
                style="
                  font-size: 40px;
                  color: black;
                  margin-top: 5px;
                  margin-bottom: 5px;
                "
                v-if="this.infoData2"
              >
              {{this.infoData2.price}}
              </div>


            </el-card>
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
</style>