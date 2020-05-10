<template>
    <div id="sidebar" class="sidebar">
        <div class="sidebar__tool-block">
            検討＋未実装事項
            <ul>
                <li>Sidebarに何を表示するか</li>
                <li>タブ切り替えを行うと時間がリセットされる仕様</li>
                <li>TimeのHTTP post</li>
                <li>errorハンドリング</li>
            </ul>    
        </div>
    </div>
</template>

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