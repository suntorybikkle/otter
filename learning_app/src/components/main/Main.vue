<template>
    <div id="main">
        <MainTimer :status="toggle" @push="pushTime"></MainTimer>
        <MainButton :name="toggle" @toggle="toggleTimer" @reset="resetTimer"></MainButton>
        <hr>
        <hr>
        <span style="font-size: 20px">↓HTTP リクエストボディイメージ(後で消す)</span>
        <pre id="tmp">{{httpJson}}</pre>
    </div>
</template>

<script>
import MainButton from '../parts/MainButton';
import MainTimer from '../parts/MainTimer';
export default {
    components:{
        MainButton,
        MainTimer
    },
    data() {
        return {
            toggle:"start",
            httpJson: {
                userId:"123",
                userName:"ktguy",
                studyInfo:{
                    studyId:"321",
                    subId:"0",
                    color:"black",
                    studyTime:"00:00:00"
                }
            }
        }
    },
    methods: {
        toggleTimer(timerStatus) {
            this.toggle = timerStatus;
        },
        pushTime(stopTime) {
            console.log(stopTime);
            this.httpJson.studyInfo.studyTime = stopTime;
        },
        resetTimer() {
            console.log("test");
        }
    }   
}
</script>

<style scoped>
#tmp {
    border: dotted 1px black;
    text-align:left;
    font-size: 15px;
}

#main{
    margin:50px;
}

</style>