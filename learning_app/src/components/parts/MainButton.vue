Start, Stopボタンの制御を行う

<template>
    <div>
        <button @click="moveStatus">{{name.toUpperCase()}}</button>
        <button @click="resetTime">RESET</button>
    </div>
</template>

<script>
export default {
    props:['name'],
    methods:{
        moveStatus() {
            if(this.name === "start") {
                this.$emit('toggle','stop');
            }else {
                this.$emit('toggle','start');
            }
        },
        resetTime() {
            this.$emit('reset',1);
        }
    }
}
</script>

<style scoped>
button {
    margin:30px;
    width:140px;
    cursor: pointer;
    font-weight: bold;
    font-size:1rem;
    padding: 10px 30px;
    background-color:#f66;
    color:#fff;
    border-style:none;
    box-shadow: 2px 2px 3px 1px #666;
}
</style>