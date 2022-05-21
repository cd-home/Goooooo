<template>
  <div class="user">
    <Header></Header>
    <div class="auth">
      <router-link to="/">首页</router-link>
      <router-link to="/user/login">登陆</router-link>
      <router-link to="/user/register">注册</router-link>
      <button @click="count++">{{ count }}</button>
      <button @click="handler">测试</button>
    </div>
    <div class="box" :style="{top: TOP}">
    </div>
    <router-view></router-view>
  </div>
</template>

<style lang="scss" scoped>
.user {
  width: 100%;
  height: 400px;
  button {
    width: 60px;
    height: 20px;
  }
}
.auth {
  width: 800px;
  height: 50px;
}

.box {
  position: absolute;
  top: -110px;
  left: 620px;
  right: 0;
  bottom: 0;
  width: 240px;
  height: 120px;
  border-radius: 10px;
  background-color: #e99e9e;
  transition: 1s top;
}
li {
  list-style: none;
  width: 500px;
  height: 20px;
  margin-top: 10px;
  background-color: #78ddb2;
}
</style>

<script>
import Header from '../components/Header.vue'
import {login} from "../api/user";

export default {
  data() {
    return {
      count: 0,
      TOP: '-500px'
    }
  },
  methods: {
    handler() {
      login({}).then(data => {
        console.log(22222)
        alert(data.data)
      }).catch(error => {
        alert(111)
        alert(error)
      })
    }
  },
  components: {
    Header
  }
}
</script>