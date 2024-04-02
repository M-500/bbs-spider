<template>
  <div class="mainContainer">
    <div class="pageLeft">
      <div class="pageContent">
        <div class="articlePage">
          <mavon-editor class="markdown" :value="get_mark_data()" :subfield="false" :defaultOpen="prop.defaultOpen" :toolbarsFlag="prop.toolbarsFlag" :editable="prop.editable" :scrollStyle="prop.scrollStyle"></mavon-editor>
        </div>

        <div class="articleComment">
          <div class="commentForm">
            <div class="commentTitle">3条评论</div>
            <div class="commentData">
              <el-input v-model="input" type="textarea" :rows="2" placeholder="请输入内容"></el-input>
            </div>
          </div>
          <div class="comment-list">
            <el-card v-for="comment in comments" :key="comment.id" class="comment-card">
              <div slot="header" class="clearfix">
                <span>评论人：{{ comment.author }}</span>
                <el-avatar :src="comment.avatar" size="small" shape="circle" class="avatar"></el-avatar>
              </div>
              <div class="comment-content">{{ comment.content }}</div>
              <div class="child-comments">
                <comment-list :comments="comment.children" v-if="comment.children && comment.children.length > 0"></comment-list>
              </div>
            </el-card>
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
export default ({
  name: 'detail',
  data() {
    return {
      artile: {
        title: "关于我是个大傻逼",
        content: "# 傻子\n\n```\nPython\n```\n\n",
      },
      comments: [
        {
          id: 1,
          author: 'Alice',
          avatar: 'https://via.placeholder.com/50',
          content: '这是一条评论。'
        },
        {
          id: 2,
          author: 'Bob',
          avatar: 'https://via.placeholder.com/50',
          content: '这是另一条评论。',
          children: [
            {
              id: 21,
              author: 'Charlie',
              avatar: 'https://via.placeholder.com/50',
              content: '这是 Bob 的回复。'
            },
            {
              id: 22,
              author: 'David',
              avatar: 'https://via.placeholder.com/50',
              content: '这是另一条回复。',
              children: [
                {
                  id: 221,
                  author: 'Eva',
                  avatar: 'https://via.placeholder.com/50',
                  content: '这是 David 的回复。'
                }
              ]
            }
          ]
        },
        {
          id: 3,
          author: 'Eve',
          avatar: 'https://via.placeholder.com/50',
          content: '这是另一条评论。'
        }
      ]
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
  border-radius: 2px;
}
.articleComment {
  margin-top: 5px;
  background-color: #f9f9f9;
  padding: 10px 24px 25px;
  border-radius: 2px;
}
</style>
