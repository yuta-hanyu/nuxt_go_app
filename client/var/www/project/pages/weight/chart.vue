<template>
  <div class="board">
    <v-container>
      <v-row justify="center">
        <v-col cols="2">
          <v-icon large right @click="calcTerm('minus')">mdi-arrow-left</v-icon>
        </v-col>
        <span v-for="(val, index) in termNumValue" :key="index">
          <v-col cols="1">
            <v-btn
              @click="changeTermType(val)"
              small
              fab
              rounded
              outlined
              elevation="10"
              :style="$FONT_STYLE.courier"
              class="subtitle-1 font-weight-bold"
            >
              {{ val }}
            </v-btn>
          </v-col>
        </span>
        <v-col align="left" cols="2">
          <v-icon large left @click="calcTerm('plus')">mdi-arrow-right</v-icon>
        </v-col>
      </v-row>

      <v-row justify="center" align-content="center">
        <LineChart
          :chartData="chartData"
          :options="options"
          class="mt-1"
          ref="lineChart"
          id="chart"
        />
        <div align="center">
          <BaseBtn
            @clickBtn="$router.push('/weight')"
            height="40"
            width="80"
            fontSize="text-h6 my-3"
          >
            Back
          </BaseBtn>
        </div>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import LineChart from '../../components/organisms/LineChart.vue'
import { Component, Vue } from 'nuxt-property-decorator'
import { WeightType } from '../../types/weight'
import { weightsStore } from '~/store'
@Component({
  name: 'Chart',
  components: {
    LineChart,
  },
})
export default class index extends Vue {
  head() {
    return {
      title: this.$PAGE_TITLE.weight,
    }
  }
  $refs!: {
    lineChart: LineChart
  }

  async mounted() {
    await weightsStore.fetchWeights() //storeに値をセット
    await this.getDispPreData() //グラフのstateをセット
  }
  /** ------------------------------ data ------------------------------ */
  //@ts-ignore
  targetValue: number = weightsStore.getTargetWeights //目標値
  termNumValue: number[] = [7, 31, 90, 180] //表示種別値
  dispTermNum: number = 7 //表示種別
  dispTermDate: Date = new Date() //表示期間基準日（今日）
  //グラフ表示必須データ
  chartData = {
    labels: [] as string[], //x軸（日付）
    datasets: [
      {
        type: 'line',
        label: '体重',
        spanGaps: true,
        data: [] as string[] | null, //y軸（体重）
        backgroundColor: 'rgba(60, 160, 220, 0.3)', //線の下側の領域の色
        borderColor: 'rgba(60, 160, 220, 0.8)', //ボーダーの色
        pointRadius: 10, //ポイントの大きさを指定
        // borderWidth: 2,//線の太さ
        yAxisID: 'y-axis-1',
      },
      {
        type: 'line',
        label: '目標',
        data: [] as string[] | null, //目標値
        borderColor: 'red',
        lineTension: 0,
        fill: false, //線の下側の領域に色を塗るか否かの指定。 false で無色になる。
        pointRadius: 0, //ポイントの大きさを指定
        yAxisID: 'y-axis-1',
      },
      {
        label: 'hidden',
        hidden: true, //データ非表示フラグ
        data: [] as string[], //hiddenデータ（id）
      },
      {
        type: 'bar',
        label: '運動',
        data: [] as string[] | null, //運動量
        backgroundColor: 'rgba(145, 196, 52, 0.28)', //ボーダーの色
        yAxisID: 'y-axis-2',
      },
      {
        type: 'bar',
        label: '食事',
        data: [] as string[] | null, //食事
        backgroundColor: 'rgba(196, 97, 52, 0.28)',
        yAxisID: 'y-axis-2',
      },
    ],
  }

