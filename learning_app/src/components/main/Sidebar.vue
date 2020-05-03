<template>
    <div id="side">
        <div id="sel-color">
            <h3>Select color</h3>
            <div class="sel-color-default c-sel-color">
                <div class="box-color sel-box-color-default"></div>default
            </div>
            <div class="sel-color-red c-sel-color">
                <div class="box-color sel-box-color-red"></div>input
            </div>
            <div class="sel-color-blue c-sel-color">
                <div class="box-color sel-box-color-blue"></div>output
            </div>
            <div class="sel-color-green c-sel-color">
                <div class="box-color sel-box-color-green"></div>relax
            </div>
            <div class="sel-color-yellow c-sel-color">
                <div class="box-color sel-box-color-yellow"></div>yellow
            </div>
        </div>
    </div>
</template>

<style scoped>
#sel-color {
    text-align: left;
}
.c-sel-color{
    margin:10px;
    display: flex;
    font-size: 1.5rem;
    cursor: pointer;
    padding: 10px;
    box-shadow: 2px 2px 3px 1px #666;
}
.box-color {
    width:30px;
    height:30px;
    border-radius:100%;
    margin: 0 10px;
    background-color: white;
    box-shadow: 2px 2px 3px 1px #666;
}
.sel-box-color-default {
    background-color: black;
}
.sel-box-color-red {
    background-color: red;
}
.sel-box-color-blue {
    background-color: blue;
}
.sel-box-color-green {
    background-color: green;
}
.sel-box-color-yellow {
    background-color: yellow;
}
</style>