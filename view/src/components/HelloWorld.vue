<template>
  <div class="mt-5">
    <link
      rel="stylesheet"
      href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css"
      integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh"
      crossorigin="anonymous"
    />
    <h1>{{msg}}</h1>
    <div class="d-flex justify-content-center">
      <form>
        <label for="inputDomain">Enter a domain:</label>
        <input
          type="text"
          v-model="searchedDomain"
          class="form-control"
          placeholder="domain.com"
          aria-describedby="emailHelp"
        />
        <br />
        <div class="row">
          <div class="col-sm-6">
            <button type="submit" v-on:click="searchDomain" class="btn btn-primary">Search domain</button>
          </div>
          <div class="col-sm-6">
            <button
              type="submit"
              v-on:click="getDomains"
              class="btn btn-primary"
            >Consult Searched Domains</button>
          </div>
        </div>
      </form>
    </div>
    <br />
    <div class="d-flex justify-content-center">
      <div class="col">
        <div class="col-sm-12">
          <div class="border border-dark-rounded">
            <p class="title-1">Consulted Domains</p>
            <ul>
              <p v-for="domain in domains" :key="domain.host">{{ domain.host }}</p>
            </ul>
          </div>
        </div>
        <div class="col-sm-16">
          <div class>
            <br />
            <div class="border border-dark-rounded">
              <h5>{{searchedDomain}}</h5>
              <p class="title-1">Servers:</p>
              <p v-for="server in servers" :key="server.address">
                IP address: {{ server.address }}
                <br />
                SSL Grade: {{server.ssl_grade}}
                <br />
                Country: {{server.country}}
                <br />
                Owner: {{server.owner}}
                <br />
              </p>
              <br />
              <p class="title-1">Domain:
              <p class="text">
              Servers Changed: {{serverChange}}
              <br />
              SSL Grade: {{sslGrade}}
              <br />
              SSL Previous Grade: {{previousGrade}}
              <br />
              Logo: {{logo}}
              <br />
              Title: {{title}}
              <br />
              Is down: {{isDown}}
              <br />
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  data() {
    return {
      searchedDomain: "",
      sslGrade: null,
      previousGrade: null,
      serverChange: null,
      title: null,
      logo: null,
      isDown: null,
      domains: null,
      servers: null
    };
  },
  methods: {
    getDomains: function(e) {
      e.preventDefault();
      axios.get("http://localhost:8000/domains").then(
        response => {
          console.log(response.data);
          this.domains = response.data.items;
        },
        error => {
          console.log(error);
        }
      );
    },
    searchDomain: function(e) {
      e.preventDefault();
      axios.get("http://localhost:8000/domains/" + this.searchedDomain).then(
        response => {
          console.log(response.data);
          this.servers = response.data.servers;
          this.serverChange = response.data.servers_changed;
          this.sslGrade = response.data.ssl_grade;
          this.previousGrade = response.data.previous_ssl_grade;
          this.logo = response.data.logo;
          this.title = response.data.title;
          this.isDown = response.data.is_down;
        },
        error => {
          console.log(error);
        }
      );
    }
  }
};
</script>


<style>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
text{
  color: black;
}
p.title-1{
  color:cornflowerblue;
  font-family: Impact, Charcoal, sans-serif;
  font-size: 30px;
}
</style>