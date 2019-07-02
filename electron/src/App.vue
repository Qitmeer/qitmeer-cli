<template>
  <div id="app">
    <div id="sidebar">
      <div id="logo"></div>
      <div id="menu">
        <h2>用户</h2>

        <a href="#">账户管理</a>

        <a href="#">地址管理</a>

        <a href="#">发送交易</a>

        <a href="#">交易记录</a>

        <a href="#">备份/恢复</a>

        <h2>数据</h2>
        <a href="#">块数据查询</a>
      </div>
      <div id="info">
        <h2>
          节点: 本地
          <a href="#">切换</a>
        </h2>
        <p>网络：Mainnet</p>
        <p>全网：100000</p>
        <p>节点：1000</p>
      </div>
      <div class="version">
        <p>版本：1.0.1 2019</p>
      </div>
    </div>
    <div class="mainwarp">
      <div class="main">
        <div class="header"></div>
        <div class="bd">{{ info }}</div>
      </div>
    </div>
  </div>
</template>

<script>
// import HelloWorld from "./components/HelloWorld.vue";
import axios from "axios";

export default {
  name: "app",
  components: {
    // HelloWorld
  },
  data() {
    return {
      info: null
    };
  },
  mounted() {
    axios({
      auth: {
        username: "admin",
        password: "123"
      },
      method: "post",
      url: "/api",
      data: {
        id: (Math.random() * 1000).toString(),
        method: "makeEntropy",
        params: null
      }
    }).then(response => (this.info = response.data));
  }
};
</script>

<style>
#app {
  display: flex;
  flex-direction: row;
  align-items: stretch;
  box-shadow: 0 1px 1px 0 rgba(0, 0, 0, 0.06), 0 2px 5px 0 rgba(0, 0, 0, 0.2);
  border-radius: 3px;

  width: 1000px;
  margin: 20px 0px;
  background-color: #fff;
}

#sidebar {
  width: 210px;

  background-color: #21252c;
  color: #fff;
  font-size: 14px;

  border-radius: 3px 0px 0px 3px;

  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: stretch;
}

#sidebar p {
  padding: 0px 20px;
}

#logo {
  height: 60px;
  border-bottom: 1px solid #404141;
  /* e5e7e9; */
}

#menu {
  flex-grow: 1;
}

#menu a {
  margin-left: 20px;
  color: #fff;
  font-size: 14px;
  line-height: 50px;
  text-decoration: none;
  display: block;
}

#sidebar h2 {
  font-size: 14px;
  color: #c3c3c3;
  line-height: 36px;
  margin: 0px 20px;
  border-bottom: 1px solid #404141;
}

#info p {
}

.version {
  padding-top: 20px;
}
.version p {
  color: #ccc;
  font-size: 12px;
  text-align: center;
}

.mainwarp {
  background-color: #f1f1f2;
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  justify-content: flex-start;
  align-items: stretch;
}
.header {
  height: 60px;
  background-color: #fff;
}

@media screen and (max-width: 1000px) {
  #app {
    width: 100%;
    height: 100%;
    margin: 0px;
  }
}
</style>
