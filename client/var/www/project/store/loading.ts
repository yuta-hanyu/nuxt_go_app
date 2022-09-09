import { Module, VuexModule, Mutation } from 'vuex-module-decorators'

@Module({ stateFactory: true, namespaced: true, name: 'loading' })
export default class loading extends VuexModule {
  loading: boolean = false

  get getLoading() {
    return this.loading
  }

  @Mutation
  setLoading(payload: boolean) {
    this.loading = payload
  }
}
