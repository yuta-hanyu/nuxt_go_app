<template>
  <v-date-picker locale="ja" v-model="date" color="green lighten-1" scrollable :max="possibleRange">
    <v-row justify="center">
      <v-col cols="3" class="mx-1">
        <BaseBtn @clickBtn="back" height="40" width="80" fontSize="text-h6">Back</BaseBtn>
      </v-col>
      <v-col cols="3" class="mx-1">
        <BaseBtn @clickBtn="next" height="40" width="80" fontSize="text-h6">OK</BaseBtn>
      </v-col>
    </v-row>
  </v-date-picker>
</template>

<script lang="ts">
import { Component, Vue, Prop, Emit } from 'nuxt-property-decorator'

@Component({
  name: 'DatePicker',
  components: {},
})
export default class DatePicker extends Vue {
  @Prop({
    type: String,
    default: '',
  })
  possibleRange!: string

  @Emit('back')
  back(): void {}

  @Emit('next')
  next(): string {
    return this.date
  }
  date: string = this.$getDayISOString()

  setdate(day: string) {
    this.$getDayISOString(day)
  }
}
</script>

<style scoped lang="scss"></style>
