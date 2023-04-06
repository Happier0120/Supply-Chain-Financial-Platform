<template>
    <div class="bill-detail-content">
        <div class="bill-detail-title-bar">
            <span>Ticket detail information</span>
        </div>
        <div class="detail-result-content">
            <div class="container">
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
                            <span>{{ item.timestamp }}</span>
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
</template>

<script>
export default {
    data() {
        return {
            infoData_public: {},
            infoData_privacy: {},
            infoData_blockchain: {},
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
                .get("/console/queryTicketPublicById", { params: { ticketID: ticketID } })
                .then((data) => {
                    console.log("infoData_public:", data.data.data);
                    this.infoData_public = data.data.data;
                });
        },
        initData2() {
            // console.log(this.$root.id);
            let ticketID = this.$route.query.id;

            this.$axios
                .get("/console/queryTicketPrivateById", { params: { ticketID: ticketID } })
                .then((data) => {
                    console.log("infoData_privacy:", data.data.data);
                    this.infoData_privacy = data.data.data;
                });
        },
        initData3() {
            // console.log(this.$root.id);
            let ticketID = this.$route.query.id;

            this.$axios
                .get("/console/queryTicketHistoryById", { params: { ticketID: ticketID } })
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

<style src="../../style/component/ticketDetail.css"></style>

<!-- <style lang="scss" scoped>
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

        .wea-type-group {
            margin-bottom: 0px;

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
                margin-top: 0px;
                margin-left: 50px;
                background-color: #dfd2ff;
            }
        }
    }
}

.page {
    margin-left: 220px;
}
</style> -->