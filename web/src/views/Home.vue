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
          <el-main>
            <el-button icon="el-icon-download" @click="downloadStream">下载</el-button>
          </el-main>
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
export default {
  data() {
    return {}
  },
  methods: {
    downloadStream() {
      let query = {"filename": "1533726291776245760Untitled Diagram.svg"}
      this.$api.downloadFileStream(query)
      .then((file) => {
        // Blob => URL
        let url = window.URL.createObjectURL(new Blob([file]));
        let link = document.createElement('a');
        link.style.display = 'none';
        link.href = url;
        link.setAttribute('download', query.filename);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link); // 下载完成移除元素
        window.URL.revokeObjectURL(url); // 释放掉blob对象
        this.$notify.SuccessNotify(response.message);
      })
      .catch(error => {
        alert(error)
        this.$notify.WarnNotify(error)
      })
    }
  }
}
</script>