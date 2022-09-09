<template>
  <v-app dark>
    <v-app-bar
      app
      color="brown lighten-3"
      elevation="10"
      dense
      height="60"
      scroll-target="#scrolling-techniques-7"
    >
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <v-spacer />
      <v-toolbar-title class="text-h5">{{ getPageName }}</v-toolbar-title>
      <v-spacer />
      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </v-app-bar>
    <v-sheet id="scrolling-techniques-7" class="overflow-y-auto mx-auto" max-height="100%">
      <v-main>
        <v-container class="container px-0 py-0">
          <Nuxt />
        </v-container>
      </v-main>
    </v-sheet>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
@Component({
  name: 'Default',
})
export default class extends Vue {
  get getPageName(): string {
    if (this.$route.path.indexOf('/weight') === 0) {
      return '体重管理'
    } else {
      return ''
    }
  }
}
</script>

<style scoped lang="SCSS">
.container{
  height: 100vh;
}
</style>
