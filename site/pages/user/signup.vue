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
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">立即创建</el-button>
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
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      }
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!');
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
  padding: auto;
}

.el-input {
  width: 50%;
}
</style>