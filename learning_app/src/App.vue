<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/record/table">Record</router-link> |
      <router-link to="/login">Login</router-link>
    </div>
    <div id="main" class="main">
      <div class="main__nav">
        <Sidebar></Sidebar>
      </div>
      <div class="main__main">
        <transition name="fade" mode="out-in">
          <router-view/>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
import Sidebar from './components/Sidebar.vue'
export default {
  components:{
    Sidebar
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #f66;
    }
  }
}

.main {
  max-width:1100px;
  margin-right:auto;
  margin-left:auto;
  border:#666 2px solid;
  border-radius: 1rem;
  display:flex;
  justify-content:center;
  .main__nav{
    flex:1;
    width:300px;
    border-right:#2c3e50 1px solid;
  }
  .main__main{
    flex:3;
  }
}

@media screen and (max-width:767px){
  .main{
    flex-direction: column;
    .main__nav{
      width:100%;
      order:1;
      border-top:#2c3e50 1px solid;
    }
  }
  #app{
    min-width: 300px;
  }
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .3s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
