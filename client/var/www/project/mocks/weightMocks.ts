import { WeightType } from '../types/weight'

let weightMocks: WeightType[] = [
  { id: 0, weight: '55', registDay: '2022-08-17', meet: 3, sports: 1, memo: 'aaaaaaaaaaaaaa' },
  { id: 1, weight: '77', registDay: '2022-08-18', meet: 3, sports: 3, memo: 'aaaaaaaaaaaaaa' },
  { id: 2, weight: '88', registDay: '2022-08-19', meet: 3, sports: 4, memo: 'aaaaaaaaaaaaaa' },
  { id: 3, weight: '22', registDay: '2022-08-20', meet: 3, sports: 5, memo: 'aaaaaaaaaaaaaa' },
  { id: 4, weight: '90', registDay: '2022-08-21', meet: 3, sports: 1, memo: 'aaaaaaaaaaaaaa' },
  { id: 5, weight: '80', registDay: '2022-08-22', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 6, weight: '70', registDay: '2022-08-23', meet: 3, sports: 3, memo: 'aaaaaaaaaaaaaa' },
  { id: 7, weight: '60', registDay: '2022-08-24', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 8, weight: '50', registDay: '2022-08-25', meet: 3, sports: 4, memo: 'aaaaaaaaaaaaaa' },
  { id: 9, weight: '100', registDay: '2022-08-26', meet: 3, sports: 5, memo: 'aaaaaaaaaaaaaa' },
  { id: 10, weight: '70', registDay: '2022-08-01', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 11, weight: '90', registDay: '2022-08-02', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 12, weight: '82', registDay: '2022-08-03', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  // {id: 13,  weight: '54', registDay: '2022-08-04' , meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 14, weight: '45', registDay: '2022-08-05', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 15, weight: '66', registDay: '2022-08-06', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 16, weight: '88', registDay: '2022-08-07', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 17, weight: '45', registDay: '2022-08-08', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 18, weight: '87', registDay: '2022-08-09', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  // {id: 19,  weight: '53', registDay: '2022-08-10' , meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 20, weight: '89', registDay: '2022-08-11', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 21, weight: '97', registDay: '2022-08-12', meet: null, sports: null, memo: '' },
  // { id: 22, weight: '40', registDay: '2022-08-13', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 23, weight: '45', registDay: '2022-08-14', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 24, weight: '53', registDay: '2022-08-15', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  { id: 25, weight: '78', registDay: '2022-08-16', meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
  // {id: 26,  weight: '53', registDay: '2022-08-10' , meet: 3, sports: 2, memo: 'aaaaaaaaaaaaaa' },
]

const targetWeight: number = 70

export default {
  get: () => [200, { weights: weightMocks, targetWeight: targetWeight }], // ???????????? [status, body?, headers?]
  post: ({ data }: { data: WeightType }) => {
    weightMocks.push(data)
    return [201]
  },
  put: ({ data }: { data: WeightType }) => {
    const index = weightMocks.findIndex((obj) => {
      return obj.id === data.id
    })
    weightMocks[index] = data
    return [201]
  },
  delete: ({ data }: { data: WeightType }) => {
    console.warn(data)
    const index = weightMocks.findIndex((obj) => {
      return obj.id === data.id
    })
    weightMocks.splice(index, 1)
    // const fixedData = weightMocks.filter((obj) => obj.id !== data.id)
    // weightMocks = { ...weightMocks, ...fixedData }
    return [201]
  },
}
