
<template>
  <div class="main">
    <div class="mainPage">
      <div class="topCard">
        <userCard></userCard>
      </div>
      <div class="content">
        <div class="leftContet">
          <el-tabs v-model="activeTabName" @tab-click="handleClick">
            <el-tab-pane label="话题" name="topic">
              <span slot="label">话题 <span class="tab-num">1</span></span>
              话题
            </el-tab-pane>
            <el-tab-pane label="文章" name="article">
              <span slot="label">文章<span class="tab-num">1</span></span>
              <article-list :articles="articleList"></article-list>
            </el-tab-pane>
            <el-tab-pane label="收藏" name="collect">
              <span slot="label">收藏<span class="tab-num">1</span></span>
              <article-list :articles="articleList"></article-list>
            </el-tab-pane>
            <el-tab-pane label="关注" name="follow">
              <span slot="label">关注<span class="tab-num">1</span></span>
              <article-list :articles="articleList"></article-list>
            </el-tab-pane>
          </el-tabs>
        </div>
        <div class="rightContet">
          
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
          summary:"猜测摘要，啊四季豆请问啊商山早行尽快哦地哦趣味，aszxdmn0quiwe",
          createTime:"2024-12-21",
          user:{
            id:1,
            nickname:"叼毛"
          }
        },
        {
          articleId:2,
          title:"测试",
          summary:"猜测摘要1，如果灭有必要，奥都请问i哦额u请问请问饿哦i几千万i哦额及其",
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
  width: 70%;
  background-color: #fff;
  padding: 0 10px 10px;
  border-radius: calc(.25rem - 1px);
  height: 100%;
}
.rightContet{
  width: 30%;
  padding: 0 0 10px 10px;
  border-radius: calc(.25rem - 1px);
  height: 100%;
  .tabs {
    margin-bottom: 5px;
  }
}
.tabs-warp {
}
// 伪类选择器，去除第一个元素的样式
.rightContet .widget:not(:first-child){
  margin-top: 10px;
}
.widget{
  background-color: #ffffff;
  border-radius: calc(.25rem - 1px);
  padding: 0 12px;
  margin-bottom: 10px;
  /* margin-top: 10px; */
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

/deep/ .el-tabs__item span {
  color: #191b1f;
    /* display: inline-block; */
    font-size: 16px;
    line-height: 22px;
    padding: 14px 0;
    position: relative;
    text-align: center;
}
.tab-num{
  color: #9196a1;
  font-size: 14px;
    font-weight: 300;
    line-height: 20px;
    margin-left: 6px;
}
</style>
