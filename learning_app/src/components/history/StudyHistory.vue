<template>
    <div>
        <h1>Study History</h1>
        <div>
            <apexchart type="line" height="500" :options="chartOptions" :series="series"></apexchart>
        </div>
        <hr>
        今週の学習時間
        <div>
            {{series[0].data}}
        </div>
        <hr>
        <ColorRange></ColorRange>
        <router-link class="menu-button" to="/">back</router-link>
    </div>
</template>

<script>
import ColorRange from './ColorRange.vue';
import VueApexCharts from 'vue-apexcharts';
import Vue from 'vue'
Vue.component('apexchart', VueApexCharts)

export default {
    components: {
        apexchart: VueApexCharts,
        ColorRange
    },
    data(){
        return {
            series: [{
                name: "Desktops",
                data: [60, 40, 35, 50, 45, 62, 120]
            }],
            chartOptions: {
                chart: {
                    height: 350,
                    width:60,
                    type: 'line',
                    zoom: {
                        enabled: false
                    }
                },
                dataLabels: {
                    enabled: false
                },
                stroke: {
                    curve: 'straight'
                },
                title: {
                    text: '今週の学習時間',
                    align: 'center'
                },
                grid: {
                    row: {
                        colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
                        opacity: 0.9
                    },
                },
                xaxis: {
                    categories: ['Sum', 'Mon', 'Tue', 'Wed', 'Thr', 'Fly', 'Sar'],
                }
            },
        }
    }        
}
    
    // data() {
    //     return {
    //         payload:[
    //             {
    //                 subject:"default",
    //                 color:"black",
    //                 time:30
    //             },
    //             {
    //                 subject:"math",
    //                 color:"red",
    //                 time:30
    //             },
    //             {
    //                 subject:"blue",
    //                 color:"blue",
    //                 time:30
    //             },
    //             {
    //                 subject:"green",
    //                 color:"green",
    //                 time:30
    //             },
    //             {
    //                 subject:"yellow",
    //                 color:"yellow",
    //                 time:30
    //             }
    //         ]
    //     }
    // }

</script>

<style scoped>
.box{
    max-width:100%;
    min-height:85vh;
    background-color:#fcc;

}
</style>