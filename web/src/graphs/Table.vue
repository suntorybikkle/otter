<template>
    <table class="record__table">
      <tr>
        <th>No.</th>
        <th>StudyId</th>
        <th>SubjectId</th>
        <th>StudyTime [s]</th>
        <th>dateTime</th>
      </tr>
      <tr v-for="(value, index) in json" :key="index">
        <td>{{index}}</td>
        <td>{{value.studyId}}</td>
        <td>{{value.subId}}</td>
        <td>{{value.studyTime}}</td>
        <td>{{value.dateTime}}</td>
      </tr>
    </table>
</template>

<script>
export default{
    props:['json']
}
</script>

<style lang="scss" scoped>
table{
  width: 95%;
  border-collapse: collapse;
  border-spacing: 0;
  margin:auto;
  th,td{
    padding:10px 0;
    text-align: center;
  }
  tr:nth-child(odd) {
    background-color: #eee;
  }
}
</style>