<template>
  <!-- <v-alert
    border="left"
    dense
    dismissible
    elevation="24"
    :type="getAlert.type"
    v-if="getAlert.dispFlag"
  >
    <div v-for="(msg, index) in getAlert.message" :key="index">
      <span v-if="getAlert.message.length > 1">・</span>{{ msg }}
    </div>
  </v-alert> -->
  <v-overlay z-index="99" :value="getLoading" opacity="0.7">
    <v-progress-circular
      size="70"
      width="10"
      color="light-blue lighten-1"
      indeterminate
      class="mx-auto align-self-center s"
    >
    </v-progress-circular>
  </v-overlay>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import { loadingStore } from '~/store'

@Component({
  name: 'Loading',
})
export default class Loading extends Vue {
  /**
   * store getter
   */
  get getLoading(): boolean {
    return loadingStore.loading
  }
}
</script>

<style></style>
