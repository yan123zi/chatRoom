<template>
  <section>
    <el-row :gutter="0" class="el-row">
      <el-col :span="17" class="el-col">
        <div class="grid-content bg-purple"></div>
      </el-col>
      <el-col :span="7" class="el-col">

        <el-dropdown trigger="click" class="avatar-div" v-if="user">
          <el-avatar icon="el-icon-user-solid" :size="58" class="avatar"></el-avatar>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item icon="el-icon-plus">黄金糕</el-dropdown-item>
            <el-dropdown-item icon="el-icon-circle-plus" divided>狮子头</el-dropdown-item>
            <el-dropdown-item icon="el-icon-circle-plus-outline">螺蛳粉</el-dropdown-item>
            <el-dropdown-item icon="el-icon-check">双皮奶</el-dropdown-item>
            <el-dropdown-item icon="el-icon-circle-check">蚵仔煎</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <div v-if="!user" class="log-reg">
          <a class="log-reg-text" href="/user/signin">登录</a>
          <el-divider direction="vertical"></el-divider>
          <a class="log-reg-text" href="/user/signup">注册</a>
        </div>
      </el-col>
    </el-row>
  </section>
</template>

<script>
export default {
  data() {
    return {

    };
  },
  methods: {

  },
  computed:{
    user(){
      return this.$store.state.user.current
    }
  }
}
</script>

<style>
.el-row {
  height: 70px;
  border-bottom: #dedae2 1px solid
}

.el-col {
  display: block;
  height: 100%;
  float: left
}

.avatar-div {
  position: relative;
  height: 100%;
  display: block;
  float: right;
  width: 60px;
  cursor: pointer;
}

.avatar {
  position: absolute;
  margin:auto;
  line-height: 60px;
  top: 0px;
  bottom: 0px;
  left: 0px;
  right: 0px;
}
.log-reg{
  /*position: absolute;*/
  display: block;
  float: right;
  line-height: 60px;
  margin-right: 20px;
}
.log-reg-text{
  cursor: pointer;
  text-decoration: none;
  color: #333;
}
.log-reg-text:hover{
  color: #409EFF;
}
</style>