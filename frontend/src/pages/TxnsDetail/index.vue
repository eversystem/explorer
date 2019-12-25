<template>
  <div class="txnsDetail-box">
    <div class="txnsDetail-header">
      <div class="my-header-container">
        <h2>Transaction Information</h2>
      </div>
    </div>

    <div class="txnsDetail-information">
      <div class="card">
        <div class="card-header">
          <h5>Transaction #{{txHash}}</h5>
        </div>
        <div class="card-body">
          <table class="table">
            <tr>
              <th>Block Height:</th>
              <td>
                <router-link
                  :to="{path:`/block/${txnDetail.blockHeight}`}"
                >{{txnDetail.blockHeight}}</router-link>
              </td>
            </tr>
            <tr>
              <th>Publisher:</th>
              <td><a href="javascript:void(0)">{{txnDetail.publisher}}</a></td>
            </tr>
            <tr>
              <th>TimeStamp:</th>
              <td>{{txnDetail.age}}({{txnDetail.utcTime}})</td>
            </tr>
            <tr>
              <th>Signers:</th>
              <td>{{txnDetail.signers}}</td>
            </tr>
            <tr>
              <th>Gas Limit:</th>
              <td>{{txnDetail.gasLimit}}</td>
            </tr>
            <tr>
              <th>ReferredTx:</th>
              <td>{{txnDetail.referredTx}}</td>
            </tr>
            <tr>
              <th>AmountLimit:</th>
              <td>{{txnDetail.amountLimit}}</td>
            </tr>
          </table>
        </div>
      </div>

      <div class="card">
        <div class="card-header">
          <h5>Receipt:</h5>
        </div>
        <div class="card-body">
          <table class="table">
            <tr>
              <th>Tx Hash:</th>
              <td>{{txnDetail.receipt.tx_hash}}</td>
            </tr>
            <tr>
              <th>Gas Usage:</th>
              <td>{{txnDetail.receipt.gas_usage}}</td>
            </tr>
            <tr>
              <th>RamUsage:</th>
              <td>{{txnDetail.receipt.ram_usage}}</td>
            </tr>
            <tr>
              <th>Returns:</th>
              <td>{{txnDetail.receipt.returns}}</td>
            </tr>
            <tr>
              <th>Status:</th>
              <td>{{txnDetail.receipt.status_code}}</td>
            </tr>
            <tr>
              <th>Message:</th>
              <td>{{txnDetail.receipt.message}}</td>
            </tr>
          </table>
          <table class="table">
            <th>
              FuncName:
            </th>
            <th>
              Content:
            </th>
            <tr v-for="receipt in txnDetail.receipt.receipts">
              <td>
                {{receipt.func_name}}
              </td>
              <td>
                {{receipt.content}}
              </td>
            </tr>
          </table>
        </div>
        <div class="card">
          <div class="card-header">
            <h5>Actions:</h5>
          </div>
          <div class="card-body">
            <table class="table">
              <th>
                Contract:
              </th>
              <th>
                Function:
              </th>
              <th>
                Data:
              </th>
              <tr v-for="action in txnDetail.actions">
                <td>
                  <router-link :to="{path:`/contract/${action.contract}`}">{{action.contract}}</router-link>
                </td>
                <td>
                  {{action.action_name}}
                </td>
                <td>
                  {{action.data}}
                </td>
              </tr>
            </table>
          </div>
        </div>        
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "Tx",
  data() {
    return {
      txHash: ""
    };
  },
  methods: {
    fetchData(r) {
      this.txHash = r.params.id;

      this.$store.dispatch("getTxnDetail", this.txHash);
    }
  },

  computed: {
    ...mapState(["txnDetail"])
  },

  watch: {
    $route: function(r) {
      this.fetchData(r);
    }
  },
  mounted: function() {
    this.fetchData(this.$route);
  }
};
</script>


<style lang="less" rel="stylesheet/less">
.txnsDetail-box {
  padding-bottom: 160px;
  background: #f6f7f8;
  .txnsDetail-header {
    background: #f6f7f8;
    box-shadow: 0 2px 3px 0 rgba(0, 0, 0, 0.1);
    .my-header-container {
      width: 1000px;
      margin: 0 auto;
      height: 60px;
      text-align: left;
      overflow: hidden;
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
  .txnsDetail-information {
    width: 1000px;
    margin: 24px auto 0;
    text-align: left;
    background: #ffffff;
    padding: 15px 50px;
    position: relative;
    box-shadow: 0 2px 3px rgba(0, 0, 0, 0.1);
    h4 {
      font-size: 14px;
      line-height: 18px;
      font-weight: bold;
      padding-bottom: 12px;
      margin: 0;
      color: #2c2e31;
      &:last-child {
        padding-bottom: 60px;
      }
    }
    pre {
      margin-bottom: 0;
    }
    p {
      /*font-size: 18px;*/
      font-size: 14px;
      line-height: 22px;
      margin-top: 20px;
      margin-bottom: 0;
      font-weight: 300;
      a {
        color: #4b78aa;
      }
    }

    > div {
      margin-bottom: 60px;
      &:last-child {
        margin-bottom: 0;
      }
    }
    > img {
      height: 50px;
    }
    .txnsDetail-hash {
      margin-top: 35px;
    }
    .txnsDetail-to {
      .my-pre {
        margin-top: 20px;
        > span {
          display: inline-block;
          margin-bottom: 12px;
        }
      }
      .my-pre-normal {
        display: flex;
        align-items: center;
        margin-top: 20px;
        > pre {
          display: inline-block;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          max-width: 850px;
        }
      }
      span {
        font-weight: bold;
      }
    }

    .txnsDetail-hash,
    .txnsDetail-from,
    .txnsDetail-to {
      > h4 {
        border-bottom: 1px solid #f6f7f8;
      }
    }
    .txnsDetail-block-time {
      display: flex;
      .txnsDetail-block {
        width: 33.3%;
      }
      > div {
        > h4 {
          border-bottom: 1px solid #f6f7f8;
        }
      }
    }
    .txnsDetail-value-gas {
      display: flex;
      > div {
        width: 33.3%;
        > h4 {
          border-bottom: 1px solid #f6f7f8;
        }
      }
    }
  }
}

@media screen and (max-width: 480px) {
  .txnsDetail-box {
    padding-bottom: 24px;
    .txnsDetail-header {
      .my-header-container {
        height: auto;
        width: 100%;
        flex-direction: column;
        padding: 0 25px;
        > h1 {
          font-size: 30px;
        }
        .my-pages {
          margin: 0;
        }
      }
    }
    .txnsDetail-information {
      width: 100%;
      padding: 15px 25px 0 25px;
      p {
        text-overflow: ellipsis;
        overflow: hidden;
      }
      .txnsDetail-block-time {
        .txnsDetail-block {
          width: 50%;
        }
      }
    }
  }
}
</style>