<template>
  <div id="app">
    <Domains msg="Welcome to the Domains Services App" />
  </div>
</template>

<script>
import Domains from "./components/Domains.vue";

export default {
  name: "App",
  components: {
    Domains
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
