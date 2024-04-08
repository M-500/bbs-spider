<template>
  <div class="mainContainer">
    <div class="pageLeft">
      <div class="pageContent">
        <div class="articlePage">
          <div v-if="isLoading">Loading</div>
          <div v-else><div v-html="article.content" class="markdown-body" style="text-align:left;margin-bottom:50px"></div></div>

          <div class="pageBottom">
            <div class="pageBottomBox">
              <img :src="require('@/assets/icon/guankan.svg')" alt="">
              <span>浏览</span>
              <span>171</span>
            </div>
            <div class="pageBottomBox">
              <img :src="require('@/assets/icon/like.svg')" alt="">
              <span>点赞</span>
            </div>
            <div class="pageBottomBox">
              <img :src="require('@/assets/icon/collect.svg')" alt="">
              <span>收藏</span>
              <span>171</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="pageRight">
      你有毒
      <div class="userCard">
      </div>
    </div>
  </div>
</template>


<script>
import { PubArticleDetailAPI } from "@/api/article/reader";
import article from "./edit/Article.vue";
// import 'mavon-editor/dist/css/index.css'
import {marked} from "marked";
export default ({
  name: 'detail',
  data() {
    return {
      isLoading: true,
      markdownOption: {
        bold: true, // 粗体
      },
      id: this.$route.params.id,
      article: {
        id: "",
        title: "",
        content: "",
      },
      comments: [
      ]
    }
  },
  methods: {
    articleContent(){
      PubArticleDetailAPI(this.id).then((res) => {
        article.id = res.id;
        article.title = res.title;
        article.content= marked(res.content);
      }).catch((e) => {
        this.$message({
          message: e.msg,
          type: "error",
        });
      });
    }
  },
  mounted() {
    PubArticleDetailAPI(this.id).then((res) => {
      this.article.id = res.id;
      this.article.title = res.title;
      this.article.content= marked(res.content);
      this.isLoading = false;
    }).catch((e) => {
      this.isLoading = false;
      this.$message({
        message: e.msg,
        type: "error",
      });
    });
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
  border-radius: 2px;
}
.articleComment {
  margin-top: 5px;
  background-color: #f9f9f9;
  padding: 10px 24px 25px;
  border-radius: 2px;
}
.pageBottom{
  display: flex;
  justify-content: space-around;

}
.pageBottomBox{
  display: flex;
  justify-content: flex-start;
  align-items: center; /* 在交叉轴上居中 */
  img{
    height: 20px;
  }
}
</style>
