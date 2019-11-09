<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png" /> -->
    <button @click="get_api">get api</button>
    <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
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
  methods: {
    
    get_api: async function() {
      // let res = await axios.get("/api/hoge", {
      //   // mode: "cors"
      // });
      let res = await axios.get(
        "https://test-gcp-ma-20191105.appspot.com/api/hoge",
        // "http://localhost:8080/api/hoge",
        {
          // mode: "cors"
        }
      );

      // let res = await axios.get("https://jsonplaceholder.typicode.com/todos/1");
      console.log(res);
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
