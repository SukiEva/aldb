<template>
    <div id="userLayout">
        <div class="login_panle">
            <div class="login_panle_form">
                <div class="login_panle_form_title">
                    <!-- <img class="login_panle_form_title_logo" :src="" alt /> -->
                    <p class="login_panle_form_title_p">数据标注存储系统</p>
                </div>
                <el-form
                    ref="loginForm"
                    :model="loginFormData"
                    :rules="rules"
                    @keyup.enter="submitForm"
                >
                    <el-form-item prop="username">
                        <el-input
                            v-model="loginFormData.username"
                            placeholder="请输入用户名"
                        >
                            <template #suffix>
                                <span class="input-icon">
                                    <el-icon>
                                        <user />
                                    </el-icon>
                                </span>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input
                            v-model="loginFormData.password"
                            :type="lock === 'lock' ? 'password' : 'text'"
                            placeholder="请输入密码"
                        >
                            <template #suffix>
                                <span class="input-icon">
                                    <el-icon>
                                        <component
                                            :is="lock"
                                            @click="changeLock"
                                        />
                                    </el-icon>
                                </span>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item prop="captcha">
                        <div class="vPicBox">
                            <el-input
                                v-model="loginFormData.captcha"
                                placeholder="请输入验证码"
                                style="width: 60%"
                            />
                            <div class="vPic">
                                <img
                                    v-if="picPath"
                                    :src="picPath"
                                    alt="请输入验证码"
                                    @click="loginVerify()"
                                />
                            </div>
                        </div>
                    </el-form-item>
                    <el-form-item>
                        <el-button
                            type="primary"
                            style="width: 46%"
                            size="large"
                            @click="checkInit"
                            >注册</el-button
                        >
                        <el-button
                            type="primary"
                            size="large"
                            style="width: 46%; margin-left: 8%"
                            @click="submitForm"
                            >登 录</el-button
                        >
                    </el-form-item>
                </el-form>
            </div>
            <div class="login_panle_right" />
            <div class="login_panle_foot">
                <div class="copyright">
                    <CopyRight />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive, ref } from "vue";

const loginFormData = reactive({
    username: "admin",
    password: "123456",
    captcha: "123",
    captchaId: "",
});
</script>

<style lang="scss" scoped>
@import "~/styles/login.scss";
</style>
