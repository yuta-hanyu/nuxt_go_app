import { Context } from '@nuxt/types'
import { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import { loadingStore } from '~/store'

// AxiosRequestConfigを拡張して、Nuxtのエラーページを表示するかどうかのフラグを保持できるようにする
// export type CustomRequestConfig = AxiosRequestConfig & {
//   dontDisplayErrorPage?: boolean
// }

/**
 *  axiosの共通設定を行う
 */
export default ({ $axios, error, app }: Context) => {
  $axios.defaults.timeout = 10000

  $axios.onRequest((config: AxiosRequestConfig) => {
    loadingStore.setLoading(true)
  })

  $axios.onResponse((response: AxiosResponse) => {
    loadingStore.setLoading(false)
  })

  $axios.onError((axiosError: AxiosError) => {
    // const res = axiosError.response
    // const status = res?.status ?? 500
    loadingStore.setLoading(false)
    app.$alert({ message: ['エラーが発生しました'], type: 'error', dispFlag: true })
    // if ((axiosError.config as CustomRequestConfig)?.dontDisplayErrorPage) {
    //   // api利用時にconfigとして`dontDisplayErrorPage: true`を渡した時のみエラーページの表示をスキップする
    //   return
    // }

    // 共通で実行したいエラー処理
    // error({ statusCode: status }) // エラーページを表示
  })
}
