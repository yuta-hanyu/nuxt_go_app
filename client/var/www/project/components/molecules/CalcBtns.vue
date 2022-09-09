<template>
  <v-row justify="center" class="mx-5">
    <v-col v-for="(val, index) in btnValue" :key="index" cols="4">
      <v-btn
        :style="$FONT_STYLE.courier"
        rounded
        :height="btnHeight"
        :width="btnWidth"
        outlined
        elevation="10"
        :class="fontSize"
        @click="setValue(val)"
        >{{ val }}</v-btn
      >
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { Component, Vue, Prop, Emit } from 'nuxt-property-decorator'

@Component({
  name: 'CalcBtns',
  components: {},
})
export default class CalcBtns extends Vue {
  @Prop({ type: String, default: '80' })
  btnHeight!: string

  @Prop({ type: String, default: '80' })
  btnWidth!: string

  @Prop({ type: String, default: 'text-h3' })
  fontSize!: string

  // 入力値を親へセット
  @Emit('setValue')
  setValue(value: string): string {
    return value
  }

  @Emit('next')
  next(): void {
    // return this.weight
  }

  @Emit('back')
  back(): void {}

  btnValue: string[] = ['1', '2', '3', '4', '5', '6', '7', '8', '9', 'C', '0', '.'] //ボタンテキスト
}
</script>

<style scoped lang="scss"></style>
