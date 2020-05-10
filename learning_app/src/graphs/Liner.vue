<template>
    <div>
      <apexchart  type="line" :options="options" :series="series"></apexchart>
    </div>
</template>

<script>
import Vue from 'vue'
import VueApexCharts from 'vue-apexcharts'

Vue.component('apexchart', VueApexCharts)

export default {
    props:["data", "cate"],  
    data: function() {
        return {
            graphData:{
                x:[], // categories
                y:[], // data
            },
            options: {
                chart: {
                    id: 'vuechart-example'
                },
                xaxis: {
                    categories: []
                }
            },
            series: [{
                name: 'series-1',
                data: []
            }]
        }
    },
    created(){
        this.getData();
    },
    methods:{
        // データ取得
        getData(){
          const newData = this.data.map((now,index)=>{
              return this.data[index];
          })
          const newCate = this.cate.map((now,index)=> {
               return this.cate[index];
          })
          console.log(newData, newCate)

            this.graphData.x = newCate;
            this.graphData.y = newData;
            this.parseData(this.graphData);

            this.series = [{
                name:'test',
                data: this.graphData.y
            }]
            this.options= {
                xaxis:{
                    categories: this.graphData.x
                }
            }
        },

        // 同一日付のデータをマージする処理
        parseData(preGraphData) {
            for(let i = 0; i < preGraphData.x.length; i++) {
                if(preGraphData.x[i] === preGraphData.x[i + 1] && i <preGraphData.x.length){
                    preGraphData.y[i + 1] = preGraphData.y[i] + preGraphData.y[i + 1];
                    preGraphData.x.splice(i, 1);
                    preGraphData.y.splice(i, 1);
                    i--;
                }
            }
        }
    }
}
</script>