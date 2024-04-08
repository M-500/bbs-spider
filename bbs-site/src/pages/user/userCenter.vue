
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
        <div class="leftContet">
          右边边的东西
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
</style>
