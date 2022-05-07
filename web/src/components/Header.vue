<template>
  <el-header class="header">
    <div class="logo">数据标注存储系统</div>
    <!-- <el-input
        v-model="input"
        placeholder="Please input"
        class="input"
        size="large"
        clearable
    /> -->

    <div class="header-right">
      <div class="header-user-con">
        <Upload/>
        <el-dropdown
            class="user-name"
            trigger="click"
            @command="handleCommand"
        >
                    <span class="el-dropdown-link">
                        <el-avatar> {{ userName }} </el-avatar>
                    </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="user"
              >个人信息
              </el-dropdown-item
              >
              <el-dropdown-item command="anno"
              >个人标注
              </el-dropdown-item
              >
              <el-dropdown-item divided command="loginout"
              >退出登录
              </el-dropdown-item
              >
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </el-header>
</template>

<script lang="ts" setup>
const props = defineProps(['userName'])
</script>

<style scoped>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  border-bottom: solid 1px #dfe1e5;
}

.header .logo {
  float: left;
  width: 250px;
  line-height: 70px;
}

.header .input {
  float: left;
  width: 500px;
  height: 70px;
  display: flex;
  align-items: center;
  padding-left: 50px;
}

.header-right {
  float: right;
  padding-right: 50px;
}

.header-user-con {
  display: flex;
  height: 70px;
  align-items: center;
}

.user-name {
  margin-left: 10px;
}

.el-dropdown-link {
  /* color: #fff; */
  cursor: pointer;
}

.el-dropdown-menu__item {
  text-align: center;
}
</style>
