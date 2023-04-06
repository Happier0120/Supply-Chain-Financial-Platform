<template>
  
  <div class="list-content">
    <div class="trace-bar">
      <div class="trace-input-box">
        <div class="query-item-box query-item-box1">
          <span>Ticket number</span>
          <el-input v-model="ticketID"></el-input>
        </div>

        <div class="query-item-box">
          <span>Issue/transfer time</span>
          <el-date-picker :value-format="'yyyyMMdd'" v-model="createTime" type="date"></el-date-picker>
        </div>
      </div>
      <div class="trace-btn">
        <el-button plain icon="el-icon-refresh-left">Reset</el-button>
        <el-button type="primary" icon="el-icon-search" @click="handleChange">Search</el-button>
      </div>
    </div>

    <div class="list-result">
      <div class="res-container">
        <el-table :data="tableData" style="width: 100%">

          <el-table-column fixed="left" prop="ticketID" label="Ticket number" width="150">
          </el-table-column>
          <el-table-column fixed="left" prop="price" label="Ticket amount" width="120">
          </el-table-column>
          <el-table-column fixed="left" prop="ownerOrg" label="Ticket issuer" width="150">
          </el-table-column>
          <el-table-column fixed="left" prop="ownerOrg" label="Ticket receiver" width="150">
          </el-table-column>
          <el-table-column fixed="left" prop="createTime" label="Issue/transfer time" width="160">
          </el-table-column>
          <el-table-column fixed="left" prop="duedate" label="Expired time" width="150">
          </el-table-column>
          <el-table-column fixed="left" prop="description" label="Remark" width="150">
          </el-table-column>

          <el-table-column fixed="left" label="Operation">
            <template slot-scope="scope">
              <el-button @click="detailButton(scope.row)" type="text" size="small">
                Trace
              </el-button>
            </template>
          </el-table-column>

        </el-table>
      </div>
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
          this.tableData = data.data.data;
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
          this.tableData = data.data.data;
        });
    }
  },
  mounted() {
    this.initData()
  },

};
</script>

<style src="../../style/component/ticketList.css"></style>