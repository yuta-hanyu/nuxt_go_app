/**
 * フラッシュメッセージ表示
 */

import { Context } from '@nuxt/types'
import { AlertType } from '~/types/alert'
import { alertStore } from '~/store'

declare module 'vue/types/vue' {
  interface Vue {
    readonly $alert: AlertType
  }
}

const alert = (payload: AlertType) => {
  alertStore.setAlert(payload)
  //3秒後に消える
  setTimeout(() => {
    alertStore.setAlert({ message: [], type: '', dispFlag: false })
  }, 3000)
}

export default ({ app }: Context, inject: any) => {
  inject('alert', alert)
}
