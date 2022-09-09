import { Context } from '@nuxt/types'

declare module 'vue/types/vue' {
  interface Vue {
    $PAGE_TITLE: { weight: string }
    FONT_STYLE: string
  }
}

const PAGE_TITLE = {
  weight: '体重管理',
}

const FONT_STYLE = {
  courier: { fontFamily: 'Courier !important' as string },
}

export default ({ app }: Context, inject: any) => {
  inject('PAGE_TITLE', PAGE_TITLE)
  inject('FONT_STYLE', FONT_STYLE)
}
