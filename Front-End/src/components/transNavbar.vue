<template>
    <div class="navbar-container">
        <el-row class="navbar">
            <el-col>
                <el-menu
                :default-active="paths[0]"
                router
                class="el-menu-vertical-demo"
                @open="handleOpen"
                @close="handleClose">
                    <el-submenu index="1">
                        <template slot="title">
                        <!-- <i class="el-icon-location"></i> -->
                        <img src="../images/home.png" alt="Home" class="icon-img">
                        <span>Home</span>
                        </template>
                        <el-menu-item :index="paths[0]">Transportation Enterprise Home</el-menu-item>
                        <el-menu-item :index="paths[1]">Tickets Traceability</el-menu-item>
                    </el-submenu>
                    <el-submenu index="2">
                        <template slot="title">
                        <!-- <i class="el-icon-location"></i> -->
                        <img src="../images/Bill.png" alt="Unpaid bills" class="icon-img">
                        <span>Untransferred tickets</span>
                        </template>
                        <el-menu-item :index="paths[2]">Ticket list</el-menu-item>
                    </el-submenu>
                    <el-submenu index="3">
                        <template slot="title">
                        <!-- <i class="el-icon-location"></i> -->
                        <img src="../images/outline-ticket-expired.png" alt="paid bills" class="icon-img">
                        <span>Transferred tickets</span>
                        </template>
                        <el-menu-item :index="paths[3]">Ticket list</el-menu-item>
                    </el-submenu>
                    <el-submenu index="4">
                        <template slot="title">
                        <!-- <i class="el-icon-location"></i> -->
                        <img src="../images/outline-ticket-expired.png" alt="paid bills" class="icon-img">
                        <span>Expired Tickets</span>
                        </template>
                        <el-menu-item :index="paths[4]">Ticket list</el-menu-item>
                    </el-submenu>
                </el-menu>
            </el-col>
        </el-row>
    </div>
</template>

<script>
export default {
    data() {
        return {
            paths: [
                '/home',
                '/ticket/traceticket',
                '/ticket/ticketList/notTransferred',
                '/ticket/ticketList/transferred',
                '/ticket/ticketList/expired'
            ]
        }
    },
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      }
    }
};
</script>

<style src="../style/component/navBar.css"></style>