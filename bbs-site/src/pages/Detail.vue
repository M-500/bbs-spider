<template>
  <div class="mainContainer">
    <div class="pageLeft">
      <div class="pageContent">
        <div class="articlePage">
          <div class="pageAuthor">
            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
            <div class="pageAuthorInfo">
              <div class="username">吴琳琳</div>
              <div class="aretDate">发布于 2012</div>
            </div>
          </div>
          <div class="pageTop">
            <h1 class="Title">{{article.title}}</h1>
            <div class="Title">{{article.summary}}</div>
          </div>
          <div v-if="isLoading">Loading</div>
          <div v-else><div v-html="article.content" class="markdown-body" style="text-align:left;margin-bottom:50px"></div></div>

          <div class="pageBottom">
            <div class="pageBottomBox">
              <img :src="require('@/assets/icon/guankan.svg')" alt="">
              <span>浏览</span>
              <span>171</span>
            </div>
            <div class="pageBottomBox">
              <img v-if="!likeForm.like" :src="require('@/assets/icon/like.svg')" alt="" @click="doLike">
              <img v-else :src="require('@/assets/icon/liked.svg')" alt="" @click="doLike">
              <span>点赞</span>
            </div>
            <div class="pageBottomBox">
              <img :src="require('@/assets/icon/collect.svg')" alt="" @click="doCollect">
              <span>收藏</span>
              <span>171</span>
            </div>
          </div>
        </div>


      </div>
      <div class="commentContent">
        <Comment :entityType=1 :entityId=2 :commentsPage={} :commentCount=12></Comment>
      </div>
    </div>
    <div class="pageRight">
      <div class="userCard">
        <div class="isCenter">
          <div class="avatarImg">
            <img src="https://cdn.jsdelivr.net/gh/M-500/picx-images-hosting@master/425380192_25678980908368088_6576232889963075542_n.8ad1dchumr.webp" alt="">
          </div>
          <div class="author-info__name">王木木</div>
          <div class="author-info__description">一个写代码的小学生</div>
        </div>
        <div class="card-info-data">
          <a href="/archives/">
            <div class="headline">文章</div>
            <div class="length-num">93</div>
          </a>
          <a href="/tags/">
          <div class="headline">标签</div>
          <div class="length-num">17</div>
        </a>
          <a href="/categories/">
          <div class="headline">分类</div>
            <div class="length-num">34</div>
          </a></div>
      </div>
    </div>

    <el-dialog
      title="添加到收藏夹"
      :visible.sync="dialogVisible"
      width="30%"
      center
      :before-close="handleClose">
<!--      <span><el-checkbox v-model="checked3" label="备选项1" border size="medium"></el-checkbox>-->
<!--        </span>-->
<!--      <span><el-checkbox v-model="checked4" label="备选项2" border size="medium"></el-checkbox></span>-->
<!--        <span slot="footer" class="dialog-footer">-->

<!--        <el-button size="mini" type="primary" @click="dialogVisible = false">确 定</el-button>-->
<!--      </span>-->
    </el-dialog>
  </div>
</template>


<script>
import Comment from "../components/comment/Comment.vue";
import { PubArticleDetailAPI,ArticleLikeDetailAPI } from "@/api/article/reader";
import article from "./edit/Article.vue";
import 'mavon-editor/dist/css/index.css'
import {marked} from "marked";
export default ({
  name: 'detail',
  components:{
    Comment
  },
  data() {
    return {
      dialogVisible: false,
      isLoading: true,
      markdownOption: {
        bold: true, // 粗体
      },
      likeForm:{
        id: parseInt(this.$route.params.id),
        like: true,
      },
      id: this.$route.params.id,
      article: {
        id: "",
        title: "",
        summary:"",
        content: "",
      },
      comments: [
      ]
    }
  },
  methods: {
    handleClose(){
      this.dialogVisible = false
    },
    doCollect(){
      this.dialogVisible = true
    },
    doLike(){
      this.likeForm.like = true
      console.log(this.likeForm)
      ArticleLikeDetailAPI(this.likeForm).then( (res) =>{
        console.log(res,"点赞结果")
      }).catch((e) => {
        this.$message({
          message: e.msg,
          type: "error",
        });
      });
    },
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
      this.article.summary = res.summary;
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
  width: 20%;
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
    transition: transform 0.2s ease;
    margin-right: 5px;
  }
  img:hover {
    transform: scale(1.5);
    animation: shake 0.2s ease-in-out;
  }
  @keyframes shake {
    0% { transform: translateX(-2px); }
    50% { transform: translateX(2px); }
    100% { transform: translateX(-2px); }
  }
}
.pageAuthor{
  display: flex;
  justify-content: flex-start;
  margin-bottom: 10px;
}
.pageAuthorInfo{
  margin-left: 10px;
  display: flex;
  justify-content: space-between;
  flex-direction: column; /* 设置主轴方向为垂直 */
  .username{
    font-size: 13px;
  }
  .aretDate{
    font-size: 11px;
    color: #3e606f;
  }
}

.userCard{
  padding: 20px 24px;
  border-radius: 8px;
  background: #fff;;
}
.isCenter{
  width: 100%;
}
.author-info__name {
  display: flex;
  justify-content: center;
  margin-top: 10px;
  font-weight: 500;
  font-size: 1.57em;
}
.author-info__description {
  display: flex;
  justify-content: center;
  margin-top: 10px;
}
.avatarImg{
  display: flex;
  justify-content: center;
  img{
    border-radius: 50%;
    height: 100px;
    overflow-clip-margin: content-box;
    overflow: clip;
  }
}
.card-info-data {
  margin: 14px 0 4px;
  display: flex;
  justify-content: space-around;
}
.card-info-data a{
  text-decoration: none; /* 去掉下划线 */
  color: #99a9bf;
  display: flex;
  flex-direction: column; /* 设置主轴方向为垂直 */
  align-content: center;
}

.card-info-data a .length-num{
  margin-top: 0.32em;
  color: #1f2d3d;
  font-size: 1.4em;
}
</style>
