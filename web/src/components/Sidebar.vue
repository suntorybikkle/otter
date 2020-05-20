<template>
    <div id="sidebar" class="sidebar">
        <div class="sidebar__tool-block">
            <Timer></Timer>
        </div>
    </div>
</template>

<script>
import Timer from './Timer'
export default {
    components:{
        Timer
    },
   
}
</script>

<style scoped lang="scss">
.sidebar{
    font-size:20px;
    .sidebar__time{
        margin:10px;
        padding:10px 0;
        border: 1px black solid;
    }
    .sidebar__time-display{
        font-size:1.6em;
        font-weight: bold;
    }
}
</style>