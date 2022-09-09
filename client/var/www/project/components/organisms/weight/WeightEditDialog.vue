<template>
  <div class="board" style="height: 75vh !important">
    <div class="text-h4 font-weight-black py-8 text-center">
      {{ targetData }}
      <span class="subtitle-1 font-weight-black">kg</span>
    </div>
    <CalcBtns btnHeight="80" btnWidth="80" fontSize="text-h3" @setValue="setValue" @next="next" />
    <v-row justify="center">
      <v-col cols="3" class="mx-1">
        <BaseBtn @clickBtn="back" height="60" width="90" fontSize="text-h5">Back</BaseBtn>
      </v-col>
      <v-col cols="3" class="mx-1">
        <BaseBtn @clickBtn="next" height="60" width="90" fontSize="text-h4">OK</BaseBtn>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Emit } from 'nuxt-property-decorator'
import BaseBtn from '../../atoms/BaseBtn.vue'

@Component({
  name: 'WeightEditDialog',
  components: {},
})
export default class WeightEditDialog extends Vue {
  @Prop({ type: String, required: true })
  data!: string

  @Emit('back')
  back(): void {
    this.targetData = this.data
  }

  @Emit('next')
  next(): string {
    return this.targetData
  }

  targetData: string = this.data

  /**
   * 体重をセット
   */
  setValue(value: string) {
    if (value === 'C') {
      this.targetData = ''
      return
    }
    if (this.targetData.replace('.', '').length >= 4) return // . を抜いた数字数は４桁まで入力可能
    this.targetData = `${this.targetData}${value}` //入力値を文字連結
  }
}
</script>

<style scoped lang="scss">
//
</style>
