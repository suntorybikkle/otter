<template>
    <div>
        <div class="time">{{ formatTime }}</div>
    </div>
</template>

<script>
export default {
    props:['status'],
    data() {
        return {
            hour:0,
            min: 0,
            sec: 0,
            timerObj: null
        }
    },
    watch:{
        status() {
            // stop処理
            if(this.status === 'start') {
                this.complete();
            }
            // start処理
            else {
                let self = this;
                this.timerObj = setInterval(function() {
                    self.count();
                }, 1000);
            }
        }
    },
    methods: {
        count() {
            this.sec++;

            // 1分経過後
            if(this.sec >= 60) {
                this.min++;
                this.sec = 0;
            }

            // 1時間経過後
            if(this.min >= 60) {
                this.hour++;
                this.min = 0;
            }
        },
        // 停止処理
        complete() {
            clearInterval(this.timerObj);
            this.$emit('push',this.formatTime);
        }
    },
    computed: {
        formatTime() {
            let timeStrings = [
                this.hour.toString(),
                this.min.toString(),
                this.sec.toString()
            ].map(function(str) {
                if(str.length < 2) {
                    return "0" + str;
                }else {
                    return str;
                }
            })
            return timeStrings[0] + ':' + timeStrings[1] + ":" + timeStrings[2];
        }
    }
}
</script>