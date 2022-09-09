import { Module, VuexModule, Mutation } from 'vuex-module-decorators'
import { AlertType } from '~/types/alert'

@Module({ stateFactory: true, namespaced: true, name: 'alert' })
export default class alert extends VuexModule {
  alert: AlertType = {
    message: [],
    type: '',
    dispFlag: false,
  }

  get getAlert() {
    return this.alert
  }

  @Mutation
  setAlert(payload: AlertType) {
    this.alert = { ...this.alert, ...payload }
  }
}
