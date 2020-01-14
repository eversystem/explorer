<template>
  <div class="login-wrapper">
    <main class="login-outer">
      <section class="login-container">
        <div class="login-header">
          <div class="explorer-title">
            <h2 >
              ATSchain Explorer
            </h2>
          </div>
        </div>
        <div class="message-container">
          <h3 id="login-title">
            ログイン
          </h3>
          <p id="login-message" >
            ATSchain Explorerに
            <br>
            ログインしてください。
          </p>
        </div>
        <div class="input-container">
          <input class="login-input username" placeholder="ユーザー名" v-model.trim="username"/>
          <input type="password" class="login-input password" placeholder="パスワード" v-model.trim="password"/>
        </div>
        <div class="button-container">
          <button type="button" class="login-button"  @click="login" :class="{active: isOK}">
            ログインする
          </button>
        </div>

      </section>
    </main>
  </div>
</template>

<script>
  import axios from 'axios';
  import { config} from '../../utils/config';

  export default {
    name: "Login",
    data() {
      return {
        username: "",
        password: ""
      }
    },
    computed: {
      isOK () {
        if (this.username != '' && this.password != '') {
          return true
        } else {
          return false
        }
      }
    },
    methods: {
      login(){
        console.log(this.username, this.password);
        if(!this.isOK) {
          return alert("ユーザー名とパスワードを入力してください")
        }
        axios({
          method: 'post',
          url: config.apis.login, 
          auth: {
            username: this.username,
            password: this.password
          }
          })
          .then( res => {
            if( res.data && res.data.jwt){
              localStorage.setItem('ex_token', res.data.jwt)
            }
            window.location.href="/"
          })
          .catch(err => {
            alert("ログインに失敗しました")
            window.location.reload()
          })
      }
    }
  }
</script>

<style lang="less" rel="stylesheet/less">
  .login-wrapper {
    width: 100%;
    height: 80vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-content: center;
  }

  .login-container {
    width: 400px;
    height: 500px;
    position: relative;
    font-size: 14px;
    margin: auto;
    border-radius: 5px;
    white-space: normal;
    background-color: #fff;
    box-shadow: 1px 5px 10px rgba(0,0,0,.12);
    color: #2d333a;
  }

  .login-header {
    width: 400px;
    height: 100px;
    background: #2149a7;
    border-top-left-radius: 5px;
    border-top-right-radius: 5px;
    color: white;
    display: flex;
    justify-content: center;
    align-content: center;
  }

  .explorer-title {
    margin: auto;

    h2 {
      font-size: 28px;
    }
  }

  .message-container {
    margin: 25px 20px 50px 20px;
  }

  #login-message {
    font-size: 14px;
  }

  #login-title {
    font-size: 24px;
    margin-bottom: 20px;
  }

  .input-container {
    width: 100%;
    height: 100px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-content: center;

    .login-input {
      width: 300px;
      height: 50px;
      margin: auto;
      border: 1px solid #c2c8d0;
      font-size: 15px;
      outline: none;
      padding: 0 15px 0 15px;
    }

    .username {
      border-radius: 5px 5px 0px 0px;
      box-shadow: 0;
    }

    .password {
      border-radius:  0px 0px 5px 5px;
    }
  }
  .login-outer {
    display: flex;
    align-items: center;
    overflow: auto;
    width: 100%;
    height: 100%;
    padding: 80px 0;
    padding: var(--outer-padding) 0;
  }

  .button-container {
    width: 100%;
    display: flex;
    justify-content: center;
    align-content: center;
    margin-top: 50px;

    .login-button {
      width: 300px;
      height: 40px;
      border-radius: 5px;
      background: #2149a7;
      color: white;
      font-size: 17px;
    }
  }

</style>