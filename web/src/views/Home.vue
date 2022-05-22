<script setup>

</script>
<template>
  <div class="common-layout">
    <el-container>
      <el-header>Header</el-header>
      <el-container>
        <el-aside width="200px">Aside</el-aside>
        <el-container>
          <el-main>Main</el-main>
          <el-footer>Footer</el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<style lang="scss" scoped>
  .el-header {
    width: 100%;
    height: 50px;
    background: pink;
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