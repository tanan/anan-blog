<template>
  <div class="health-data-table">
    <table>
        <thead>
          <th>日付</th>
          <th>体重</th>
          <th>体脂肪</th>
        </thead>
        <tbody>
          <tr v-for="(k, index) in Object.keys(body)" :key="index">
            <td>{{ getDisplayDate(k) }}</td>
            <td>{{ body[k].weight }}Kg</td>
            <td>{{ body[k].fat }}%</td>
          </tr>
        </tbody>
      </table>
  </div>    
</template>

<script>
export default {
  name: 'HealthDataTable',
  props: {
    body: Object
  },
  methods: {
    getDisplayDate(v) {
      let year = v.substring(0, 4)
      let month = v.substring(4, 6)
      let day = v.substring(6, 8)
      let hour = v.substring(8, 10)
      let minutes = v.substring(10, 12)
      return year + "-" + month + "-" + day + " " + hour + ":" + minutes
    }
  }
}
</script>

<style scoped lang="scss">
  .health-data-table {
    table,td,th {
      margin: 0 auto;
      border-collapse: collapse;
      border: 1px #808080 solid;
    }
    th {
      padding: 8px 16px;
      text-align: center;
    }

    td {
      padding: 8px 16px;
      text-align: center;
    }

  }
</style>