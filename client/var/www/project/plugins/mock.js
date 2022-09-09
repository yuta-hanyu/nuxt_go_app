import mock from '~/mocks/$mock.ts'

export default ({ $axios }) => {
  mock($axios).setDelayTime(1000) // axiosのモック有効化
  // mock($axios)
}
