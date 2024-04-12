

<template>
  <el-dialog class="dialogPage" title="添加收藏" :visible.sync="dialogVisible" closeOnClickModal="false" width="30%" center>
    <div class="coll-body">
      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>

      <div class="item">
        <div class="title">
          <span>我的收藏</span>
          <i class="el-icon-lock"></i>
        </div>
        <div class="mark">108条内容</div>
        <el-button class="col_btn" size="mini">收藏</el-button>
      </div>
    </div>

    <div slot="footer" class="dialog-footer">
      <el-button size="mini" type="primary" @click="innerVisible = true">创建收藏夹</el-button>
    </div>

    <el-dialog width="30%" title="创建新收藏夹" center :visible.sync="innerVisible" append-to-body>

      <div class="coll-body">
        <el-form :model="collForm">
          <el-form-item>
            <el-input v-model="collForm.name" placeholder="收藏标题"></el-input>
          </el-form-item>
          <el-form-item>
            <el-input type="textarea" :autosize="{ minRows: 4, maxRows: 7}" placeholder="收藏描述（可选）" v-model="collForm.desc"></el-input>
          </el-form-item>
          <el-form-item>
            <el-radio-group v-model="collForm.isPublic">
              <el-radio>
                <span class="radio-title">公开</span>
                <span class="radio-mark">有其他人关注此收藏夹时不可设置为私密</span>
              </el-radio>
              <el-radio>
                <span class="radio-title">私密</span>
                <span class="radio-mark">只有你自己可以查看这个收藏夹</span>
              </el-radio>
            </el-radio-group>
          </el-form-item>

        </el-form>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button size="mini" @click="innerVisible = false">返回</el-button>
        <el-button size="mini" type="primary" @click="innerVisible = true">确认创建</el-button>
      </div>
    </el-dialog>
  </el-dialog>
</template>

<script>
export default {
  props: {
    dialogVisible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      innerVisible: false,
      collForm: {
        name: "",
        desc: "",
        isPublic: false
      }
    }
  }
}
</script>

<style scoped>
/deep/ .el-dialog__header {
  padding: 10px 5px 10px;
}
.dialogPage {
  border-radius: 8px;
}
.coll-body {
  height: 450px;
  padding: 10px 20px 10px;
  overflow-y: auto;
}
.item {
  padding: 13px 5px 8px;
  border-bottom: 1px solid black; /* 下边框样式：1px的实线，颜色为黑色 */
  position: relative; /* 父元素需要设置相对定位 */
}
.col_btn {
  position: absolute; /* 子元素设置绝对定位 */
  top: 10px; /* 子元素距离父元素顶部的距离 */
  right: 0px; /* 子元素距离父元素右边的距离 */
}
.title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  padding-bottom: 10px;
}

.radio-mark{
  color: grey;
}
</style>
