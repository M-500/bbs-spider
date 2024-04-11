<template>
  <div class="comments">
    <div class="comment">
      <div class="comment-item-left">
        <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
      </div>
      <div class="comment-item-main">
        <div class="comment-meta">
          <a href="/user/10024" class="comment-nickname">cityben</a>
          <div class="comment-meta-right">
            <time class="comment-time">17天前</time>
          </div>
        </div>
        <div class="comment-content-wrapper">
          <div class="comment-content content">怎么安装的？需要什么环境？</div>
        </div>
        <div class="comment-actions">
          <div class="comment-action-item">
            <i class="iconfont icon-like"></i>
            <span>点赞</span>
            <span>1</span>
          </div>
          <div class="comment-action-item">
            <i class="iconfont icon-comment"></i>
            <span>回复</span>
            <span>1</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
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
  },
  data() {
    return {
      showReplyCommentId: 0,
      reply: {
        commentId: 0,
        value: {
          content: '',
          imageList: [],
        },
      },
    }
  },
  computed: {
    user() {
      return this.$store.state.user.current
    },
    isLogin() {
      return this.$store.state.user.current != null
    },
  },
  methods: {
    append(data) {
      if (data) {
        this.$refs.commentsLoadMore.unshiftResults(data)
      }
    },
    async like(comment) {
      try {
        await this.$axios.post(`/api/comment/like/${comment.commentId}`)
        comment.liked = true
        comment.likeCount = comment.likeCount + 1
        this.$message.success('点赞成功')
      } catch (e) {
        if (e.errorCode === 1) {
          this.$msgSignIn()
        } else {
          this.$message.error(e.message || e)
        }
      }
    },
    switchShowReply(comment) {
      if (!this.user) {
        this.$msgSignIn()
        return
      }

      if (this.reply.commentId === comment.commentId) {
        this.hideReply(comment)
      } else {
        this.reply.commentId = comment.commentId
        setTimeout(() => {
          this.$refs[`editor${comment.commentId}`][0].focus()
        }, 0)
      }
    },
    hideReply(comment) {
      this.reply.commentId = 0
      this.reply.value.content = ''
      this.reply.value.imageList = []
    },
    async submitReply(parent) {
      try {
        const ret = await this.$axios.post('/api/comment/create', {
          entityType: 'comment',
          entityId: parent.commentId,
          content: this.reply.value.content,
          imageList:
            this.reply.value.imageList && this.reply.value.imageList.length
              ? JSON.stringify(this.reply.value.imageList)
              : '',
        })
        this.hideReply()
        this.appendReply(parent, ret)
        this.$message.success('发布成功')
      } catch (e) {
        if (e.errorCode === 1) {
          this.$msgSignIn()
        } else {
          this.$message.error(e.message || e)
        }
      }
    },
    onReply(parent, comment) {
      this.appendReply(parent, comment)
    },
    appendReply(parent, comment) {
      if (parent.replies && parent.replies.results) {
        parent.replies.results.push(comment)
      } else {
        parent.replies = {
          results: [comment],
        }
      }
    },
  },
}
</script>

<style scoped >
.comments {
  padding: 10px 10px 10px 0;
  font-size: 14px;

  .comment {
    display: flex;
    padding: 10px 0;

    &:not(:last-child) {
      border-bottom: 1px solid var(--border-color);
    }

    .comment-item-main {
      flex: 1 1 auto;
      margin-left: 16px;

      .comment-meta {
        display: flex;
        justify-content: space-between;
        .comment-nickname {
          font-size: 14px;
          font-weight: 600;
          color: var(--text-color);

          &:hover {
            color: var(--text-link-color);
          }
        }

        .comment-meta-right {
          .comment-time {
            font-size: 13px;
            color: var(--text-color3);
          }
          .comment-ip-area {
            font-size: 13px;
            color: var(--text-color3);
            margin-left: 10px;
          }
        }
      }

      .comment-content-wrapper {
        .comment-content {
          margin-top: 10px;
          margin-bottom: 0;
          color: var(--text-color);
          white-space: pre-wrap;
        }
        .comment-image-list {
          margin-top: 10px;

          img {
            width: 72px;
            height: 72px;
            line-height: 72px;
            cursor: pointer;
            &:not(:last-child) {
              margin-right: 8px;
            }

            object-fit: cover;
            transition: all 0.5s ease-out 0.1s;

            &:hover {
              transform: matrix(1.04, 0, 0, 1.04, 0, 0);
              backface-visibility: hidden;
            }
          }
        }
      }

      .comment-actions {
        margin-top: 10px;
        display: flex;
        align-items: center;

        .comment-action-item {
          line-height: 22px;
          font-size: 13px;
          cursor: pointer;
          color: var(--text-color3);
          user-select: none;

          &:hover {
            color: var(--text-link-color);
          }

          &.active {
            color: var(--text-link-color);
            font-weight: 500;
          }

          &:not(:last-child) {
            margin-right: 16px;
          }
        }
      }

      .comment-reply-form {
        margin-top: 10px;
      }

      .comment-replies {
        margin-top: 10px;
        // padding: 10px;
        background-color: var(--bg-color2);
      }
    }
  }

  .reply {
    display: flex;
  }
}
</style>
