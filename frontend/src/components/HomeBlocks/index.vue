<template>
  <div class="homeBlocks-box">
    <div class="my-list-banner">
      <div class="my-list-title">
        BLOCKS
      </div>
      <router-link to="/blocks">View All</router-link>
    </div>
    <ul class="my-list-body">
      <li my-list-wrap v-for="block in indexBlockList">
        <p>
          <img src="/static/img/block.png" alt="">
          Block#:  
          <router-link :to="{path:`/block/${block.height}`}">{{block.height}}</router-link>
        </p>
        <p>Mint: 
        <router-link :to="{path:`/block/${block.height}`}">{{block.witness}}</router-link>
        </p>
        <p>Txn: 
        <router-link :to="{path:`/txs/?b=${block.height}`}" style="text-align: left;">{{block.txn}} Transactions</router-link>
        </p>
        <p>TimeStamp: 
          {{block.age}} ({{block.utcTime}})
        </p>
      </li>
    </ul>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: "IndexBlock",

    computed: {
      ...mapState(['indexBlockList'])
    },

    mounted () {
      this.$store.dispatch('getIndexBlockList')
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .homeBlocks-box {
    width: 485px;
    border: 1px solid #F1F1F1;
    background-color: #FFFFFF;
    box-shadow: 0 10px 40px 0 rgba(0, 0, 0, .1);
    .my-list-banner {
      background: #F6F7F8;
      padding: 0 10px 0 10px;
      display: flex;
      justify-content: space-between;
      align-items: center;
      height: 54px;
      box-shadow: 0 2px 3px rgba(0,0,0,0.1);
      .my-list-title {
        font-size: 20px;
        line-height: 24px;
        color: #2C2E31;
        font-weight: bolder;
      }
      > a{
        font-size: 14px;
        line-height: 18px;
        color: #4b78aa;
      }
    }
    .my-list-body {
      padding-left: 0;
      padding-top: 8px;
      margin-bottom: 0;
      padding-bottom: 10px;

      li {
        list-style: none;
        text-align: left;
        padding: 10px 10px 10px 10px;
        line-height: 1;
        > p{
          /*font-size: 12px;*/
          font-size: 12px;
          line-height: 12px;
          color: #2C2E31;
          margin-bottom: 5px;
          padding-left: 0px;
          > img {
            width: 24px;
            height: 24px;
          }
          > a{
            padding-left: 5px;
            font-size: 14px;
            line-height: 18px;
            font-weight: bold;
            color: #4b78aa;
          }
          &:nth-of-type(1) {
            display: flex;
            padding-left: 0;
            align-items: center;
            padding-bottom: 5px;
          }
          &:nth-of-type(2) {
            margin-top: 5px;
          }
          &:nth-of-type(3) {
            margin-top: 5px;
          }
        }
        > a {
          display: inline-block;
          color: #4b78aa;
          padding-left: 50px;
          padding-top: 14px;
          /*font-size: 12px;*/
          font-size: 14px;
          line-height: 12px;

          &:first-of-type {
            width: 380px;
            text-overflow: ellipsis;
            overflow: hidden;
          }
          &:nth-of-type(2) {
            width: 380px;
            text-overflow: ellipsis;
            overflow: hidden;
          }
        }
        &:nth-child(2n) {
          background: #F6F7F8;
          > p {
            &:nth-of-type(1) {
              border-bottom: 1px solid #FFFFFF;
            }
          }
        }
        &:nth-child(2n+1) {
          background: #FFFFFF;
          > p {
            &:nth-of-type(1) {
              border-bottom: 1px solid #f6f7f8;
            }
          }
        }
      }
    }
  }
  @media screen and (max-width:480px) {
    .homeBlocks-box {
      width: 100%;
      margin-bottom: 15px;
      .my-list-body {
        li {
          > a {
            &:first-of-type {
              width: 340px;
            }
          }
        }
      }
    }
  }

  @media screen and (max-width:375px) {
    .homeBlocks-box {
      .my-list-body {
        li {
          > a {
            &:first-of-type {
              width: 300px;
            }
          }
        }
      }
    }
  }

</style>
