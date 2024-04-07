<script>
import { PubArticleListAPI } from "@/api/article/reader";
export default {
  name: 'home',
  data () {
    return  {
      form:{
        page_num: 0,
        page_size: 15,
      },
      articles:[
        {},
        {},
        {},
      ],
    }
  },
  methods: {
    toDetail(id) {
      this.$router.push('/detail/' + id)
    },
    getArticleList(){
      PubArticleListAPI(this.form).then((res) => {
        this.articles = res
      }).catch((e) => {
        this.$message({
          message: e.msg,
          type: "error",
        });
      });
    }
  },
  created() {
    this.getArticleList()
  }
}
</script>

<template>
  <div class="pages">
    <div class="left-main">
      <div class="left-main-container">
        <div class="tag-list">
          <div class="tag-box">
            <img :src="require('@/assets/icon/a-shijianzuijin.svg')" alt="">
            <div class="title">最新</div>
          </div>

          <div class="tag-box">
            <img :src="require('@/assets/icon/a-shijianzuijin.svg')" alt="">
            <div class="title">推荐</div>
          </div>

          <div class="tag-box">
            <img :src="require('@/assets/icon/a-shijianzuijin.svg')" alt="">
            <div class="title">关注</div>
          </div>

          <div class="tag-box">
            <img :src="require('@/assets/icon/a-shijianzuijin.svg')" alt="">
            <div class="title">收藏</div>
          </div>


          <div class="tag-box">
            <img :src="require('@/assets/icon/a-shijianzuijin.svg')" alt="">
            <div class="title">我的</div>
          </div>
        </div>
      </div>
    </div>

    <div class="right-main">
      <div class="left-item">
        <div class="article-list" v-for="(item,index) in articles" :key="index" >
          <div class="article-item" @click="toDetail(item.id)">
            <div class="avatar">
              <el-avatar :size="50" src="https://mlog.club/images/avatars/80.png"></el-avatar>
            </div>
            <div class="article-main-content">
              <div class="top">
                <div class="uInfo">大喵喵酱</div>
                <div class="uTime"> 发布于{{item.ctime}}</div>
              </div>
              <div class="content">
                <div class="title">{{item.title}}</div>
                <div class="abs">{{item.summary}}...</div>
              </div>
              <div class="bottom">
                <div class="left">
                  <div class="box">
                    <img :src="require('@/assets/icon/dianzan.svg')" alt="">
                    <span>赞 4</span>
                  </div>
                  <div class="box">
                    <img :src="require('@/assets/icon/pinglun.svg')" alt="">
                    <span>评论 4</span>
                  </div>
                  <div class="box">
                    <img :src="require('@/assets/icon/guankan.svg')" alt="">
                    <span>查看 4</span>
                  </div>
                </div>
                <div class="right">
                  <el-tag type="success">摸鱼</el-tag>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="right-item">
        <div class="rightBox">
          <div class="notice">
            <div class="cardTitle">公告</div>
            <div class="cardContent">欢迎访问 码农俱乐部 点击这里设置您的邮箱 可以接收站内跟帖、回复邮件提醒，不错过任何一条消息。</div>
          </div>
          <div class="ranking">
            <div class="rankingTitle">积分排行</div>
            <div class="rankingContent">
              <div class="rankingCard">
                <div class="left">
                  <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<style scoped>
.pages {
  display: flex;
  width: 100%;
}
.left-main {
  display: flex;
  justify-content: flex-end;
  width: 20%;
}
.right-main {
  width: 80%;
  display: flex;
  justify-content: flex-start;
}
.left-item {
  width: 60%;
}
.right-item {
  width: 25%;
  margin-left: 10px;
}
.left-main-container {
  .tag-list {
    padding: 10px 0 10px;
    display: flex;
    flex-direction: column; /* 纵向排列 */
    background-color: #f9f9f9;
    .tag-box {
      border-radius: 1px;
      border: aliceblue 1px solid;
      margin-bottom: 10px;
      margin-right: 20px;
      padding: 8px 15px 8px;
      display: flex;
      justify-content: flex-start;
      img {
        width: 20px;
        background-color: white;
      }
      .title {
        margin-left: 18px;
        font-size: 13px;
      }
    }
  }
}

.article-list {
  display: flex;
  flex-direction: column; /* 纵向排列 */
  .article-item {
    margin: 0 10px 15px;
    background-color: #f9f9f9;
    display: flex;
    border: #f9f9ff 1px solid;
    border-radius: 4px;
    padding: 10px;
    justify-content: flex-start; /* 横向排列并靠左对齐 */
    .avatar img {
      width: 50px;
    }
    .article-main-content {
      margin-left: 10px;
      .top {
        margin-bottom: 5px;

        display: flex;
        justify-content: space-between;
        .uInfo {
          color: #70727c;
          font-weight: bold;
        }
        .uTime {
        }
      }
      .content {
        .title {
          font-weight: bold;
          font-size: 17px;
          margin-bottom: 2px;
        }
      }
    }
    .bottom {
      margin-top: 10px;
      display: flex;
      justify-content: space-between;
      .left {
        display: flex;
        justify-content: flex-start;
        .box {
          margin-right: 15px;
          align-items: center; /* 垂直居中 */
          display: flex;
          img {
            display: flex;
            justify-content: center; /* 水平居中 */
            align-items: center; /* 垂直居中 */
            width: 15px;
            height: auto;
          }
          span {
            font-size: 12px;
          }
        }
      }
      .right {
        display: flex;
        justify-content: flex-start;
      }
    }
  }
}

.rightBox {
  display: flex;
  margin-bottom: 10px;
  flex-direction: column; /* 纵向排列 */
}
.notice {
  background-color: #f9f9f9;
  padding: 10px 24px 25px;
  border-radius: 5px;
  display: flex;

  flex-direction: column; /* 纵向排列 */
}
.cardTitle {
  padding-bottom: 12px;
  border-bottom: #70727c 1px solid;
}
.cardContent {
  padding-top: 8px;
  font-weight: 400;
  font-size: small;
}
.ranking {
  margin-top: 12px;
  background-color: #f9f9f9;
  padding: 10px 24px 25px;
  border-radius: 5px;
  display: flex;
  flex-direction: column; /* 纵向排列 */
}
.rankingTitle {
  padding-bottom: 12px;
  border-bottom: #70727c 1px solid;
}
</style>
