<template>
  <div class="bg">
    <div class="home-box">
      <Search/>
      <div class="blockTxns-container">
        <IndexBlock/>
        <IndexTxns/>
      </div>
    </div>
  </div>
</template>

<script>
  import IndexBlock from "../../components/HomeBlocks";
  import IndexTxns from "../../components/HomeTxns";
  import Search from '../../components/Search'
  import Chart from "../../components/HomeChart";

  export default {
    name: "Home",

    components: {Search, IndexTxns, IndexBlock, Chart},
  }
</script>

<style lang="less" rel="stylesheet/less">
.bg {
/*  background-image: url("/static/img/bg.png"); */
/*  background-color: #2149A7; */
  background-size: 100%;
  background-repeat: no-repeat;
  margin-top: -90px;
  /*距离底部的高度，添加到blockTxns-container类上，因为要显示阴影*/
  .home-box {
    width: 1040px;
    margin: 0 auto;
    overflow: hidden;
    /* background-color: #2149A7; */
    .blockTxns-container {
      padding: 0 20px 100px;
      display: flex;
      justify-content: space-between;
      margin-top: 75px;
    }
  }
}

@media screen and (max-width:480px){
  .bg {
    width: 100%;
    background-image: none;
    .home-box {
      width: 100%;
      margin: 0 auto;
      overflow: hidden;
      /* background-color: #2149A7; */
      .blockTxns-container {
        padding: 0 10px 24px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        margin-top: 24px;
      }
    }
  }
}


</style>
