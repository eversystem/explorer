<template>
  <div class="accountDetail-box">

    <div class="accountDetail-header">
      <div class="my-header-container">
        <h1>Contract</h1>
        <p>{{address}}</p>
      </div>
    </div>


    <div class="accountDetail-info">
      <div class="accountDetail-tips">
        <img src="/static/img/iostWhite.png" alt="">
      </div>

      <div class="my-tab-content">
        <div class="my-tab-pane2">
          <!--
          <div class="pane2-tips1">
            <h4>Publisher:</h4>
            <p></p>
          </div>
          <div class="pane2-tips2">
            <h4>Time:</h4>
            <p>{{ContractInfo.time}}</p>
          </div>
          -->
          <div class="pane2-tips3">
            <h4>Code:</h4>
            <pre>{{ContractInfo.code}}</pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import { config, axiosWithJWT } from '../../utils/config'
  const { apis } = config

  export default {
    name: "ContractDetail",
    data() {
      return {
        id: this.$route.params.id,
        // account: {},
        // txnList: {},
        // txnLen: '',
        ContractInfo: {},
        // currentTab: false,
        isShow: false,   // 当接口code返回1时，切换方法不执行
      }
    },

    computed: {
      ...mapState(['accountDetail', 'accountTxnInfo']),
      showContractCode () {
        // return this.address.slice(0,4) != 'IOST'
        return this.address.slice(0,8) == 'Contract'
      }
    },

    methods: {

      fetchData(r) {
        this.address = r.params.id

        //this.$store.dispatch('getAccountDetail', this.address)

        //this.$store.dispatch('getAccountTxnInfo', this.address)


        // axiosWithJWT()..get('https://127.0.0.1/api/account/' + this.address).then((response) => {
        //   this.account = response.data
        // })
        // axiosWithJWT()..get('https://127.0.0.1/api/account/' + this.address + '/txs').then((response) => {
        //   this.txnList = response.data.txn_list
        //   this.txnLen = response.data.txn_len
        // })

        // this.addressx = this.address.replace('Contract','')
        // axiosWithJWT()..get('https://127.0.0.1/api/contracts/' + this.address).then((response) => {
        const token = window.localStorage.getItem('ex_token');
        axiosWithJWT(token).get(`${apis.contract}${this.id}`).then((response) => {
          if (response.data.code == 1) {
            this.isShow = true
            return
          }
          this.ContractInfo = response.data.data
        })
      },
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
  .accountDetail-box {
    padding-bottom: 160px;
    background: #F6F7F8;
    .accountDetail-header {
      background: #F6F7F8;
      box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
      .my-header-container {
        width: 1000px;
        height: 120px;
        margin: 0 auto;
        text-align: left;
        overflow: hidden;
        > h1 {
          margin: 21px 0 15px;
          color: #273238;
          font-size: 36px;
          line-height: 44px;
          font-weight: bold;

        }
        > p {
          margin: 0;
          font-size: 14px;
          line-height: 18px;
        }
      }
    }


    .accountDetail-info {
      width: 1000px;
      margin: 24px auto 0;
      padding-bottom: 15px;
      background: #FFFFFF;
      text-align: left;
      .accountDetail-tips {
        > img {
          height: 50px;
        }
        padding: 15px 100px 60px 50px;
        box-shadow: 0 2px 3px 0 rgba(0, 0, 0, .1);
        .accountDetail-height-value {
          display: flex;
          .accountDetail-height, .accountDetail-value {
            width: 50%;
          }
        }
        h4 {
          font-size: 14px;
          line-height: 18px;
          font-weight: bold;
          color: #2C2E31;
          padding-bottom: 12px;
          padding-top: 20px;
          border-bottom: 1px solid #F1F1F1;
        }

        p {
          /*font-size: 18px;*/
          font-size: 14px;
          line-height: 22px;
          color: #2C2E31;
          margin: 0;
          a {
            /*font-size: 18px;*/
            font-size: 14px;
            line-height: 22px;
            color: #2C2E31;
          }
        }
      }


      .my-tab {
        display: flex;
        margin: 0;
        padding: 0;
        height: 60px;
        font-size: 14px;
        color: #2C2E31;
        font-weight: bold;
        line-height: 60px;
        > li {
          padding-left: 50px;
          list-style: none;
          width: 50%;
          cursor: pointer;
          &.active {
            background: #F6F7F8;
          }
        }
      }

      .my-tab-content {
        .my-tab-pane1 {
          .pane1-header {
            padding: 0 50px;
            .pane1-title {
              /*font-size: 18px;*/
              font-size: 14px;
              line-height: 22px;
              color: #2C2E31;
              padding-top: 20px;
            }
            .pane1-content {
              display: flex;
              justify-content: space-between;
              height: 50px;
              /*font-size: 18px;*/
              font-size: 14px;
              line-height: 22px;
              color: #2C2E31;
              margin-top: 60px;
            }
            p {
              margin-bottom: 0;
              a {
                color: #4b78aa;
              }
            }
          }

          .accountDetail-list {
            .my-list-header {
              /*font-size: 16px;*/
              font-size: 14px;
              font-weight: 500;
              border-bottom: 1px solid #f6f7f8;
              border-top: 1px solid #f6f7f8;
            }
            .list-wrap {
              padding-bottom: 60px;
              .list-body-wrap {
                width: 100%;
                &:last-child {
                  .my-list-body {
                    border: 0;
                  }
                }
                &:hover {
                  box-shadow: 0 8px 30px 0 rgba(0, 0, 0, .15);
                }
              }
            }

            ul {
              display: flex;
              padding: 17px 0 17px 20px;
              width: 1000px;
              margin: 0 auto;
              border-bottom: 1px solid #f6f7f8;
              li {
                list-style: none;
                text-align: left;
                color: #2c2e31;
                a {
                  color: #4b78aa;
                }
                &:first-child {
                  width: 100px;
                }
                &:nth-child(2) {
                  width: 200px;
                  padding-right: 40px;
                  text-overflow: ellipsis;
                  overflow: hidden;
                }

                &:nth-child(3) {
                  width: 200px;
                  padding-right: 40px;
                  text-overflow: ellipsis;
                  overflow: hidden;
                }
                &:nth-child(4) {
                  width: 200px;
                  padding-right: 40px;
                  text-overflow: ellipsis;
                  overflow: hidden;
                }
                &:nth-child(5) {
                  width: 100px;
                }
                &:nth-child(6) {
                  width: 100px;
                }
                &:nth-child(7) {
                  width: 80px;
                }
              }
            }
          }
        }

        .my-tab-pane2 {
          padding: 0 50px;
          > div {
            margin-top: 60px;
            h4 {
              font-size: 14px;
              line-height: 18px;
              font-weight: bold;
              color: #2C2E31;
              padding-bottom: 12px;
              border-bottom: 1px solid #F1F1F1;
              margin: 0;
              &:last-child {
                padding-bottom: 60px;
              }
            }

            p {
              font-size: 14px;
              line-height: 22px;
              color: #2C2E31;
              margin: 20px 0 0;
            }
          }
        }
      }
    }
  }

  @media screen and (max-width:480px) {
    .accountDetail-box {
      padding-bottom: 24px;
      .accountDetail-header {
        .my-header-container {
          height: auto;
          width: 100%;
          flex-direction: column;
          padding: 0 25px;
          > p {
            text-overflow: ellipsis;
            overflow: hidden;
          }
          .my-pages {
            margin: 0;
          }
        }
      }

      .accountDetail-info {
        width: 100%;
        .accountDetail-tips {
          padding: 15px 25px 60px;
        }

        .my-tab {
          > li {
            padding-left: 25px;
          }
        }

        .my-tab-content {
          .my-tab-pane1 {
            .pane1-header {
              padding: 0 25px;
              .pane1-content {
                margin: 30px 0;
                > p {
                  &:first-child {
                    width: 254px;
                  }
                }
              }
            }

            .accountDetail-list {
              overflow: auto;
            }
          }

          .my-tab-pane2 {
            padding: 0 25px;
          }
        }
      }
    }
  }
</style>
