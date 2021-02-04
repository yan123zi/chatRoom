<template>
  <section>
    <el-row :gutter="0" style="height: 70px;border-bottom: #dedae2 1px solid">
      <el-col :span="17" style="display: block;height: 100%;float: left">
        <div class="grid-content bg-purple"></div>
      </el-col>
      <el-col :span="7" style="display: block;float: left;height: 100%">
        <div style="position: relative;height:100%;display:block;float: right;width:60px">
          <el-avatar icon="el-icon-user-solid" style="position: absolute;margin: auto;top:0px;bottom: 0px;"></el-avatar>
        </div>
      </el-col>
    </el-row>
  </section>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: '1',
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style>
.clear{
  margin: 0px 0px;
}
</style>