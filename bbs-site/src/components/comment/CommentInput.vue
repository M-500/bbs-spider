<template>
  <div class="box-normal">
    <div class="reply-box-avatar">
      <div class="avatar">
        <img src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" alt="">
      </div>
      <div class="reply-box-wrap">
        <div class="textarea-wrap">
          <textarea name="" id="" cols="30" rows="1" placeholder="诶哟不错哦，发条评论吧"></textarea>
          <el-button size="mini" class="pubBtn">发布</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TextEditor from "./TextEditor.vue";
export default {
  components:{
    TextEditor
  },
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
  },
  data() {
    return {
      value: {
        content: '', // 内容
        imageList: [],
      },
      sending: false, // 发送中
      quote: null, // 引用的对象
    }
  },
  computed: {
    btnName() {
      return this.sending ? '正在发表...' : '发表'
    },
    user() {
      return this.$store.state.user.current
    },
  },
  methods: {
    async create() {
      if (!this.value.content) {
        this.$message.error('请输入评论内容')
        return
      }
      if (this.sending) {
        console.log('正在发送中，请不要重复提交...')
        return
      }
      if (this.$refs.simpleEditor && this.$refs.simpleEditor.isOnUpload()) {
        this.$message.warning('正在上传中...请上传完成后提交')
        return
      }
      this.sending = true
      try {
        const data = await this.$axios.post('/api/comment/create', {
          contentType: this.contentType,
          entityType: this.entityType,
          entityId: this.entityId,
          content: this.value.content,
          imageList:
            this.value.imageList && this.value.imageList.length
              ? JSON.stringify(this.value.imageList)
              : '',
          quoteId: this.quote ? this.quote.commentId : '',
        })
        this.$emit('created', data)

        this.value.content = ''
        this.value.imageList = []
        this.quote = null
        this.$refs.simpleEditor.clear()
        this.$message.success('发布成功')
      } catch (e) {
        console.error(e)
        this.$message.error(e.message || e)
      } finally {
        this.sending = false
      }
    },
    reply(quote) {
      this.quote = quote
      this.$refs.commentEditor.scrollIntoView({
        block: 'start',
        behavior: 'smooth',
      })
    },
    cancelReply() {
      this.quote = null
    },
  },
}
</script>

<style scoped >
.box-normal{
  margin-top: 10px;
  font-weight: 400;
  height: 50px;
}
.reply-box-avatar{
  display: flex;
  justify-content: flex-start;
  margin-right: 20px;
}
.avatar{
  margin-right: 10px;
  img{
    height: 50px;
    border-radius: 50%;
  }
}
.reply-box-wrap{
  width: 100%;
  display: flex;
  flex-direction: column; /* 纵向排列 */
  justify-content: center;
}
.textarea-wrap{
  width: 100%;
  display: flex;
  flex-direction: column; /* 纵向排列 */
  justify-content: center;
  position: relative; /* 使得子元素相对于父元素进行定位 */
  textarea{
    width: 100%;
    padding: 10px 0 10px;
    resize: none;
        overflow: hidden;
  }
  .pubBtn{
    position: absolute; /* 绝对定位 */
        right: 3px;
        background-color:cornsilk;
  }
}
</style>
