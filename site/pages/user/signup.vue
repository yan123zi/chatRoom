<template>
  <section>
    <el-card class="card">
      <el-form ref="form" :model="form" label-width="80px" class="form">
        <el-form-item label="用户名:">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="密码:">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="重复密码:">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="验证码:">
          <el-input v-model="form.name"></el-input>
          <img :src="form.url" style="width: 80px; height: 40px;position: absolute;margin-left: 20px" @click="showCaptcha"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">立即注册</el-button>
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
        name: '',
        url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
      }
    }
  },
  mounted() {
    this.showCaptcha()
  },
  methods: {
    onSubmit() {
      console.log('submit!');
    },
    async showCaptcha(){
      const codeObj=await this.$axios.get("/api/captcha/request")
      console.log(codeObj.data)
      console.log(codeObj.data.data)
      this.form.url=codeObj.data.data.captchaUrl
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