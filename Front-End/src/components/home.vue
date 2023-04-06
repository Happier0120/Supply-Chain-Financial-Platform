<template>
    <div class="home-content">
        <div class="cmp-container">
            <div class="avatar-container">
                <div class="img-container">
                    <img src="../images/avatar14.png" alt="enterprise-icon">
                </div>
                <div class="detail-container">
                    <span>{{this.global.orgName}}</span>
                    <div class="member-container">
                        <img src="../images/diamond.png" alt="diamond">
                        <span>Member</span>
                    </div>
                </div>
            </div>
            <button class="view-btn" @click="handleClick">
                <span>View details</span>
                <img src="../images/right-arrow.png" alt="right-arrow">
            </button>
        </div>
        <div class="mid-container">
            <div class="item-record">
                <img src="../images/outline-3dcube (1).png" alt="blockNum">
                <div class="word-info">
                    <span class="rec-title">Block number</span>
                    <span class="record">{{this.infoData.blockNumber}}</span>
                </div>
            </div>
            <div class="item-record">
                <img src="../images/channel.png" alt="channelNum">
                <div class="word-info">
                    <span class="rec-title">Channel number</span>
                    <span class="record">1</span>
                </div>
            </div>
            <div class="item-record">
                <img src="../images/node.png" alt="nodeNum">
                <div class="word-info">
                    <span class="rec-title">Node number</span>
                    <span class="record">{{this.infoData.nodesNumber}}</span>
                </div>
            </div>
            <div class="item-record">
                <img src="../images/transaction.png" alt="transactionNum">
                <div class="word-info">
                    <span class="rec-title">Transactions number</span>
                    <span class="record">{{this.infoData.transactionCount}}</span>
                </div>
            </div>
        </div>
        <div class="bot-container">
            <div class="left-area" id="left-chart"></div>
            <div class="right-area" id="right-chart"></div>
        </div>
    </div>
</template>

<script>
export default {
    data () {
        return {
            infoData: '',
            infoData2: '',            
            chart1Data: '',
            chart2Data: ''

        }
    },
    created() {
        console.log("跳转到票据详情");
        // this.initData();
        // this.initData2();
    },
    methods: {
        initData() {
            // 返回数据可用
            this.$axios
                .get("/console/queryChaincodeInfo")
                .then((data) => {
                this.infoData = data.data.data;
                // this.$forceUpdate()
            });
        },

        initData2() {

            return new Promise((resolve, reject) => {
                this.$axios
                    .get("/console/queryTicketsInfo")
                    .then((data) => {
                        console.log(data);
                        console.log(data.data.data.todayPrice);
                        this.infoData2 = data.data.data;                
                        var bill_data = [];
                        var amount_data = [];
                        bill_data.push(data.data.data.todayTicketsNumber);
                        bill_data.push(data.data.data.ticketsNumber);
                        bill_data.push(data.data.data.todayExpiredTicketsNumber);
                        bill_data.push(data.data.data.expiredTicketsNumber);

                        amount_data.push(data.data.data.todayPrice);
                        amount_data.push(data.data.data.price);
                        amount_data.push(data.data.data.todayExpiredPrice);
                        amount_data.push(data.data.data.expiredPrice);

                        this.$forceUpdate()
                        this.chart1Data = bill_data;
                        this.chart2Data = amount_data;
                        resolve();
                });

            })
            
            
        },
        drawChart1() {
            this.chartObj1 = this.$echarts.init(document.getElementById("left-chart"));
            this.chartObj1.setOption({
                title: {
                    text: 'Ticket statistics',
                    left: 'center',
                    textStyle: {
                        color: '#283593',
                        fontFamily: '"Alkatra", cursive',
                        fontFamily: '"Bungee Spice", cursive',
                        fontFamily: '"Exo 2", sans-serif'
                    }
                },
                tooltip: {},
                xAxis: {
                    data: ['Tickets issued today', 'Tickets issued (total)', 'Tickets expired today', 'Tickets exprired (total)'],
                    axisLabel: {
                        interval: 0
                    }
                },
                yAxis: {},
                series: [
                    {
                    name: 'Num',
                    type: 'bar',
                    data: this.chart1Data
                    }
                ]
            });
        },
        drawChart2() {
            this.chartObj2 = this.$echarts.init(document.getElementById("right-chart"));
            this.chartObj2.setOption({
                title: {
                    text: 'Amount statistics',
                    left: 'center',
                    textStyle: {
                        color: '#283593',
                        fontFamily: '"Alkatra", cursive',
                        fontFamily: '"Bungee Spice", cursive',
                        fontFamily: '"Exo 2", sans-serif'
                    }
                },
                tooltip: {},
                xAxis: {
                    data: ['Amount issued today', 'Amount issued (total)', 'Amount due today', 'Amount due (total)'],
                    axisLabel: {
                        interval: 0,
                        // rotate: 30
                    }
                },
                yAxis: {},
                series: [
                    {
                    name: 'Num',
                    type: 'bar',
                    data: this.chart2Data
                    }
                ]
            });
        },
        handleClick() {
            this.$notify.info({
            title: 'Notice',
            message: 'This feature is not yet available.'
        });
        }
    },
    mounted() {
        this.initData();

        this.initData2().then(() => {
            this.drawChart1();
            this.drawChart2();
        })

    }
}
</script>
<style>
  @import url('https://fonts.googleapis.com/css2?family=Alkatra&family=Bungee+Spice&family=Exo+2&display=swap');
</style>
<style src="../style/component/home.css"></style>