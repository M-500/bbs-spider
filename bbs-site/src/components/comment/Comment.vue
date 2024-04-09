<template>
  <div class="comment-component">
    <div class="comment-header">
      <span>评论</span>
      <span class="commentCnt">162</span>
      <div class="selector">
        <span class="selector-active">最热</span>
        <span>|</span>
        <span>最新</span>
      </div>
    </div>

    <template v-if="isLogin">
      <template >
        <comment-input
          ref="input"
          :entity-id="entityId"
          :entity-type="entityType"
          @created="commentCreated"
        />
      </template>
    </template>
    <div v-else class="comment-not-login">
      <div class="comment-login-div">
        请
        <a style="font-weight: 700" @click="toLogin">登录</a>后发表观点
      </div>
    </div>

    <comment-list/>
  </div>
</template>

<script>
import CommentInput from "./CommentInput.vue";
import CommentList from "./CommentList.vue";
export default {
  props: {
    entityType: {
      type: String,
      default: '',
      required: true,
    },
    entityId: {
      type: Number,
      default: 0,
      required: true,
    },
    commentsPage: {
      type: Object,
      default() {
        return {}
      },
    },
    commentCount: {
      type: Number,
      default: 0,
    },
  },
  components:{
    CommentInput,CommentList
  },
  computed: {
    isLogin() {
      // return this.$store.state.user.current != null
      return true
    },
    user() {
      return true
      // return this.$store.state.user.current
    },
    config() {
      return true
      // return this.$store.state.config.config
    },
    // 是否需要先邮箱认证
    isNeedEmailVerify() {
      return (
        this.config.createCommentEmailVerified &&
        this.user &&
        !this.user.emailVerified
      )
    },
  },
  methods: {
    commentCreated(data) {
      this.$refs.list.append(data)
      this.$emit('created')
    },
    reply(quote) {
      this.$refs.input.reply(quote)
    },
    toLogin() {
      this.$toSignin()
    },
  },
}
</script>
<style scoped>
.comment-component {
  background-color: var(--bg-color);
  border-radius: 3px;
  .comment-header {
    display: flex;
    padding-top: 20px;
    align-items: center;
    margin: 0 10px;
    color: #1f2d3d;
    font-size: 18px;
    font-weight: 500;
    .commentCnt{
      font-size: 14px;
    }
    span{
      margin-right: 5px;
    }
    .selector{
      margin-left: 30px;
      font-size: 12px;
      span{
        margin-left: 5px;
        color: #0ff123;
      }
      .selector-active{
        color: #1f2d3d;
      }
    }
  }

  .comment-not-login {
    margin: 10px;
    border: 1px solid var(--border-color);
    border-radius: 0;
    overflow: hidden;
    position: relative;
    padding: 10px;
    box-sizing: border-box;

    .comment-login-div {
      color: var(--text-color4);
      cursor: pointer;
      border-radius: 3px;
      padding: 0 10px;

      a {
        margin-left: 10px;
        margin-right: 10px;
      }
    }
  }
}
</style>
