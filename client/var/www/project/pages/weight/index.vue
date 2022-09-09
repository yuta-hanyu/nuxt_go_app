<template>
  <div class="board">
    <v-container>
      <!-- <client-only><Alert /></client-only> -->
      <Alert />
      <Loading />
      <v-row justify="center" align-content="center">
        <v-col cols="10" md="8" align="center">
          <v-row justify="center" align-content="center" class="main-display my-2">
            <v-col cols="6" class="my-auto">
              <div class="text-caption">
                前回
                <span v-if="getWeight.length > 0">{{ getPreData.day }}</span>
              </div>
              <div class="text-h6">
                <span v-if="getWeight.length > 0">{{ getPreData.weight }}</span> kg
              </div>
            </v-col>
            <v-col cols="6">
              <div class="text-subtitle-1">{{ dispWeight }}</div>
              <div class="text-h4 font-weight-black">
                {{ newWeight.weight }}
                <span class="subtitle-1 font-weight-black">kg</span>
              </div>
            </v-col>
          </v-row>
        </v-col>
        <v-col cols="10" align="center">
          <div class="sab-display">
            <v-icon color="white" size="55">mdi-cog</v-icon>
            <v-icon color="white" size="55" @click="$router.push('/weight/chart')" class="mx-10">
              mdi-chart-timeline-variant
            </v-icon>
            <v-icon color="white" size="55" @click="openday">mdi-clock</v-icon>
          </div>
        </v-col>
        <v-col md="5" sm="7" align="center">
          <CalcBtns btnHeight="80" btnWidth="80" fontSize="text-h3" @setValue="setValue" />
        </v-col>
        <v-col cols="12" class="mx-1" align="center">
          <BaseBtn
            @clickBtn="weightDetailSetDialog = !weightDetailSetDialog"
            height="60"
            width="90"
            fontSize="text-h4"
            >OK
          </BaseBtn>
        </v-col>
      </v-row>
    </v-container>
    <v-dialog v-model="datePickerDialog" persistent eager>
      <DatePicker
        @back="datePickerDialog = !datePickerDialog"
        ref="datePickerRef"
        @next="setday"
        :possibleRange="$getDayISOString()"
      ></DatePicker>
    </v-dialog>
    <v-dialog v-model="weightDetailSetDialog" persistent>
      <WeightDetailSetDialog
        :data="newWeight"
        @registWeight="registWeight"
        @back="weightDetailSetDialog = !weightDetailSetDialog"
      ></WeightDetailSetDialog>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import DatePicker from '~/components/organisms/DatePicker.vue' //$refsを使用するためimport
import { WeightType } from '../../types/weight'
import { weightsStore, loadingStore } from '~/store'
//@ts-ignore
import mock from '../../mocks/$mock.ts'
import BaseBtn from '../../components/atoms/BaseBtn.vue'
// mock() // axiosのモック有効化

@Component({
  name: 'Weight',
  components: {
    DatePicker,
  },
})
export default class index extends Vue {
  head() {
    return {
      title: this.$PAGE_TITLE.weight,
    }
  }
  $refs!: {
    datePickerRef: DatePicker
  }

  async mounted() {
    await weightsStore.fetchWeights()
  }

  /** ------------------------------ data ------------------------------ */
  datePickerDialog: boolean = false //日付時刻ダイアログ表示フラグ
  weightDetailSetDialog: boolean = false //詳細設定ダイアログ表示フラグ

  //登録データ
  newWeight: WeightType = {
    id: 0,
    weight: '',
    day: this.$getDayISOString(),
    meet: 3,
    sports: 3,
    memo: '',
  }

  /** ------------------------------ methos ------------------------------ */
  /**
   * データ登録
   */
  registWeight(data: WeightType) {
    this.weightDetailSetDialog = !this.weightDetailSetDialog
    loadingStore.setLoading(true)
    this.newWeight = data
    weightsStore.addWeight(this.newWeight).then(() => {
      //@ts-ignore
      this.$alert({ message: ['登録しました'], type: 'success', dispFlag: true })
      // 初期値へ戻す
      this.newWeight = {
        id: 0,
        weight: '',
        day: this.$getDayISOString(''),
        meet: 3,
        sports: 3,
        memo: '',
      }
    })
  }

  /**
   * 体重をセット
   */
  setValue(value: string) {
    if (value === 'C') {
      this.newWeight.weight = ''
      return
    }
    if (this.newWeight.weight.replace('.', '').length >= 4) return // . を抜いた数字数は４桁まで入力可能
    this.newWeight.weight = `${this.newWeight.weight}${value}` //入力値を文字連結
  }

  /**
   * datepickerダイアログ展開
   */
  openday() {
    this.datePickerDialog = !this.datePickerDialog
    const today = this.$getDayISOString() // 今日の日付をdatepickerにセット
    this.$refs.datePickerRef.setdate(today)
  }

  /**
   * 選択した日付をnewWeightにセット
   */
  setday(day: string): void {
    this.newWeight.day = day
    this.datePickerDialog = !this.datePickerDialog
  }

  /** ------------------------------ computed ------------------------------ */
  /**
   * store getter
   */
  get getWeight(): WeightType[] {
    return weightsStore.getAllWeights
  }

  /**
   * 表示用日付時刻取得
   */
  get dispWeight(): string {
    return this.newWeight.day.replace(/-/g, '/')
  }

  /**
   * 前回データを取得
   */
  get getPreData(): WeightType {
    let result = JSON.parse(JSON.stringify(this.getWeight)) //ディープコピー
    result = result.sort((a: WeightType, b: WeightType) => {
      return a.day > b.day ? -1 : 1 //オブジェクトソート
    })
    result = result.shift()
    result.day = result.day.replace(/-/g, '/')
    return result
  }
}
</script>

<style lang="scss" scoped>
.sab-display {
  background-color: rgb(114, 114, 110);
  height: 60px;
  opacity: 0.5;
  border-radius: 5%;
}
</style>
