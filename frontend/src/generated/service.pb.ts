/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "./fetch.pb"
export type HelloReq = {
  name?: string
}

export type HelloRes = {
  msg?: string
}

export class TestService {
  static Hello(req: HelloReq, initReq?: fm.InitReq): Promise<HelloRes> {
    return fm.fetchReq<HelloReq, HelloRes>(`/v1/Hello`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}