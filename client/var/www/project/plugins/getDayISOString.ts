/**
 * ISO 8601 date 文字列 (YYYY-MM-DD) のstringデータを作成
 */

import { Context } from '@nuxt/types'
import Vue from 'vue'

declare module 'vue/types/vue' {
  interface Vue {
    $getDayISOString(day?: string): string
  }
}

const getDayISOString = (day?: string): string => {
  if (day === undefined) {
    //引数なしの場合は今日を返す
    return new Date(Number(Date.now()) - new Date().getTimezoneOffset() * 60000)
      .toISOString()
      .substr(0, 10)
  } else {
    return new Date(Number(new Date(day)) - new Date().getTimezoneOffset() * 60000)
      .toISOString()
      .substr(0, 10)
  }
}

export default ({ app }: Context, inject: any) => {
  inject('getDayISOString', getDayISOString)
}
