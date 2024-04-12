
<template>
  <div class="main">
    <div class="mainPage">
      <div class="topCard">
        <userCard></userCard>
      </div>
      <div class="content">
        <div class="leftContet">
          <div class="widget">
            <div class="widget-header">
              <span>个人成就</span>
            </div>
            <div class="widget-content extra-info">
              <ul>
                <li>
                  <span>积分</span>
                  <br>
                  <a href="">12</a>
                </li>
              </ul>
              <ul>
                <li>
                  <span>话题</span>
                  <br>
                  <a href="">1</a>
                </li>
              </ul>
              <ul>
                <li>
                  <span>评论</span>
                  <br>
                  <a href="">12</a>
                </li>
              </ul>
              <ul>
                <li>
                  <span>收藏</span>
                  <br>
                  <a href="">12</a>
                </li>
              </ul>
            </div>
          </div>

          <div class="widget">
            <div class="widget-header">
              <span>个人资料</span>
            </div>
            <div class="widget-content stable">
              <div class="str">
                <div class="slabel">昵称</div>
                <div class="svalue">damn bro 名利</div>
              </div>
              <div class="str">
                <div class="slabel">签名</div>
                <div class="svalue">古希腊掌管bug的神</div>
              </div>
            </div>
          </div>

          <div class="widget">
            <div class="widget-header">
              <span>粉丝</span>
            </div>
            <div class="widget-content stable">
              <div class="str">
                <div class="slabel">昵称</div>
                <div class="svalue">damn bro 名利</div>
              </div>
              <div class="str">
                <div class="slabel">签名</div>
                <div class="svalue">古希腊掌管bug的神</div>
              </div>
            </div>
          </div>

        </div>
        <div class="rightContet">
          <el-tabs v-model="activeTabName" @tab-click="handleClick">
            <el-tab-pane label="话题" name="topic">
              <span slot="label"><i class="el-icon-chat-line-square"></i> 话题</span>
              话题
            </el-tab-pane>
            <el-tab-pane label="文章" name="article">
              <span slot="label"><i class="el-icon-tickets"></i> 文章</span>
              <article-list :articles="articleList"></article-list>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>

  </div>

</template>
<script>
import userCard from "../../components/userCenter/userCard.vue";
import ArticleList from "../../components/article/ArticleList.vue";
export default {
  name: 'userCenter',
  components:{
    userCard,ArticleList
  },
  data () {
    return  {
      activeTabName:'article',
      articleList:[
        {
          articleId: 1,
          title:"测试",
          summary:"猜测摘要",
          createTime:"2024-12-21",
          user:{
            id:1,
            nickname:"叼毛"
          }
        },
        {
          articleId:2,
          title:"测试",
          summary:"猜测摘要",
          createTime:"2024-12-21",
          user:{
            id:1,
            nickname:"叼毛"
          }
        }
      ],
    }
  },
  methods:{
    handleClick(){},
  }
}
</script>

<style scoped>
.main{
  width: 100%;
  display: flex;
  justify-content: center;
}
.mainPage{
  width: 80%;
}
.content{
  display: flex;
  justify-content: flex-start;
  width: 100%;
}
.leftContet{
  width: 30%;
  margin-right: 10px;
}
.rightContet{
  width: 70%;
  background-color: #fff;
  padding: 0 10px 10px;
  border-radius: calc(.25rem - 1px);
  height: 100%;
  .tabs {
    margin-bottom: 5px;
  }
}
.tabs-warp {
}
// 伪类选择器，去除第一个元素的样式
.leftContet .widget:not(:first-child){
  margin-top: 10px;
}
.widget{
  background-color: #ffffff;
  border-radius: calc(.25rem - 1px);
  padding: 0 12px;
  margin-top: 10px;
}

.widget-header{
  align-items: center;
  border-bottom: 1px solid #e9e9e9;
  display: flex;
  font-size: 16px;
  font-weight: 700;
  justify-content: space-between;
  padding: 8px 0;
}

.widget-content{
  padding: 8px 0;
  //word-break: break-all;
}
.extra-info{
  display: flex;
  justify-content: space-around;
  width: 100%;
  ul {
    list-style: none;
    display: flex;
    flex-direction: column; /* 纵向排列 */
    justify-content: center;
    text-align: center;
    li {

      text-align: center;
      width: 100%;
      span {
        font-weight: 400;
        font-size: 13px;
        color: #70727c;;
      }
      a{
        color: #485fc7;
        cursor: pointer;
        text-decoration: none;
      }
    }
  }
}
.stable .str:not(:last-child){
  border-bottom: 1px solid #e9e9e9;
}
.stable .str {
  align-items: center;
  display: flex;
  font-size: 14px;
  justify-content: space-between;
  padding: 12px 0;
}

.stable .str .svalue {
  color: #333;
  margin-left: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  word-break: break-all;
}

</style>
