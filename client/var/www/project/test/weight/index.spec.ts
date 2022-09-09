import { mount, createLocalVue, shallowMount } from '@vue/test-utils'
import index from '~/pages/weight/index.vue'
import Vuex from 'vuex'
import * as weights from '~/store/weights'
import DatePicker from '~/components/organisms/DatePicker.vue'
import getDayISOString from '~/plugins/getDayISOString'

describe('index', () => {
  let wrapper
  let getters
  let store: any
  const localVue = createLocalVue()
  // localVue.use(Vuex)

  beforeEach(() => {
    getters = { getAllWeights: jest.fn() }

    store = new Vuex.Store({
      state: {
        weights,
      },
      getters,
    })
  })

  // })
  it('コンポーネントが表示される', () => {
    wrapper = shallowMount(index, {
      localVue,
      //@ts-ignore
      store,
      mocks: { getDayISOString, weights },
    })

    // wrapperが存在する
    expect(wrapper.exists()).toBe(true)
    // expect(wrapper.vm).toBeTruthy()
  })
})
