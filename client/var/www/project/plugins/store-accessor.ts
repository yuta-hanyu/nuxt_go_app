import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import weights from '~/store/weights'
import alert from '~/store/alert'
import loading from '~/store/loading'

// eslint-disable-next-line import/no-mutable-exports
let weightsStore: weights
let alertStore: alert
let loadingStore: loading

/**
 * ストアを初期化する（型推論できるモジュールとして取得する）
 */
function initializeStores(store: Store<any>): void {
  // weights を型推論できるストアモジュール化
  weightsStore = getModule(weights, store)
  alertStore = getModule(alert, store)
  loadingStore = getModule(loading, store)
}

export { initializeStores, weightsStore, alertStore, loadingStore }
