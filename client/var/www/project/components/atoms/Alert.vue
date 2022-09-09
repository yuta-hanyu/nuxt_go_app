<template>
  <v-alert
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
  </v-alert>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'nuxt-property-decorator'
import { alertStore } from '~/store'
import { AlertType } from '~/types/alert'

@Component({
  name: 'Alert',
})
export default class Alert extends Vue {
  /**
   * store getter
   */
  get getAlert(): AlertType {
    return alertStore.alert
  }
}
</script>

<style></style>
