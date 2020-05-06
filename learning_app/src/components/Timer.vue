<template>
    <div class="timer">
        <div class="timer__display">
            {{formatTime}}
        </div>
        <div class="timer__toggle" @click="toggleTime">
            {{status}}
        </div>
        <div class="timer__reset" @click="resetTime">
            RESET
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            status:'START',
            hour:0,
            min: 0,
            sec: 0,
            timerObj: null
        }
    },
    watch:{
        status() {
            // stop処理
            if(this.status === 'START') {
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
        },
        toggleTime() {
            if(this.status === "START") {
                this.status = "STOP";
            }else {
                this.status = "START";
            }
        },
        resetTime() {
            this.complete();
            this.initTime();
        },
        initTime() {
            this.hour = 0;
            this.min = 0;
            this.sec = 0;
            this.status = 'START';
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
            document.title = timeStrings[0] + ':' + timeStrings[1] + ":" + timeStrings[2];
            return timeStrings[0] + ':' + timeStrings[1] + ":" + timeStrings[2];
        }
    }
}
</script>

<style lang="scss" scoped>
.timer__display {
  font-size:4rem;
  margin:1rem;
}
.timer__toggle, .timer__reset {
  cursor: pointer;
  margin:0 1em;
  min-width: 3.5em;
  display: inline-block;
  padding: 0.5em 1em;
  text-decoration: none;
  background: #f66;/*ボタン色*/
  color: #FFF;
  border-bottom: solid 5px #c44;
  border-radius: 3px;
  &:hover {
    transition: transform .1s;
    transform: scale(1.05);
    &:active{
      /*ボタンを押したとき*/
      transform: translateY(5px);/*下に動く*/
      box-shadow: 0px 0px 1px rgba(0, 0, 0, 0.2);/*影を小さく*/
      border-bottom: none;
    }
  }
}

</style>