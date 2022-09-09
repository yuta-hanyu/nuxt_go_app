<script lang="ts">
import { Line } from 'vue-chartjs'
import { Prop } from 'nuxt-property-decorator'
import Component, { mixins } from 'vue-class-component'

@Component({
  extends: Line,
  name: 'LineChart',
})
export default class LineChart extends mixins(Line) {
  @Prop({ type: Object, default: null })
  chartData!: any

  @Prop({ type: Object, default: null })
  options!: any

  mounted() {
    this.renderChart(this.chartData, this.options)
  }

  /**
   * グラフを再描画
   */
  update() {
    this.$forceUpdate()
    console.warn(this.chartData)
    this.renderChart(this.chartData, this.options)
  }
}
</script>
