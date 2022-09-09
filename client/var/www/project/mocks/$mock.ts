/* eslint-disable */
import { AxiosInstance } from "axios"
import mockServer from "axios-mock-server"
import mock0 from "./weightMocks"

export default (client?: AxiosInstance) =>
  mockServer(
    [
      {
        path: "/weightMocks",
        methods: mock0 as any,
      },
    ],
    client,
    ""
  )
