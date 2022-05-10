<template>
  <Header :user-name="userName"/>
  <el-main>
    <Main :user-email="userEmail"/>
  </el-main>
</template>

<script lang="ts" setup>
// 获取用户信息
import {useRouter} from "vue-router";
import {getUser} from "~/api/user";

const router = useRouter()
const userEmail = sessionStorage.getItem("UserEmail") as string
if (userEmail == null) {
  ElMessage.error("未登录")
  router.push({name: "Login"})
}
const userName = ref("用户")
const fetchUser = () => {
  getUser(userEmail).then((res) => {
    userName.value = res.data.name
  })
}
fetchUser()
</script>

<style scoped>
</style>