  options = {
    legend: {
      display: true, //項目ごとのラベル表示するか否か
      labels: {
        //hiddenのラベル非表示
        filter: (items: any) => {
          return items.text != 'hidden'
        },
      },
    },
    responsive: true,
    onClick: (e: Event, item: Event) => {
      this.showDetail(item)
    },
    maintainAspectRatio: false,
    // title: {
    //   display: true,
    //   text: '履歴',
    //   fontSize: 20,
    // },
    tooltips: false,
    scales: {
      yAxes: [
        {
          id: 'y-axis-1',
          type: 'linear',
          position: 'left',
          ticks: {
            beginAtZero: true,
            suggestedMax: 100,
            min: 40,
            position: 'left',
          },
        },
        {
          id: 'y-axis-2',
          type: 'linear',
          min: 0,
          max: 5,
          position: 'right',
        },
      ],
    },
  }

  /** ------------------------------ methos ------------------------------ */
  /**
   * 詳細ページ展開
   */
  showDetail(item: any) {
    if (!(item[0] instanceof Number) && !item[0]) return
    const index = item[0]._index
    const id = item[0]._chart.config.data.datasets[2].data[index]
    this.$router.push(`/weight/${id}`)
  }

  /**
   * 表示期間種別変更
   */
  changeTermType(num: number) {
    this.dispTermNum = num
    this.getDispPreData()
  }

  /**
   * Y軸、x軸のデータ定義
   */
  async getDispPreData() {
    // dataを初期化
    this.chartData.labels = []
    for (let i = 0; i < 5; i++) {
      this.chartData.datasets[i].data = []
    }

    // storeデータをディープコピー
    // const sortObj = this.getWeight.map((weight) => ({ ...weight }))

    const lastDay = new Date(this.dispTermDate.getTime()) // ディープコピー
    let startDay = new Date(lastDay.getTime()) // 一時的に開始日を入れる（参照コピー回避）
    startDay.setDate(startDay.getDate() - (this.dispTermNum - 1)) // 開始日からn日前に変更

    for (let d = new Date(startDay); d <= new Date(lastDay); d.setDate(d.getDate() + 1)) {
      // 日付配列（x軸）を生成
      const month = d.getMonth() + 1
      const day = d.getDate()
      this.chartData.labels.push(`${month}/${day}`)

      // 体重配列（y軸）、id配列（hiddenデータ）を生成
      let pushWeight: null | string = null
      let pushId: null | string = null
      let pushSports: null | string = null
      let pushMeet: null | string = null

      this.getWeight.map((obj: WeightType) => {
        const dayData = new Date(obj.day)
        const dayMonth = dayData.getMonth() + 1
        const dayDay = dayData.getDate()
        if (`${dayMonth}/${dayDay}` === `${month}/${day}`) {
          pushWeight = obj.weight
          pushId = String(obj.id)
          pushSports = String(obj.sports)
          pushMeet = String(obj.meet)
        }
      })
      this.chartData.datasets[0].data!.push(pushWeight!) // 体重配列（y軸）
      this.chartData.datasets[2].data!.push(String(pushId)) //id配列（hiddenデータ）
      this.chartData.datasets[3].data!.push(String(pushSports)) //運動配列（y軸）
      this.chartData.datasets[4].data!.push(String(pushMeet)) //食事配列（y軸）
    }
    //目標値配列を生成
    for (let i = 0; i < this.chartData.datasets[0].data!.length; i++) {
      this.chartData.datasets[1].data!.push(String(this.targetValue)) //TODO 目標値は後で代入できるよう修正
    }

    // グラフを再描画する
    this.$refs.lineChart && this.$refs.lineChart.update()
  }

  /**
   * 表示範囲を変更
   */
  calcTerm(str: 'plus' | 'minus') {
    str === 'plus'
      ? this.dispTermDate.setDate(this.dispTermDate.getDate() + this.dispTermNum)
      : this.dispTermDate.setDate(this.dispTermDate.getDate() - this.dispTermNum)
    this.getDispPreData() // 表示データを再生成
  }

  /** ------------------------------ computed ------------------------------ */
  /**
   * store getter
   */
  get getWeight(): WeightType[] {
    return weightsStore.getAllWeights
  }
}
</script>

<style scoped lang="scss">
#chart {
  width: 95%;
  background: rgb(55, 56, 55);
}
</style>
