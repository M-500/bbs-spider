
<template>
  <div>
    <el-button size="mini" class="button follow-btn" icon="el-icon-plus"  :class="{ 'is-followed': followed }" @click="follow" >
      {{ followed ? '已关注' : '关注' }}
    </el-button>
  </div>
</template>

<script >
export default {
  name: 'FollowBtn',
  props: {
    userId: {
      type: Number,
      required: true,
    },
    followed: {
      type: Boolean,
      default: false,
    },
  },
  methods:{
    follow(){

    },
  }
}
</script>


<style scoped>
.follow-btn {
  font-size: 12px;
  //height: 25px;
  background-color: #2469f6; // TODO
  border-color: #2469f6;
  color: #fff;

  &:hover,
  &.is-followed {
    background-color: #7ba5f9; // TODO
    border-color: #7ba5f9;
  }
  i {
    font-size: 12px;
    margin-right: 5px;
  }
}
</style>
