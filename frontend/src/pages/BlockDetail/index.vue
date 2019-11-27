<template>
  <div class="blockDetail-box">
    <div class="blockDetail-header">
      <div class="my-header-container">
        <h2>Block Information</h2>
      </div>
    </div>
    <div class="blockDetail-information">
      <div class="card">
        <div class="card-header">
          <h5>Block #{{blockHeight}}</h5>
        </div>
        <div class="card-body">
          <table class="table">
            <tr><th>Height: </th><td>{{blockHeight}}</td></tr>
            <tr><th>Hash: </th><td>{{blockDetail.blockHash}}</td></tr>
            <tr><th>TimeStamp: </th><td>{{blockDetail.age}} ({{blockDetail.utcTime}})</td></tr>
            <tr><th>Transactions: </th><td>
              <router-link :to="{path:`/txs/?b=${blockHeight}`}">{{blockDetail.txn}}</router-link></td></tr>
            <tr><th>Tx: </th><td>{{blockDetail.txList}}</td></tr>
            <tr><th>Witness: </th><td><router-link :to="{path:`/account/${blockDetail.witness}`}">{{blockDetail.witness}}</router-link></td></tr>
            <tr><th>Total Gas Limit: </th><td>{{blockDetail.totalGasLimit}}</td></tr>
            <tr><th>Average Gas Price: </th><td>{{blockDetail.avgGasPrice}}</td></tr>
            <tr><th>Parent Hash: </th><td>{{blockDetail.parentHash}}</td></tr>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: "Block",
    data() {
      return {
        blockHeight: '',
      }
    },
    methods: {
      fetchData(r) {
        this.blockHeight = r.params.id
        this.$store.dispatch('getBlockDetail', this.blockHeight)
      }
    },

    computed: {
      ...mapState(['blockDetail'])
    },

    watch: {
      '$route': function (r) {
        this.fetchData(r)
      }
    },
    mounted: function () {
      this.fetchData(this.$route)
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .blockDetail-box {
    padding-bottom: 160px;
    background: #F6F7F8;
    .blockDetail-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        width: 1000px;
        margin: 0 auto;
        text-align: left;
        overflow: hidden;
        height: 60px;
        > h2 {
          font-size: 28px;
          line-height: 30px;
          margin: 10px 0 10px;
          font-weight: bold;
        }
        > p {
          font-size: 14px;
          height: 18px;
          margin: 0;
        }
      }
    }
    .blockDetail-information {
      width: 1000px;
      margin: 24px auto 0;
      background: #FFFFFF;
      text-align: left;
      padding: 15px 50px 0 50px;
      box-shadow: 0 2px 3px rgba(0,0,0,0.1);
      > img {
        height: 50px;
      }
      h4 {
        font-size: 14px;
        line-height: 18px;
        font-weight: bolder;
        padding-bottom: 12px;
        margin: 0;
        color: #2c2e31;
        border-bottom: 1px solid #F1F1F1;
      }
      > div {
        margin-bottom: 60px;
        &:last-child {
          margin-bottom: 0;
        }
      }
      .blockDetail-height-txns {
        display: flex;
        margin-top: 35px;
        .blockDetail-height, .blockDetail-txns {
          width: 50%;
        }
      }
      .blockDetail-witness {
        padding-bottom: 60px;
        a {
          color: #4b78aa;
        }
      }

      p {
        /*font-size: 18px;*/
        font-size: 14px;
        line-height: 22px;
        color: #2C2E31;
        margin-top: 20px;
        margin-bottom: 0;
        a {
          /*font-size: 18px;*/
          font-size: 14px;
          line-height: 22px;
          color: #2C2E31;
        }
      }
    }
  }

  @media screen and (max-width:480px) {
    .blockDetail-box {
      padding-bottom: 24px;
      .blockDetail-header {
        .my-header-container {
          height: auto;
          width: 100%;
          flex-direction: column;
          padding: 0 25px;
          .my-pages {
            margin: 0;
          }
        }
      }
      .blockDetail-information {
        width: 100%;
        padding: 15px 25px 0 25px;
        p {
          text-overflow: ellipsis;
          overflow: hidden;
        }
      }
    }
  }
</style>
