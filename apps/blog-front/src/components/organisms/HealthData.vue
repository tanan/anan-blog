<template>
  <div class="health-data-section">
    <h2>身体データ記録</h2>
    <p>タニタの体組成計で記録したデータを載せるページです。APIの都合上、体重と体脂肪のみ連携します。</p>
    <HealthDataTable :body=body />
  </div>    
</template>

<script>
import HealthDataTable from '@/components/molecules/HealthDataTable.vue'
export default {
  name: 'HealthData',
  components: {
    HealthDataTable
  },
  data() {
    return {
      body: {}
    };
  },
  async created() {
    // this.body = await this.getBodyData();
    let resp = await this.getBodyData();
    let body = {}
    resp.forEach(v => {
      let d = {}
      if (v.date in body) {
        d = body[v.date]
      }
      if(v.tag === "6021") {
        d.weight = v.keydata
      } else {
        d.fat = v.keydata
      }
      body[v.date] = d
    });
    console.log(body)
    this.body = body
  },
  methods: {
    getBodyData: async () => {
      const host = "https://health-planet-api-cmfot4feta-uc.a.run.app"
      const fetchUrl = host + "/v1/me"
      const fetchOptions = {
        method: "GET",
        headers: {
          "Content-Type": "application/json"
        },
      }

      try {
        const response = await fetch(fetchUrl, fetchOptions).then(resp =>
          resp.json()
        )
        return response
      } catch (error) {
        throw new Error("Could not receive the data from health care API!")
      }
    }
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .health-data-section {
    width: 58%;
    margin-right: auto;
    margin-left: auto;

    h2 {
      font-size: 1.8em;
      text-align: center;
    }

    p {
      text-align: center;
    }
  }
</style>