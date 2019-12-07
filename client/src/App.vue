<template>
  <div id="app">
    <h2>CRUD</h2>
    <h3>Read all</h3>
    <button @click="getevents">"GET" /events</button>
    <h3>Read one</h3>
    <input type="text" name="idRead" placeholder="id" v-model="idRead" required />
    <button @click="geteventsId">"GET" /events/{id}</button>
    <h3>Create one</h3>
    <input type="text" name="idCreate" placeholder="id" v-model="idCreate" required />
    <input type="text" name="titleCreate" placeholder="title" v-model="titleCreate" required />
    <input
      type="text"
      name="descriptionCreate"
      placeholder="description"
      v-model="descriptionCreate"
      required
    />
    <button @click="addevent">"POST" /events</button>
    <h3>Update one</h3>

    <input type="text" name="idUpdate" placeholder="id" v-model="idUpdate" required />
    <input type="text" name="titleUpdate" placeholder="title" v-model="titleUpdate" required />
    <input
      type="text"
      name="descriptionUpdate"
      placeholder="description"
      v-model="descriptionUpdate"
      required
    />
    <button @click="updateevent">"PATCH" /events</button>
    <h3>Delete one</h3>
    <input type="text" name="idDelete" placeholder="id" v-model="idDelete" required />
    <button @click="removeAtricleId">"DELETE" /events</button>
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
      events: [],
      api_url: "http://localhost:8090",
      // api_url: "https://api-dot-vue-go-gcp-20191111.appspot.com",
      idRead: "",
      idCreate: "",
      titleCreate: "",
      descriptionCreate: "",
      idDelete: "",
      idUpdate: "",
      titleUpdate: "",
      descriptionUpdate: ""
    };
  },
  methods: {
    getevents: async function() {
      let res = await axios.get(this.api_url + "/events");
      console.log(res.data);
    },
    geteventsId: async function() {
      let res = await axios.get(this.api_url + "/events/" + this.idRead);
      console.log(res.data);
    },
    addevent: async function() {
      let res = await axios.post(this.api_url + "/event", {
        id: this.idCreate,
        title: this.titleCreate,
        description: this.descriptionCreate
      });
      console.log(res);
    },
    updateevent: async function() {
      let res = await axios.patch(this.api_url + "/events/" + this.idUpdate, {
        id: this.idUpdate,
        title: this.titleUpdate,
        description: this.descriptionUpdate
      });
      console.log(res);
    },
    removeAtricleId: async function() {
      let res = await axios.delete(this.api_url + "/events/" + this.idDelete);
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
