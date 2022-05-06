<script setup>

</script>
<template>
  <div class="home">
    <Header></Header>
    <router-link to="/user">登陆或者注册</router-link>
  </div>
</template>

<style lang="scss" scoped>
.home {
  width: 100%;
  height: 400px;
  background-color: #559b90;
}
</style>

<script>
import Header from "../components/Header.vue";

export default {
  data() {
    return {}
  },
  components: {
    Header
  }
}
</script>