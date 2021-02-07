<template>
  <section>
    <el-card class="card">
      <el-form ref="form" :model="form" label-width="80px" class="form">
        <el-form-item label="用户名:">
          <el-input v-model="form.nickname"></el-input>
        </el-form-item>
        <el-form-item label="密码:">
          <el-input v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item label="重复密码:">
          <el-input v-model="form.rePassword"></el-input>
        </el-form-item>
        <el-form-item label="验证码:">
          <el-input v-model="form.captchaCode"></el-input>
          <img :src="form.captchaUrl" style="width: 80px; height: 40px;position: absolute;margin-left: 20px"
               @click="showCaptcha"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="signUp">立即注册</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </section>
</template>

<script>
export default {
  name: "signup",
  data() {
    return {
      form: {
        nickname: '',
        password: '',
        rePassword: '',
        captchaId: '',
        captchaUrl: '',
        captchaCode: '',
      }
    }
  },
  mounted() {
    this.showCaptcha()
  },
  methods: {
    async signUp() {
      try {
        await this.$store.dispatch('user/signup',{
          nickname:this.form.nickname,
          password:this.form.password,
          rePassword:this.form.rePassword,
          captchaId:this.form.captchaId,
          captchaCode:this.form.captchaCode
        })
      }catch (e) {
        this.$message.error(e.message||e)
      }
    },
    async showCaptcha() {
      try {
        const codeObj = await this.$axios.get("/api/captcha/request")
        this.form.captchaUrl = codeObj.captchaUrl
        this.form.captchaId=codeObj.captchaId
      }catch (e) {
        this.$message.error(e.message||e)
      }
    }
  }
}
</script>

<style scoped>
.card {
  height: 600px;
  width: 70%;
  margin: 50px auto;
}

.form {
  display: block;
  margin: 50px auto;
  width: 80%;
  padding-left: 25%;
}

.el-input {
  width: 50%;
}
</style>