<template>
  <div class="main-box">
    <div class="bottom-box">
      <div class="wea-info">
        <el-row>
          <el-col >
              <div class="query-box">
               编号
                <el-input
                  class="query-input"
                  v-model="prodNumber"
                  placeholder="请输入待溯源票据编号"
                  @input="handleChange"
                ></el-input>
              </div> 
              <el-button type="primary" @click="handleChange"> 开始溯源</el-button>
          </el-col>
         
        </el-row>
        

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
.query-box {
  float: left;
  margin-top: 0px;
}
.query-item {
  float: left;
  width: 200px;
  margin-left: 0px;
  .query-input {
    width: 100px;
  }
}
.el-input {
  display: inline-block;
  width: 200px;
  margin-left: 20px;
}
.main-box {
  padding-left: 50px;
  margin-left: 280px;
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