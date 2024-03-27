<script setup>
export default {
  name: 'home',
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
        </div>
      </div>
    </div>

    <div class="right-main">
      <div class="left-item">
        <div class="article-list">
          <div class="article-item">
            <div class="avatar">
              <el-avatar :size="50" src="https://mlog.club/images/avatars/80.png"></el-avatar>
            </div>
            <div class="article-main-content">
              <div class="top">
                <div class="uInfo">大喵喵酱</div>
                <div class="uTime"> 发布于2024-02-22 09:54:00</div>
              </div>
              <div class="content">
                <div class="title">Go 语言社区系统 BBS-GO 4.0.0 发布，升级vue3、nuxt3</div>
                <div class="abs">文档地址 帮助文档：https://bbs-go.com 官网交流：https://mlog.club 问题反馈：https://mlog.club/topics/node/3 功能建议收集：https://mlog.club/topic/609 git...</div>
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

          <div class="article-item">
            <div class="avatar">
              <img src="https://mlog.club/images/avatars/80.png" alt="">
            </div>
            <div class="article-main-content">
              <div class="top">
                <div class="uInfo">大喵喵酱</div>
                <div class="uTime"> 发布于2024-02-22 09:54:00</div>
              </div>
              <div class="content">
                <div class="title">Go 语言社区系统 BBS-GO 4.0.0 发布，升级vue3、nuxt3</div>
                <div class="abs">文档地址 帮助文档：https://bbs-go.com 官网交流：https://mlog.club 问题反馈：https://mlog.club/topics/node/3 功能建议收集：https://mlog.club/topic/609 git...</div>
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

          <div class="article-item">
            <div class="avatar">
              <img src="https://mlog.club/images/avatars/80.png" alt="">
            </div>
            <div class="article-main-content">
              <div class="top">
                <div class="uInfo">大喵喵酱</div>
                <div class="uTime"> 发布于2024-02-22 09:54:00</div>
              </div>
              <div class="content">
                <div class="title">Go 语言社区系统 BBS-GO 4.0.0 发布，升级vue3、nuxt3</div>
                <div class="abs">文档地址 帮助文档：https://bbs-go.com 官网交流：https://mlog.club 问题反馈：https://mlog.club/topics/node/3 功能建议收集：https://mlog.club/topic/609 git...</div>
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

          <div class="article-item">
            <div class="avatar">
              <img src="https://mlog.club/images/avatars/80.png" alt="">
            </div>
            <div class="article-main-content">
              <div class="top">
                <div class="uInfo">大喵喵酱</div>
                <div class="uTime"> 发布于2024-02-22 09:54:00</div>
              </div>
              <div class="content">
                <div class="title">Go 语言社区系统 BBS-GO 4.0.0 发布，升级vue3、nuxt3</div>
                <div class="abs">文档地址 帮助文档：https://bbs-go.com 官网交流：https://mlog.club 问题反馈：https://mlog.club/topics/node/3 功能建议收集：https://mlog.club/topic/609 git...</div>
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

          <div class="article-item">
            <div class="avatar">
              <img src="https://mlog.club/images/avatars/80.png" alt="">
            </div>
            <div class="article-main-content">
              <div class="top">
                <div class="uInfo">大喵喵酱</div>
                <div class="uTime"> 发布于2024-02-22 09:54:00</div>
              </div>
              <div class="content">
                <div class="title">Go 语言社区系统 BBS-GO 4.0.0 发布，升级vue3、nuxt3</div>
                <div class="abs">文档地址 帮助文档：https://bbs-go.com 官网交流：https://mlog.club 问题反馈：https://mlog.club/topics/node/3 功能建议收集：https://mlog.club/topic/609 git...</div>
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

      </div>
    </div>
  </div>
</template>

<style scoped>

.pages{
  display: flex;
  width: 100%;
}
.left-main{
  display:flex;
  justify-content: flex-end;
  width: 20%;
}
.right-main{
  width: 80%;
  .left-item{
    width: 60%;
  }
  .right-item{
    width: 40%;
  }
}
.left-main-container{
  //width: 30%;
  .tag-list{
  padding: 10px 20px 10px;
    display: flex;
    flex-direction: column; /* 纵向排列 */
    background-color: #f9f9f9;
    .tag-box{
      border-radius: 1px;
      border: aliceblue 1px solid;
      margin-bottom: 10px;
      padding: 8px 15px 8px;
      display: flex;
      justify-content: flex-start;
      img{
        width: 20px;
        background-color: white;
      }
      .title{
        margin-left: 18px;
        font-size: 13px;
      }
    }
  }
}


.article-list{
  display: flex;
  flex-direction: column; /* 纵向排列 */
  .article-item{
    margin: 0 10px 15px;
    background-color: #f9f9f9;
    display: flex;
    border: #f9f9ff 1px solid;
    border-radius: 4px;
    padding: 10px;
    justify-content: flex-start; /* 横向排列并靠左对齐 */
    .avatar img{
      width: 50px;
    }
    .article-main-content{
      margin-left: 10px;
      .top{

        margin-bottom: 5px;

        display: flex;
        justify-content: space-between;
        .uInfo{
          color: #70727c;
          font-weight: bold;
        }
        .uTime{

        }
      }
      .content{
        .title{
          font-weight: bold;
          font-size: 17px;
          margin-bottom: 2px;
        }
      }
    }
    .bottom{
      margin-top: 10px;
      display: flex;
      justify-content: space-between;
      .left{
        display: flex;
        justify-content: flex-start;
        .box{
          margin-right: 15px;
          align-items: center; /* 垂直居中 */
          display: flex;
          img{
            display: flex;
            justify-content: center; /* 水平居中 */
            align-items: center; /* 垂直居中 */
            width: 15px;
            height: auto;
          }
          span{
            font-size: 12px;
          }
        }
      }
      .right{
        display: flex;
        justify-content: flex-start;
      }
    }
  }
}

</style>
