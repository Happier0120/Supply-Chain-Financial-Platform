<template>
    <div>
        <el-form :rules="rules" ref="loginForm" :model="loginForm" class="loginContainer">
            <div class="loginTitle">
                系统登录
            </div>
            <el-form-item prop="username">
                <el-input type="text" v-model="loginForm.orgName" placeholder="请输入用户名">
                </el-input>
            </el-form-item>
            <el-form-item prop="password">
                <el-input type="password" v-model="loginForm.password" placeholder="请输入密码">
                </el-input>
            </el-form-item>

            <!--  <el-form-item prop="code">
          <el-input type="text" auto-complete="false" v-model="loginForm.code" placeholder="点击图片更换验证码" style="width: 250px;margin-right: 5px">
          </el-input>
         <img :src="captchaUrl">
      </el-form-item> -->

            <!-- <el-checkbox v-model="checked" class="loginRemember">记住我</el-checkbox> -->
            <el-button type="primary" style="width:100%" @click="submitLogin">登录</el-button>
        </el-form>
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
            checked: true,
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
                    alert('核心企业登录成功');
                    this.$store.commit("changeNavbar",true)  //改变Navbar显示与否的控制
                    sessionStorage.setItem('navbar_show',this.$store.state.navbar_show);
                    this.$router.push({ path: `/dashboard`, query: { account: this.loginForm.orgName } });
                    console.log(" this.$Navbar_show:", this.$store.state.navbar_show);
                    this.global.orgName = this.loginForm.orgName  //用全局变量接住登录的用户名字
                    this.global.type = "核心企业"
                    // this.$router.push({ path: `/Navbar`, query: { account: this.loginForm.orgName } });

                } else if (res.data.code == 200 && res.data.data == "1") {
                    alert('运输企业登录成功');
                    this.$store.commit("changeNavbar2",true)
                    sessionStorage.setItem('navbar_show',this.$store.state.navbar_show);
                    this.$router.push({ path: `/dashboard`, query: { account: this.loginForm.orgName } });
                    console.log(" this.$Navbar2_show:", this.$store.state.navbar2_show);

                    this.global.orgName = this.loginForm.orgName
                    this.global.type = "运输企业"

                } else {
                    alert('账户密码错误');
                }

            }).catch(err => {
                console.log("err.data:", err.data);

            })

        }
    }
};
</script>

<style lang="less" scoped>
.loginContainer {
    border-radius: 15px;
    background-clip: padding-box;
    margin: 180px auto;
    width: 350px;
    padding: 15px 35px 15px 35px;
    background: aliceblue;
    border: 1px solid blueviolet;
    box-shadow: 0 0 25px #f885ff;
}

.loginTitle {
    margin: 0px auto 48px auto;
    text-align: center;
    font-size: 40px;
}

.loginRemember {
    text-align: left;
    margin: 0px 0px 15px 0px;
}

body {
    //background-image: url("../assets/background.jpg") ;
    background-size: 100%;
}
</style>

