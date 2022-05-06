<template>
  <Header/>
  <el-main>
    <div class="flex-auto overflow-y-auto">
      <Waterfall
          :list="list"
          :row-key="options.rowKey"
          :gutter="options.gutter"
          :has-around-gutter="options.hasAroundGutter"
          :width="options.width"
          :breakpoints="options.breakpoints"
          :img-selector="options.imgSelector"
          :background-color="options.backgroundColor"
          :animation-effect="options.animationEffect"
          :animation-duration="options.animationDuration"
          :animation-delay="options.animationDelay"
          :lazyload="options.lazyload"
          :load-props="options.loadProps"
      >
        <template #item="{ item, url, index }">
          <div class="card">
            <LazyImg :url="url"/>
            <p class="text">{{ index }}</p>
          </div>
        </template>
      </Waterfall>
    </div>
    <el-backtop :right="100" :bottom="100"/>
  </el-main>
</template>

<script lang="ts" setup>
import {LazyImg, Waterfall} from "vue-waterfall-plugin-next";
import "vue-waterfall-plugin-next/style.css";
import loading from "~/assets/loading.png";
import error from "~/assets/error.png";
import {getData} from "~/api/algae";

// 瀑布流图片设置
function useWaterfall() {
  const options = reactive({
    // 唯一key值
    rowKey: "id",
    // 卡片之间的间隙
    gutter: 10,
    // 是否有周围的gutter
    hasAroundGutter: true,
    // 卡片在PC上的宽度
    width: 320,
    // 自定义行显示个数，主要用于对移动端的适配
    breakpoints: {
      1200: {
        // 当屏幕宽度小于等于1200
        rowPerView: 4,
      },
      800: {
        // 当屏幕宽度小于等于800
        rowPerView: 3,
      },
      500: {
        // 当屏幕宽度小于等于500
        rowPerView: 2,
      },
    },
    // 动画效果
    animationEffect: "animate__fadeInUp",
    // 动画时间
    animationDuration: 1000,
    // 动画延迟
    animationDelay: 300,
    // 背景色
    backgroundColor: "#ffffff",
    // imgSelector
    imgSelector: "src",
    // 加载配置
    loadProps: {
      loading,
      error,
    },
    // 是否懒加载
    lazyload: true,
  });
  return options;
}
const options = useWaterfall();
// 数据获取
const list = ref([]);
const fetchData = () => {
  getData({}).then((res) => {
    list.value = res.data;
  });
};
fetchData()
</script>

<style scoped></style>
