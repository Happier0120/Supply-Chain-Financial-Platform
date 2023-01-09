  <template>
  <div>
    <div class="query-box">
      <div class="query-item">
        收单企业:
        <el-input
          class="query-input"
          v-model="prodName"
          placeholder="请输入内容"
          @input="handleChange"
        ></el-input>
      </div>
      <div class="query-item">
        开单日期:
        <el-date-picker
          class="query-input"
          v-model="createTime"
          type="date"
          placeholder="选择日期"
          @change="handleChange"
        >
        </el-date-picker>
      </div>
      <div class="query-item">
        票据编号:
        <el-input
          class="query-input"
          v-model="prodNumber"
          placeholder="请输入内容"
          @input="handleChange"
        ></el-input>
      </div>
      <el-button type="primary" @click="handleChange">查 询</el-button>
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
        <el-table-column
          fixed="left"
          prop="ticketID"
          label="票据编号"
          width="250"
        >
        </el-table-column>
        <el-table-column
          fixed="left"
          prop="guarantor"
          label="担保机构"
          width="250"
        >
        </el-table-column>
         <el-table-column
          fixed="left"
          prop="price"
          label="票据金额"
          width="250"
        >
        </el-table-column>
        <el-table-column
          fixed="left"
          prop="ownerOrg"
          label="票据接收方"
          width="250"
        >
        </el-table-column>
        <el-table-column
          fixed="left"
          prop="createTime"
          label="开单日期"
          width="250"
        >
        </el-table-column>
        <el-table-column
          fixed="left"
          prop="duedate"
          label="自动兑付日期"
          width="250"
        >
        </el-table-column>
        <el-table-column
          fixed="left"
          prop="description"
          label="备注"
          width="200"
        >
        </el-table-column>

        <el-table-column fixed="right" label="操作">
          <template slot-scope="scope">
            <el-button
              @click="detailButton(scope.row)"
              type="text"
              size="small"
            >
              详情
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
      tableData: [],
      prodName: "",
      createTime: "",
      prodNumber: "",
    };
  },
  methods: {
    handleChange() {},
    // clickCreate() {
    //   console.log();
    //   this.$router.push({ path: `/order/createOrder`});
    // },
    detailButton(row) {
      //跳转到详情
      console.log("row:",row);
      this.$router.push({ path: `/order/queryOrder`,query:{id:row.ticketID}});
    },
    editButton() {
      //跳转到编辑
      this.$router.push({ path: "/order/editOrder", params: { id: 1 } });
    },
    removeButton() {
      console.log("删除");
    },
  },
  mounted() {
    this.$axios.get("/console/queryNotExpiredTickets").then((data) => {
      console.log("data:", data.data.data);
      this.tableData = data.data.data;
    });
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
  width: 400px;
  margin-left: 50px;
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
</style>