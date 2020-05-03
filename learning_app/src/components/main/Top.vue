<template>
    <div>
      <div id="main-title">~ がくしゅうあぷり（仮） ~</div>
      <router-link class="menu-button" to="/login" tag="button">ログイン</router-link>
      <router-link class="menu-button" to="/history" tag="button">学習履歴</router-link>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
#main-title {
    font-size: 2rem;
}
.menu-button {
    margin:10px;
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