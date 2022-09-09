<template>
  <v-btn
    @click="clickBtn"
    :style="$FONT_STYLE.courier"
    rounded
    :height="height"
    :width="width"
    outlined
    elevation="10"
    :class="fontSize"
  >
    <slot></slot>
  </v-btn>
</template>

<script lang="ts">
import { Component, Vue, Emit, Prop } from 'nuxt-property-decorator'

@Component({
  name: 'BaseBtn',
})
export default class BaseBtn extends Vue {
  @Prop({ type: String, default: 40 })
  height!: string
  @Prop({ type: String, default: 80 })
  width!: string
  @Prop({ type: String, default: 'text-h6' })
  fontSize!: string
  @Emit('clickBtn')
  clickBtn(): void {}
}
</script>

<style></style>
