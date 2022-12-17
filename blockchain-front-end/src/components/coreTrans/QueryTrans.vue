<template>
  <div class="main-box">
    <!-- <div class="top-box">
      <div class="pic"></div>
      <div class="info">
        <div class="name">
          {{ item.name }}
          <el-tag
            style="margin-left: 100px"
            :type="item.statu ? 'success' : 'info'"
            effect="dark"
          >
            {{ item.statu ? "在 线" : "离 线" }}
          </el-tag>
        </div>
        <div>
          <span class="info-item">编号: {{ item.wea_id }}</span>
          <span class="info-item">物联网卡号: {{ item.sensor_card }}</span>
          <span class="info-item">所属地块: {{ item.greenhouse }}</span>
        </div>
      </div>
    </div> -->
    <div class="bottom-box">
      <!-- <el-empty
        v-if="Object.keys(infoData).length == 0"
        description="暂无数据"
      ></el-empty> -->
      <!-- <div v-else class="wea-info"></div> -->
      <div class="wea-info">
        <div class="wea-type-group">
          <div class="type-info">
            <span>票据编号: {{ infoData.billCode }} </span>
            <span> 担保机构: {{ infoData.billCode }} </span>
            <span> 票据金额: {{ infoData.billCode }} </span>
          </div>
        </div>

        <div class="wea-type-group">
          <div class="type-info">
            <span>票据接收方: {{ infoData.billCode }} </span>
            <span> 开单日期: {{ infoData.billCode }} </span>
            <span> 自动兑付日期: {{ infoData.billCode }} </span>
          </div>
        </div>
        <div class="wea-type-group">
          <div class="type-info">
            <span>备注: {{ infoData.billCode }} </span>
          </div>
        </div>
        <div class="wea-type-group" style="margin: auto">
          <div id="app">
            <div v-for="j in 6" :key="j">
              <el-card class="box-card">
                <div slot="header" class="clearfix">

                    <div>{{ "组织MSP：" }}</div>
                  <div>{{  j }}</div>

                  <div>{{ "背书时间：" }}</div>
                  <div>{{  j }}</div>
                </div>
                 <div class="text item">
                  {{ "通道：xxxxxxx"  }}
                </div>
                <div class="text item">
                  {{ "交易ID：xxxxxxx"  }}
                </div>
                 <div class="text item">
                  {{ "区块Hash：xxxxxxx"  }}
                </div>
                  <div class="text item">
                  {{ "票据编号：xxxxxxx"  }}
                </div>

              </el-card>
              
              <div >
                  <v-img  src="../../assets/left.png" style="width:100px ;height:100px;float: left;margin-left: 30px;margin-top: 180px">
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
      infoData: {},
    };
  },
  created() {
    console.log("跳转到票据详情");
  },
  methods: {
    initData() {
      // console.log(this.$root.id);
      let id = this.$route.query.id;
      this.$axios
        .get("/console/queryTmsOrderById", { params: { id: id } })
        .then((data) => {
          this.infoData = data.data.data;
        });
    },
  },
  mounted() {
    this.initData();
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
        background-color:#dfd2ff	;
      }
    }
  }
}
</style>