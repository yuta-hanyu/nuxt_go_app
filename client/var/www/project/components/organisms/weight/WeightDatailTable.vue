<template>
  <div>
    <v-simple-table dark>
      <tbody>
        <tr>
          <td class="text-h5 pr-0" width="30%"><v-icon>mdi-dumbbell</v-icon> 体重</td>
          <td class="text-h5 text-center font-weight-black pl-0" @click="openEditDaialog">
            {{ targetData?.weight }} kg
          </td>
        </tr>
        <tr>
          <td class="text-h6 pr-0">　- 前回</td>
          <td class="text-h6 text-center font-weight-black pl-0">
            {{ getPreData.weight }} kg（{{ getPreData.registDay }}）
          </td>
        </tr>
        <tr>
          <td class="text-h6 pr-0">　- 目標</td>
          <td class="text-h6 text-center font-weight-black pl-0">
            {{ getTargetWeights }}
          </td>
        </tr>
        <tr>
          <td class="text-h6 pr-0"><v-icon>mdi-food</v-icon> 食事</td>
          <td class="text-center pl-0">
            <v-rating
              background-color="warning lighten-1"
              color="warning"
              empty-icon="mdi-food-apple-outline"
              full-icon="mdi-food-apple"
              length="5"
              size="28"
              v-model="targetData.meet"
              @input="setData"
            ></v-rating>
          </td>
        </tr>
        <tr>
          <td class="text-h6 pr-0"><v-icon>mdi-walk</v-icon> 運動</td>
          <td class="text-center pl-0">
            <v-rating
              background-color="warning lighten-1"
              color="warning"
              empty-icon="mdi-star-outline"
              full-icon="mdi-star"
              length="5"
              size="28"
              v-model="targetData.sports"
              @input="setData"
            ></v-rating>
          </td>
        </tr>
        <tr>
          <td class="text-h6 pr-0"><v-icon>mdi-file-edit-outline</v-icon> メモ</td>
          <td class="text-center pl-0">
            <v-textarea
              clearable
              background-color="grey darken-3"
              class="py-4"
              filled
              v-model="targetData.memo"
              dense
              hide-details="false"
              @change="setData"
            ></v-textarea>
          </td>
        </tr>
      </tbody>
    </v-simple-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Emit } from 'nuxt-property-decorator'
import { WeightType } from '~/types/weight'
import { weightsStore } from '~/store'

@Component({
  name: 'WeightDatailTable',
  components: {},
})
export default class WeightDatailTable extends Vue {
  @Prop({ type: Object, required: true })
  data!: WeightType

  // 体重編集ダイアログ展開
  @Emit('openEditDaialog')
  openEditDaialog(): void {}

  @Emit('setData')
  setData(): WeightType {
    return this.targetData
  }

  /**
   * 編集用weightデータ
   */
  get targetData(): WeightType {
    return this.data
  }

  /**
   * 前回データを取得
   */
  get getPreData(): WeightType {
    //配列を日付の降順にソート
    let sortWeight = JSON.parse(JSON.stringify(weightsStore.getAllWeights)) //ディープコピー
    sortWeight = sortWeight.sort((a: WeightType, b: WeightType) => {
      return a.registDay > b.registDay ? -1 : 1 //オブジェクトソート
    })
    //前回日付となるデータのみを取得
    let result = null
    for (let i = 0; i < sortWeight.length; ++i) {
      if (this.targetData!.registDay > sortWeight[i].day && !result) {
        result = sortWeight[i]
      }
    }
    return result
  }

  /**
   * 目標値
   */
  get getTargetWeights(): string {
    const calcWeight = Number(this.targetData?.weight) - weightsStore.getTargetWeights
    if (Math.sign(calcWeight) === 1) {
      return `${weightsStore.getTargetWeights}kg（あと${Math.round(calcWeight)}kg）`
    } else {
      return `${weightsStore.getTargetWeights}kg（目標達成です）`
    }
  }
}
</script>

<style scoped lang="scss"></style>
