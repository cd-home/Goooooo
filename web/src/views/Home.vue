<script setup>

</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-header></el-header>
      <el-container>
        <el-aside>
        </el-aside>
        <el-container>
          <el-main></el-main>
          <el-footer></el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<style lang="scss" scoped>
  .el-container {
    .el-aside {
      position: absolute;
      width: 24rem;
      height: 90vh;
      left: 1vw;
      border: 1px solid rgb(196, 187, 187);
      border-radius: 5px;
    }
  }
</style>

<script>
import Header from "@/components/Header.vue";
export default {
  data() {
    return {}
  },
  components: {
    Header
  }
}
</script>