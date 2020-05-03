<template>
    <div id="chart">
        <apexchart type="heatmap" height="350" :options="chartOptions" :series="series"></apexchart>
    </div>
</template>

<script>
import VueApexCharts from 'vue-apexcharts'
import Vue from 'vue'
Vue.component('apexchart', VueApexCharts)

export default {
    components: {
        apexchart: VueApexCharts,
    },
    data() {
        return {
           series: [{
              name: 'Jan',
              data: [10,20,30,40,50,60,70]
            }
          ],
          chartOptions: {
            chart: {
              height: 350,
              type: 'heatmap',
            },
            plotOptions: {
              heatmap: {
                shadeIntensity: 0.5,
                radius: 0,
                useFillColorAsStroke: true,
                colorScale: {
                  ranges: [{
                      from: -30,
                      to: 5,
                      name: 'low',
                      color: '#00A100'
                    },
                    {
                      from: 6,
                      to: 20,
                      name: 'medium',
                      color: '#128FD9'
                    },
                    {
                      from: 21,
                      to: 45,
                      name: 'high',
                      color: '#FFB200'
                    },
                    {
                      from: 46,
                      to: 55,
                      name: 'extreme',
                      color: '#FF0000'
                    }
                  ]
                }
              }
            },
            dataLabels: {
              enabled: false
            },
            stroke: {
              width: 1
            },
            title: {
              text: 'HeatMap Chart with Color Range'
            },
          },
        }      
    }
}
</script>