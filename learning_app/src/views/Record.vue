<template>
  <div class="record">
    <div id="nav">
      <router-link to='/record/liner'>liner</router-link> |
      <router-link to='/record/table'>table</router-link>
    </div>
    <transition name="fade" mode="out-in">
      <router-view :data="data" :cate="cate" :json="json"></router-view>
    </transition>
  </div>
</template>

<script>
export default {
  data(){
    return{
      json:{},
      cate:[],
      data:[]
    }
  },
  methods:{
    get() {
      let body = this;
      fetch('http://127.0.0.1:8080/record/',{
        mode: 'cors'
      })
      .then(res => res.json())
      // データをjsonに格納
      .then(data => {
        body.json = data.studyInfos;
      })
      // 
      .then(() => {
        const newCate = body.json.map((cateVal, index) => {
          const date = body.json[index].dateTime.split(" ");
          return date[0];
        });

        const newData = body.json.map((dataVal, index) => {
          return body.json[index].studyTime;
        });

        body.cate = newCate;
        body.data = newData;

        console.log(body.cate, body.data)
      });
    },
    substituteFields(json){
      console.log(json);
      const newCate = this.cate.map((cateVal, index) => {
        const date = this.json[index].dateTime.spirit(" ");
        
        return date[0];
      })
      console.log(newCate);
    }
  },
  created(){
    this.get();
  }
}

</script>

<style lang="scss" scoped>
  
.fade-enter-active, .fade-leave-active {
  transition: opacity .3s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>
