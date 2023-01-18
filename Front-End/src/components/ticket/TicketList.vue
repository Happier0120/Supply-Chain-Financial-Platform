<template>
  <div class="page">
    <div class="query-box">
      <div class="query-item">
        票据编号:
        <el-input class="query-input" v-model="ticketID" placeholder="请输入内容"></el-input>
      </div>
      <div class="query-item">
        开立/转让日期:
        <el-date-picker :value-format="'yyyyMMdd'" class="query-input" v-model="createTime" type="date"
          placeholder="选择日期">
        </el-date-picker>
      </div>
      <!-- <div class="query-item">
        票据编号:
        <el-input class="query-input" v-model="prodNumber" placeholder="请输入内容" @input="handleChange"></el-input>
      </div> -->
      <el-button type="primary" @click="handleChange" style="  margin-left: 20px; ">查 询</el-button>
      <!-- <div class="crate-order">
        <el-button type="primary" @click="clickCreate">开立票据</el-button>
      </div> -->
    </div>


    <div style="margin-top:20px">
      <el-table :data="tableData" style="width: 100%">
        <!-- <el-table-column
          fixed="left"
          prop="id"
          label="序号"
          width="100"
          align="center"
        >
        </el-table-column> -->
        <el-table-column fixed="left" prop="ticketID" label="票据编号" width="150">
        </el-table-column>
        <el-table-column fixed="left" prop="price" label="票据金额" width="100">
        </el-table-column>
        <el-table-column fixed="left" prop="ownerOrg" label="票据发起方" width="150">
        </el-table-column>
        <el-table-column fixed="left" prop="ownerOrg" label="票据接收方" width="150">
        </el-table-column>
        <el-table-column fixed="left" prop="createTime" label="开立/转让日期" width="150">
        </el-table-column>
        <el-table-column fixed="left" prop="duedate" label="兑付日期" width="150">
        </el-table-column>
        <el-table-column fixed="left" prop="description" label="备注" width="50">
        </el-table-column>

        <el-table-column fixed="left" label="操作">
          <template slot-scope="scope">
            <el-button @click="detailButton(scope.row)" type="text" size="small">
              溯源信息
            </el-button>
            <el-button @click="transferButton(scope.row)" type="text" size="small" v-if="transfer">
              转让票据
            </el-button>
          </template>
        </el-table-column>

      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      listId: "",
      tableData: [],
      transfer: false,
      createTime: "",
      duedate: "",
      ticketID: "",
      type: "",
    };
  },
  watch: {
    "$route.params.id": {  //监控$route.params.id
      handler() {
        this.initData()
      },
      deep: true
    }
  },
  methods: {
    handleChange() {
      this.$axios
        .get("/console/queryTickets", { params: { createTime: this.createTime, duedate: this.duedate, ticketID: this.ticketID, type: this.type } })
        .then((data) => {
          console.log("infoData_public:", data.data.data);
          this.tableData = data.data.data
        });
      console.log("createTime:", this.createTime)
    },
    detailButton(row) {
      //跳转到详情
      console.log("row:", row);
      this.$router.push({ path: `/ticket/ticketdetail`, query: { id: row.ticketID } });
    },
    transferButton(row) {
      //跳转到转让
      console.log("row:", row);
      this.$router.push({ path: `/ticket/transferticket`, query: { id: row.ticketID } });

    },
    initData() {
      this.listId = this.$route.params.id
      if (this.listId == "notExpired") {
        this.type = "notExpired"
        console.log("notExpired")
      } else if (this.listId == "expired") {
        this.type = "expired"
        console.log("expired")
      } else if (this.listId == "notTransferred") {
        this.type = "notTransferred"
        console.log("notTransferred")
        this.transfer == true
      } else if (this.listId == "transferred") {
        this.type = "transferred"
        console.log("transferred")
      }
      this.$axios
        .get("/console/queryTickets", { params: { createTime: this.createTime, duedate: this.duedate, ticketID: this.ticketID, type: this.type } })
        .then((data) => {
          console.log("infoData_public:", data.data.data);
          this.tableData = data.data.data
        });
    }
  },
  mounted() {
    this.initData()
  },

};
</script>

<style lang="scss">
.query-box {
  float: left;
  margin-top: 20px;
}

.query-item {
  float: left;
  width: 400;
  margin-left: 20px;

  .query-input {
    width: 250px;
  }
}

.el-input {
  display: inline-block;
  width: 350px;
  margin-left: 10px;
}

.crate-order {
  margin-left: 20px;
  margin-top: 20px;
  margin-bottom: 20px;
}

.page {
  margin-left: 280px;
}
</style>