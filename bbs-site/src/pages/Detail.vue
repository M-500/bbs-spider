<template>
  <div class="mainContainer">
    <div class="pageLeft">
      <div class="pageContent">
        <div class="articlePage">
          <mavon-editor class="markdown" :value="get_mark_data()" :subfield="false" :defaultOpen="prop.defaultOpen" :toolbarsFlag="prop.toolbarsFlag" :editable="prop.editable" :scrollStyle="prop.scrollStyle"></mavon-editor>
        </div>

        <div class="articleComment">
          这里是评论区域
        </div>
      </div>
    </div>
    <div class="pageRight">
      你有毒
    </div>
  </div>
</template>


<script>
export default ({
  name: 'detail',
  data() {
    return {
      artile: {
        title: "关于我是个大傻逼",
        content: "# 傻子\n\n```\nPython\n```\n\n",
      },
    }
  },
  methods: {
    get_mark_data() {
      return "# 傻子\n\n```\nPython\n```\n\n"
    }
  },
  computed: {
    prop() {
      let data = {
        subfield: false,// 单双栏模式
        defaultOpen: 'preview',//edit： 默认展示编辑区域 ， preview： 默认展示预览区域 
        editable: false,
        toolbarsFlag: false,
        scrollStyle: true
      }
      return data
    }
  },
})
</script>

<style scoped>
.markdown {
  box-shadow: rgba(0, 0, 0, 0.1) 0px 0px 0px 0px !important;
}
.mainContainer {
  display: flex;
  justify-content: center;
  width: 100%;
}
.pageLeft {
  width: 58%;
  display: flex;
  margin-right: 18px;
  flex-direction: column; /* 纵向排列 */
}
.pageRight {
  width: 17%;
}

.articlePage {
  background-color: #f9f9f9;
  padding: 10px 24px 25px;
  border-radius: 5px;
}
</style>
