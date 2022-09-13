import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { WeightType } from '../types/weight'
import { $axios } from '~/utils/api'
import moment from 'moment'

@Module({ stateFactory: true, namespaced: true, name: 'weights' })
export default class weights extends VuexModule {
  weights: WeightType[] = []
  targetWeight: number = 0

  get getAllWeights() {
    return this.weights
  }
  get getTargetWeights() {
    return this.targetWeight
  }

  @Mutation
  setWeights(payload: WeightType[]) {
    this.weights = payload
  }

  @Mutation
  setTargetWeight(payload: number) {
    this.targetWeight = payload
  }

  @Action({})
  async fetchWeights() {
    const response = await $axios.$get('/weightMocks')
    const res = await $axios.$get('/weights')
    this.setWeights(res)
    this.setTargetWeight(70)
  }

  @Action({ rawError: true })
  async addWeight(payload: WeightType) {
    //TODO オートインクリメントのためバックエンドできたら削除
    // const idAarray: number[] = []
    // this.weights.map((obj) => {
    //   idAarray.push(obj.id)
    // })
    // payload.id = Math.max(...idAarray) + 1
    /////////////////////
    payload.registDay = moment(payload.registDay).format()
    payload.createdAt = moment(payload.registDay).format()
    let res = await $axios.$post('/weights', payload)
    console.warn(res)

    this.fetchWeights()
  }

  @Action({})
  async editWeight(payload: WeightType) {
    await $axios.$put('/weightMocks', payload)
    this.fetchWeights()
  }

  @Action({ rawError: true })
  async deleteWeight(payload: WeightType) {
    await $axios.$delete('/weightMocks', { data: payload })
    this.fetchWeights()
  }

  // /**
  //  * Todo を削除する
  //  * @param todo 削除する Todo インスタンス
  //  */
  // @Mutation
  // remove(todo: Todo) {
  //   this.todos.splice(this.todos.indexOf(todo), 1)
  // }
}
