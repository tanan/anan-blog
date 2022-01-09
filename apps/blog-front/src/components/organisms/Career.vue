<template>
  <div class="career">
    <div class="title">職歴・経歴</div>
    <CareerContent :career="career" v-for="career in careers" :key="career.id" />
  </div>
</template>

<script>
import CareerContent from '@/components/molecules/CareerContent.vue'
export default {
  name: 'Career',
  components: {
    CareerContent
  },
  data () {
    return {
      careers: [
        {
          id: '1',
          period: '2021年10月 - 現在',
          company: {
            name: '静岡大学大学院',
            url: 'https://www.shizuoka.ac.jp/'
          },
          role: '総合科学研究科情報学専攻修士課程',
          description: '量子情報科学分野の研究。現在は主に関節制御やエラー軽減等の研究を行っている',
          technology: 'python, Qiskit, Qulacs'
        },
        {
          id: '2',
          period: '2016年7月 - 現在',
          company: {
            name: '株式会社ユーザベース',
            url: 'https://www.uzabase.com/'
          },
          role: 'Site Reliablity Engineer',
          description: 'SaaSサービスの開発、運用',
          technology: 'on-premises, Xen Server, Proxmox, AWS, GCP, Linux, Docker, Kubernetes, Istio, MySQL, PostgreSQL, Elasticsearch, Aerospike, Apache, Nginx, Squid, LVS, Keepalived, HAProxy, Fluentd, Unbound, Ansible, Terraform, Hinemos, Jenkins, Buildkite, ArgoCD, Prometheus, Nagios, Cacti, Golang, Python, Java, Kotlin, Angular Dart'
        },
        {
          id: '3',
          period: '2012年4月 - 2016年6月',
          company: {
            name: '京セラコミニュケーションシステム株式会社',
            url: 'https://www.kccs.co.jp/'
          },
          role: 'Infra Engineer',
          description: '広告配信システムの開発、インフラ環境構築と運用',
          technology: 'on-premises, AWS, C#, Java, Ansible, Windows Server 2008/2012, HyperV, IIS, Aerospike, MySQL, SQLServer, Fluentd, Treasure Data'
        },
        {
          id: '4',
          period: '2008年4月 - 2012年3月',
          company: {
            name: '熊本大学',
            url: 'https://www.kumamoto-u.ac.jp/'
          },
          role: '理学部理学科物理プログラム',
          description: '物性理論、計算機シミュレーションの研究室に所属',
          technology: 'Fortran 90'
        },
        {
          id: '5',
          period: '2005年4月 - 2008年3月',
          company: {
            name: '大分県立竹田高等学校',
            url: 'http://kou.oita-ed.jp/taketa/'
          },
          role: '普通科',
          description: '',
          technology: ''
        },
      ]
    }
  }
}
</script>

<style lang="scss">
  .career {
    margin-left: auto;
    margin-right: auto;

    .title {
      padding-top: 8%;
      text-align: center;
      font-size: 28px;
      font-weight: 600;
    }
  }
</style>