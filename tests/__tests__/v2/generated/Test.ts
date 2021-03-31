// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'

import { GoogleProtobuf } from './google/protobuf/Timestamp'
import { Common as Common_Common } from './common/Common'
import { Common as Common_Extra } from './common/Extra'

export namespace Foo {
  export type EnumType = 'UNKNOWN' | 'ADMIN' | 'USER'
  export type Role = 'VIEW' | 'EDIT'

  export interface IRequest {
    id?: number
  }

  /**
   * @deprecated
   */
  export interface INested {
    title?: string
  }

  export interface ITest {
    fieldInt32?: number
    fieldInt32Repeated?: number[]
    fieldDouble?: number
    fieldDoubleRepeated?: number[]
    fieldFloat?: number
    fieldFloatRepeated?: number[]
    fieldUint32?: number
    fieldUint32Repeated?: number[]
    fieldUint64?: number
    fieldUint64Repeated?: number[]
    fieldSint32?: number
    fieldSint32Repeated?: number[]
    fieldSint64?: number
    fieldSint64Repeated?: number[]
    fieldFixed32?: number
    fieldFixed32Repeated?: number[]
    fieldFixed64?: number
    fieldFixed64Repeated?: number[]
    fieldSfixed32?: number
    fieldSfixed32Repeated?: number[]
    fieldSfixed64?: number
    fieldSfixed64Repeated?: number[]
    fieldBool?: boolean
    fieldBoolRepeated?: boolean[]
    fieldString?: string
    fieldStringRepeated?: string[]
    fieldBytes?: Uint8Array
    fieldBytesRepeated?: Uint8Array[]
    fieldEnum?: EnumType
    fieldEnumRepeated?: Role[]
    message?: INested
    messageRepeated?: INested[]
    timestamp?: GoogleProtobuf.ITimestamp
    timestampRepeated?: GoogleProtobuf.ITimestamp[]
    otherPkgMessage?: Common_Common.IOtherPkgMessage
    otherPkgMessageRepeated?: Common_Common.IOtherPkgMessage[]
    fieldInt64?: number
    fieldInt64Repeated?: number[]
  }

  export interface ISecondTest {
    extraPkgMessage?: Common_Extra.IExtraPkgMessage
  }
}
