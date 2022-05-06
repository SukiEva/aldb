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
        <el-button round @click="dialogFormVisible = true">
          添加
          <i-ep-upload/>
        </el-button>
        <el-dropdown
            class="user-name"
            trigger="click"
            @command="handleCommand"
        >
                    <span class="el-dropdown-link">
                        <el-avatar> 用户 </el-avatar>
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

  <el-dialog v-model="dialogFormVisible" title="藻类图像上传" width="30%">
    <el-form :model="form">
      <el-form-item label="名称">
        <el-input v-model="form.name"/>
      </el-form-item>
      <el-form-item label="所属河流">
        <el-space
        >
          <el-select v-model="form.river" placeholder="选择河流">
            <el-option
                v-for="item in rivers"
                :key="item.name"
                :label="item.name"
                :value="item.name"
            />
          </el-select>
          <el-button
              type="primary"
              size="small"
              @click="riverFormVisible = true"
          >添加
          </el-button
          >
        </el-space
        >
      </el-form-item>
      <el-form-item label="图像">
        <el-upload drag :action="uploadUrl" class="avatar-uploader" :show-file-list="false" :on-success="afterUpload">
          <img v-if="form.src" :src="form.src" class="avatar"/>
          <el-icon v-else class="avatar-uploader-icon">
            <Plus/>
          </el-icon>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
            <span class="dialog-footer">
                <el-button @click="formCancel">取消</el-button>
                <el-button type="primary" @click="formSubmit">确认</el-button>
            </span>
    </template>
  </el-dialog>
  <el-dialog v-model="riverFormVisible" title="添加河流" width="30%" center>
    <el-form :model="river">
      <el-form-item label="名称">
        <el-input v-model="river.name"/>
      </el-form-item>
      <el-form-item label="位置">
        <el-input v-model="river.address"/>
      </el-form-item>
    </el-form>
    <template #footer>
            <span class="dialog-footer">
                <el-button @click="riverFormCancel">取消</el-button>
                <el-button type="primary" @click="riverFormSubmit"
                >添加</el-button
                >
            </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {addAlga, addRiver, getRivers} from "~/api/algae";
import type {UploadFile, UploadFiles, UploadProps} from 'element-plus'
import {Plus} from '@element-plus/icons-vue'

const input = ref("");

// 添加图像
const dialogFormVisible = ref(false);
const form = reactive({
  name: "",
  src: "",
  river: "",
});
const clearForm = () => {
  form.name = "";
  form.src = "";
  form.river = "";
};
const formCancel = () => {
  dialogFormVisible.value = false;
  clearForm();
};
const formSubmit = () => {
  addAlga(form).then((res) => {
    if (res.code != 200) {
      ElMessage.error("上传失败")
      return
    }
    ElMessage.success("上传成功")
    formCancel()
  })
}
// 上传图片
const disabled = ref(false)
const uploadUrl = "http://127.0.0.1:8080/api/upload"
const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
  if (response.code != 200) {
    ElMessage.error("上传失败")
    return
  }
  form.src = response.data.url
}
// 获取河流
const rivers = ref([]);
const fetchRivers = () => {
  getRivers({}).then((res) => {
    rivers.value = res.data;
  });
};
fetchRivers();
// 添加河流
const river = reactive({
  name: "",
  address: "",
});
const riverFormVisible = ref(false);
const riverFormSubmit = () => {
  if (river.name == "" || river.address == "") {
    ElMessage.warning("请填写完整信息")
    return;
  }
  addRiver(river).then((res) => {
    console.log(res)
    if (res.code != 200) {
      return;
    }

    ElMessage.success("添加河流成功")
    fetchRivers();
    riverFormVisible.value = false;
  });
};
const riverFormCancel = () => {
  river.name = "";
  river.address = "";
  riverFormVisible.value = false;
};
</script>

<style scoped>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  /* color: #fff; */
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

.avatar-uploader .avatar {
  width: 100%;
  height: 100%;
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>