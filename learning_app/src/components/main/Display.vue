<template>
  <div id="display">
    <Sidebar id="sidebar"></Sidebar>
    <div id="main-area">
      <Top></Top>
      <Main id="main"></Main>
    </div>
  </div>
</template>

<script>
import Top from './Top.vue';
import Sidebar from './Sidebar.vue';
import Main from './Main.vue';

export default {
  data() {
    return {
      result:0
    }
  },
  components: {
    Top,
    Sidebar,
    Main
  }
}
</script>

<style>
body {
  background-color: rgba(255, 255, 255, .5);
}
#main-area {
  flex:2;
  display:flex;
  flex-direction: column;
  justify-content: center;
}
#main{
  flex:1;
  font-size:5rem;
}
#sidebar{
  flex:1;
  max-width:300px;
}
#display {
  display: flex;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>