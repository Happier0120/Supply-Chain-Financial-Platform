<template>
    <div class="content">
        <div class="left">
            <div class="icon-container">
                <img src="../images/供应链-copy.png" alt="供应链icon">
                <span class="prodIcon">Supply-Chain Financial Platform</span>
            </div>
        </div>
        <div class="right">
            <div class="form-container">
                <el-form :rules="rules" ref="loginForm" :model="loginForm" class="loginContainer">
                    <div class="loginTitle">
                        <span class="top-title">Sign in</span>
                        <span class="notation notation1">Welcome to logistics supply-chain platform.</span>
                        <span class="notation notation2">Sign in to experience</span>
                    </div>
                    <el-form-item prop="username">
                        <span class="notation">Username</span>
                        <el-input type="text" v-model="loginForm.orgName">
                        </el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <span class="notation">Password</span>
                        <el-input type="password" v-model="loginForm.password">
                        </el-input>
                    </el-form-item>
        
                    <!--  <el-form-item prop="code">
                  <el-input type="text" auto-complete="false" v-model="loginForm.code" placeholder="点击图片更换验证码" style="width: 250px;margin-right: 5px">
                  </el-input>
                 <img :src="captchaUrl">
              </el-form-item> -->
        
                    <el-checkbox v-model="checked" class="loginRemember">I agree to the terms of service</el-checkbox>
                    <el-button type="primary" style="width:100%" @click="submitLogin">Sign in</el-button>
                </el-form>
                <div class="sign-up-container">
                    <span class="notation">New to SCFP? </span>
                    <button class="goForSign" @click="handleClick">Create an account</button>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
export default {
    name: "Login",
    data() {
        return {
            captchaUrl: "",
            loginForm: {
                orgName: "",
                password: "",
                //code:''
            },
            checked: false,
            rules: {
                orgName: [{ required: true, message: "请输入用户名", trigger: "blur" }, { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
                ],
                password: [{ required: true, message: "请输入密码", trigger: "blur" }, , { min: 6, message: '密码长度要大于6', trigger: 'blur' }],
                //code:[{required:true,message:"请输入验证码",trigger:"blur"}],
            }

        }
    },
    methods: {
        created() { },
        submitLogin() {
            console.log("登录");
            this.$axios({
                method: 'post',
                url: '/console/submit',  ///开放接口
                headers: {
                    'Content-Type': 'application/json'  //设置Content-Type为jsonss
                },
                data: this.loginForm,
            }).then(res => {
                console.log("res.data:", res.data);
                if (res.data.code == 200 && res.data.data == "0") {
                    alert('Core enterprise login successful');
                    // this.$store.commit("changeNavbar",true)  //改变Navbar显示与否的控制
                    // sessionStorage.setItem('navbar_show',this.$store.state.navbar_show);
                    // this.$router.push({ path: `/dashboard`, query: { account: this.loginForm.orgName } });
                    this.$router.push({ path: '/home' });
                    // console.log(" this.$Navbar_show:", this.$store.state.navbar_show);
                    this.global.orgName = this.loginForm.orgName  //用全局变量接住登录的用户名字
                    this.global.type = "核心企业"
                    // this.$router.push({ path: `/Navbar`, query: { account: this.loginForm.orgName } });

                } else if (res.data.code == 200 && res.data.data == "1") {
                    alert('Transportation enterprise login successful');
                    // this.$store.commit("changeNavbar2",true)
                    // sessionStorage.setItem('navbar_show',this.$store.state.navbar_show);
                    this.$router.push({ path: `/home`});
                    // console.log(" this.$Navbar2_show:", this.$store.state.navbar2_show);

                    this.global.orgName = this.loginForm.orgName
                    this.global.type = "运输企业"

                } else {
                    alert('账户密码错误');
                }

            }).catch(err => {
                console.log("err.data:", err.data);

            })

        },
        handleClick() {
            this.$notify.info({
                title: 'Notice',
                message: 'This feature is not yet available.'
            });
        }
    }
};
</script>

<style lang="less" scoped>
// .loginContainer {
//     border-radius: 15px;
//     background-clip: padding-box;
//     margin: 180px auto;
//     width: 350px;
//     padding: 15px 35px 15px 35px;
//     background: aliceblue;
//     border: 1px solid blueviolet;
//     box-shadow: 0 0 25px #f885ff;
// }

// .loginTitle {
//     margin: 0px auto 48px auto;
//     text-align: center;
//     font-size: 40px;
// }

.loginRemember {
    text-align: left;
    margin: 0px 0px 15px 0px;
}

body {
    //background-image: url("../assets/background.jpg") ;
    background-size: 100%;
}
</style>

<style src="../style/login.css"></style>
