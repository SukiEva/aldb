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
                    添加 <i-ep-upload />
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
                                >个人信息</el-dropdown-item
                            >
                            <el-dropdown-item command="anno"
                                >个人标注</el-dropdown-item
                            >
                            <el-dropdown-item divided command="loginout"
                                >退出登录</el-dropdown-item
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
                <el-input v-model="form.name" />
            </el-form-item>
            <el-form-item label="所属河流">
                <el-space
                    ><el-select v-model="form.river" placeholder="选择河流">
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
                        >添加</el-button
                    ></el-space
                >
            </el-form-item>
            <el-form-item label="图像">
                <el-upload drag action="" :limit="1">
                    <el-icon class="el-icon--upload"
                        ><i-ep-upload-filled
                    /></el-icon>
                    <div class="el-upload__text">
                        Drop file here or <em>click to upload</em>
                    </div>
                    <template #tip>
                        <div class="el-upload__tip">jpg or png files</div>
                    </template>
                </el-upload>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="formCancel">Cancel</el-button>
                <el-button type="primary" @click="dialogFormVisible = false"
                    >Confirm</el-button
                >
            </span>
        </template>
    </el-dialog>
    <el-dialog v-model="riverFormVisible" title="添加河流" width="30%" center>
        <el-form :model="river">
            <el-form-item label="名称">
                <el-input v-model="river.name" />
            </el-form-item>
            <el-form-item label="位置">
                <el-input v-model="river.address" />
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
import { addRiver, getRivers } from "~/api/algae";
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
// 获取河流
const rivers = ref([]);
const fetchRivers = () => {
    getRivers().then((res) => {
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
        ElMessage({
            message: "请填写完整信息",
            type: "warning",
        });
        return;
    }
    addRiver(river).then((ele) => {
        if (ele.code != 200) {
            return;
        }

        ElMessage({
            message: "添加河流成功",
            type: "success",
        });
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
</style>
