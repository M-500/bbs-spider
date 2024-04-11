
<template>
  <div class="main">
    <div class="mainPage">
      <div class="topCard">
        <userCard></userCard>
      </div>
      <div class="content">
        <div class="leftContet">
          左边的东西
        </div>
        <div class="rightContet">
          傻逼吗这不是
        </div>
      </div>
    </div>

  </div>

</template>
<script>
import userCard from "../../components/userCenter/userCard.vue";
export default {
  name: 'userCenter',
  components:{
    userCard
  },
  data () {
    return  {

    }
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
}
.rightContet{
  width: 70%;
}
</style>
