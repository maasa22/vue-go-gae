<template>
  <div id="app">
    <button @click="get_todos">api button</button>
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
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
      todos: []
    };
  },
  methods: {
    get_todos: async function() {
      // let res = await axios.get("https://jsonplaceholder.typicode.com/todos/1"); // dummy api on net
      let res = await axios.get("http://localhost:8090/api/todos"); // local go server
      // let res = await axios.get(
      //   "https://api-dot-vue-go-gcp-20191111.appspot.com"
      // ); // server on gae
      this.todos = res.data;
      console.log(this.todos);
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
