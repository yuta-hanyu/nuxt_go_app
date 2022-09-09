<template>
  <div>
    <v-simple-table dark>
      <thead>
        <tr>
          <td class="text-h5 text-center" colspan="2">extra-info</td>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="text-h6 pr-0" width="35%"><v-icon>mdi-food</v-icon> 食事</td>
          <td class="text-center pl-0">
            <v-rating
              background-color="warning lighten-1"
              color="warning"
              empty-icon="mdi-food-apple-outline"
              full-icon="mdi-food-apple"
              length="5"
              size="28"
              v-model="targetData.meet"
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
              dense
              hide-details="false"
              v-model="targetData.memo"
            ></v-textarea>
          </td>
        </tr>
        <tr>
          <td colspan="2">
            <v-row justify="center">
              <v-col cols="3">
                <BaseBtn @clickBtn="back" height="30" width="40" fontSize="text-h6">Back</BaseBtn>
              </v-col>
              <v-col cols="3">
                <BaseBtn @clickBtn="registWeight" height="30" width="40" fontSize="text-h6"
                  >OK</BaseBtn
                >
              </v-col>
            </v-row>
          </td>
        </tr>
      </tbody>
    </v-simple-table>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Emit } from 'nuxt-property-decorator'
import { WeightType } from '~/types/weight'

@Component({
  name: 'WeightDetailSetDialog',
  components: {},
})
export default class WeightDetailSetDialog extends Vue {
  @Prop({ type: Object, required: true })
  data!: WeightType

  @Emit('back')
  back(): void {}

  // 詳細データ登録
  @Emit('registWeight')
  registWeight(): WeightType {
    return this.targetData
  }

  /**
   * 登録用データ
   */
  get targetData(): WeightType {
    console.warn(this.data)
    return this.data
  }
}
</script>

<style scoped lang="scss"></style>
