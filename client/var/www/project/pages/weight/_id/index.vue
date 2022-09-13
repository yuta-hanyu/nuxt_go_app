<template>
  <div class="board">
    <v-container>
      <Alert />
      <Loading />
      <v-row justify="center" align-content="center">
        <v-col cols="12" class="my-2">
          <div class="text-h4 text-center" :style="$FONT_STYLE.courier">{{ dispDay }}</div>
        </v-col>
        <v-col cols="12">
          <WeightDatailTable
            @setData="setData"
            :data="targetData"
            @openEditDaialog="weightEditDialog = !weightEditDialog"
          />
        </v-col>
        <v-col cols="12">
          <v-row justify="center" align-content="center">
            <v-col cols="3" class="mx-1">
              <BaseBtn
                height="40"
                width="80"
                fontSize="text-h6"
                @clickBtn="$router.push('/weight/chart')"
                >Back</BaseBtn
              >
            </v-col>
            <v-col cols="3" class="mx-1">
              <BaseBtn height="40" width="80" fontSize="text-h6" @clickBtn="deleteWeight">
                Delete
              </BaseBtn>
            </v-col>
            <v-col cols="3" class="mx-1">
              <BaseBtn height="40" width="80" fontSize="text-h6" @clickBtn="editWeight">OK</BaseBtn>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-container>
    <v-dialog v-model="weightEditDialog" persistent eager>
      <WeightEditDialog
        :data="targetData?.weight"
        @next="setWeight"
        @back="weightEditDialog = !weightEditDialog"
      ></WeightEditDialog>
    </v-dialog>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { WeightType } from '../../../types/weight'
import { weightsStore } from '~/store'

@Component({
  name: 'DetailWeight',
  components: {},
})
export default class DetailWeight extends Vue {
  head() {
    return {
      title: this.$PAGE_TITLE.weight,
    }
  }
  async mounted() {
    await weightsStore.fetchWeights() //storeに値をセット
  }

  /** ------------------------------ data ------------------------------ */
  weightEditDialog: boolean = false //weightEditDialog表示フラグ
  targetData: WeightType | undefined = this.weight //編集用データ

  /** ------------------------------ method ------------------------------ */
  /**
   * 削除開始
   */
  deleteWeight() {
    if (this.targetData === undefined) return
    weightsStore.deleteWeight(this.targetData).then(() => {
      //@ts-ignore
      this.$alert({ message: ['削除しました'], type: 'success', dispFlag: true })
      this.$router.push('/weight/chart')
    })
  }

  /**
   * 編集データへ体重をセット
   */
  setWeight(data: string) {
    this.targetData!.weight = data
    this.weightEditDialog = false
  }

  /**
   * 詳細データをセット
   */
  setData(data: WeightType) {
    this.targetData!.meet = data.meet
    this.targetData!.sports = data.sports
    this.targetData!.memo = data.memo
  }

  /**
   * 編集開始
   */
  editWeight() {
    weightsStore.editWeight(this.targetData!).then(() => {
      //@ts-ignore
      this.$alert({ message: ['更新しました'], type: 'success', dispFlag: true })
    })
  }

  /** ------------------------------ computed ------------------------------ */
  /**
   * 対象データ取得
   */
  get weight(): WeightType | undefined {
    let targetWeight: WeightType | undefined = weightsStore.getAllWeights.find(
      (obj) => Number(this.$route.params.id) === obj.id
    )
    if (targetWeight === undefined) {
      this.$router.push('/weight/chart')
    }
    return targetWeight
  }

  /**
   * 対象データの日付を整形
   */
  get dispDay(): string {
    const day = new Date(this.weight!.registDay)
    const year = day.getFullYear()
    const month = day.getMonth() + 1
    const date = day.getDate()
    const dayOfWeek = day.getDay() // 曜日(数値)
    const dayOfWeekStr = ['日', '月', '火', '水', '木', '金', '土'][dayOfWeek] // 曜日(日本語表記)
    return `${year}年${month}月${date}日(${dayOfWeekStr})`
  }
}
</script>

<style></style>
