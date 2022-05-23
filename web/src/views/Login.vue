<template>
  <el-container>
      <el-main>
         <el-form
          :label-position="labelPosition"
          label-width="100px"
          :model="form"
          style="max-width: 460px"
          :rules="rules"
          ref="form"
        >
          <el-form-item label="Account" prop="account">
            <el-input v-model="form.account" autofocus="true" />
          </el-form-item>
          <el-form-item label="Password" prop="password">
          <el-input v-model="form.password" type="password" show-password oncopy="false" />
          </el-form-item>
          <el-button type="" size="small" @click="LoginHand">登录</el-button>
        </el-form>
      </el-main>
    </el-container>
</template>

<style lang="scss">
  .el-container {
    background-image: url("@/assets/loginbg.jpeg");
    background-size: cover;
    width: 100%;
    height: 100%;
    position: relative;
    .el-main {
      user-select: none;
      position: absolute;
      top: 10vh;
      right: 8vh;
      width: 56vh;
      height: 60vh;
      .el-form {
        margin-top: 10vh;
        .el-form-item {
           .el-form-item__label {
            color: #fff;
          }
          .el-input {
            font-size: 1.5rem;
            --el-input-font-color: rgb(59, 137, 201);
            --el-input-background-color: transparent;
            --el-input-border-radius: 5px;
          }
        }
        .el-button {
          margin-left: 28rem;
          background-color: transparent;
          font-size: 1rem;
          color: #fff;
        }
      }
    }
  }
</style>

<script>
import { reactive, ref } from 'vue'

const labelPosition = ref('right')

const form = reactive({
  account: '',
  password: '',
})

const rules = reactive({
  account: [
    {required: true, message: '账号不能为空', type: "string", trigger: 'blur'},
    {min: 4, max: 18, message: '账号长度必须4-18位', trigger: 'blur'}
  ],
  password: [
    {required: true, message: '密码不能为空', trigger: 'blur'},
    // {min: 6, max: 18, message: '密码长度必须4-18位', trigger: 'blur'},
    // {
    //   pattern: /^(?=.*\d)(?=.*[a-zA-Z])(?=.*[~!@#$%^&*])[\da-zA-Z~!@#$%^&*]{6,18}$/,
    //   message: '密码长度6-18位, 必须包含数字、字母、特殊字符',
    //   trigger: 'blur'
    // },
  ]
})

export default {
  data() {
    return {
      form,
      labelPosition,
      rules,
    }
  },
  methods: {
   LoginHand() {
      //  let body = new FormData();
      //  body.append("account", this.form.account);
      //  body.append("password", this.form.password);
      this.$api.login(this.form).then(response => {
      this.$notify.SuccessNotify(response.message)
      this.$store.commit('login', response.data)
      console.log(this.$store.state.user)
      this.$router.push("/home")
     }).catch(error => {
       this.$notify.ErrorNotify(error)
     })
   }
  }
}
</script>