// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'
import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from './google/protobuf/Timestamp'
import { Common as Common_Common } from './common/Common'
import { Common as Common_Extra } from './common/Extra'

export namespace Foo {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export enum EnumType_Enum {
    UNKNOWN = 0,
    ADMIN = 1,
    USER = 2,
  }
  export type EnumType = keyof typeof EnumType_Enum

  export enum Role_Enum {
    VIEW = 0,
    EDIT = 1,
  }
  export type Role = keyof typeof Role_Enum

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

  export interface ICustomOptionsTest {
    requiredField: Common_Extra.IExtraPkgMessage
    typicalOptionalField?: number
    customOptionalField?: number
  }

  export interface IRequiredPropertiesTest {
    requiredField: number
    customRequiredField: number
    optionalField?: number
  }

  @protobufjs.Type.d('Request')
  export class Request
    extends protobufjs.Message<Request>
    implements ConvertibleTo<IRequest> {
    @protobufjs.Field.d(1, 'int32')
    public id?: number

    public asInterface(): IRequest {
      return this
    }
  }

  /**
   * @deprecated
   */
  @protobufjs.Type.d('Nested')
  export class Nested
    extends protobufjs.Message<Nested>
    implements ConvertibleTo<INested> {
    @protobufjs.Field.d(1, 'string')
    public title?: string

    public asInterface(): INested {
      return this
    }
  }

  @protobufjs.Type.d('Test')
  export class Test
    extends protobufjs.Message<Test>
    implements ConvertibleTo<ITest> {
    @protobufjs.Field.d(1, 'int32')
    public fieldInt32?: number

    @protobufjs.Field.d(2, 'int32', 'repeated')
    public fieldInt32Repeated?: number[]

    @protobufjs.Field.d(3, 'double')
    public fieldDouble?: number

    @protobufjs.Field.d(4, 'double', 'repeated')
    public fieldDoubleRepeated?: number[]

    @protobufjs.Field.d(5, 'float')
    public fieldFloat?: number

    @protobufjs.Field.d(6, 'float', 'repeated')
    public fieldFloatRepeated?: number[]

    @protobufjs.Field.d(7, 'uint32')
    public fieldUint32?: number

    @protobufjs.Field.d(8, 'uint32', 'repeated')
    public fieldUint32Repeated?: number[]

    @protobufjs.Field.d(9, 'uint64')
    public fieldUint64?: number

    @protobufjs.Field.d(10, 'uint64', 'repeated')
    public fieldUint64Repeated?: number[]

    @protobufjs.Field.d(11, 'sint32')
    public fieldSint32?: number

    @protobufjs.Field.d(12, 'sint32', 'repeated')
    public fieldSint32Repeated?: number[]

    @protobufjs.Field.d(13, 'sint64')
    public fieldSint64?: number

    @protobufjs.Field.d(14, 'sint64', 'repeated')
    public fieldSint64Repeated?: number[]

    @protobufjs.Field.d(15, 'fixed32')
    public fieldFixed32?: number

    @protobufjs.Field.d(16, 'fixed32', 'repeated')
    public fieldFixed32Repeated?: number[]

    @protobufjs.Field.d(17, 'fixed64')
    public fieldFixed64?: number

    @protobufjs.Field.d(18, 'fixed64', 'repeated')
    public fieldFixed64Repeated?: number[]

    @protobufjs.Field.d(19, 'sfixed32')
    public fieldSfixed32?: number

    @protobufjs.Field.d(20, 'sfixed32', 'repeated')
    public fieldSfixed32Repeated?: number[]

    @protobufjs.Field.d(21, 'sfixed64')
    public fieldSfixed64?: number

    @protobufjs.Field.d(22, 'sfixed64', 'repeated')
    public fieldSfixed64Repeated?: number[]

    @protobufjs.Field.d(23, 'bool')
    public fieldBool?: boolean

    @protobufjs.Field.d(24, 'bool', 'repeated')
    public fieldBoolRepeated?: boolean[]

    @protobufjs.Field.d(25, 'string')
    public fieldString?: string

    @protobufjs.Field.d(26, 'string', 'repeated')
    public fieldStringRepeated?: string[]

    @protobufjs.Field.d(27, 'bytes')
    public fieldBytes?: Uint8Array

    @protobufjs.Field.d(28, 'bytes', 'repeated')
    public fieldBytesRepeated?: Uint8Array[]

    @protobufjs.Field.d(29, EnumType_Enum)
    public fieldEnum?: EnumType_Enum

    @protobufjs.Field.d(30, Role_Enum, 'repeated')
    public fieldEnumRepeated?: Role_Enum[]

    @protobufjs.Field.d(33, Nested)
    public message?: Nested

    @protobufjs.Field.d(34, Nested, 'repeated')
    public messageRepeated?: Nested[]

    @protobufjs.Field.d(35, GoogleProtobuf.Timestamp)
    public timestamp?: GoogleProtobuf.Timestamp

    @protobufjs.Field.d(36, GoogleProtobuf.Timestamp, 'repeated')
    public timestampRepeated?: GoogleProtobuf.Timestamp[]

    @protobufjs.Field.d(37, Common_Common.OtherPkgMessage)
    public otherPkgMessage?: Common_Common.OtherPkgMessage

    @protobufjs.Field.d(38, Common_Common.OtherPkgMessage, 'repeated')
    public otherPkgMessageRepeated?: Common_Common.OtherPkgMessage[]

    @protobufjs.Field.d(39, 'int64')
    public fieldInt64?: number

    @protobufjs.Field.d(40, 'int64', 'repeated')
    public fieldInt64Repeated?: number[]

    public asInterface(): ITest {
      return {
        fieldInt32: this.fieldInt32,
        fieldInt32Repeated: this.fieldInt32Repeated,
        fieldDouble: this.fieldDouble,
        fieldDoubleRepeated: this.fieldDoubleRepeated,
        fieldFloat: this.fieldFloat,
        fieldFloatRepeated: this.fieldFloatRepeated,
        fieldUint32: this.fieldUint32,
        fieldUint32Repeated: this.fieldUint32Repeated,
        fieldUint64: this.fieldUint64,
        fieldUint64Repeated: this.fieldUint64Repeated,
        fieldSint32: this.fieldSint32,
        fieldSint32Repeated: this.fieldSint32Repeated,
        fieldSint64: this.fieldSint64,
        fieldSint64Repeated: this.fieldSint64Repeated,
        fieldFixed32: this.fieldFixed32,
        fieldFixed32Repeated: this.fieldFixed32Repeated,
        fieldFixed64: this.fieldFixed64,
        fieldFixed64Repeated: this.fieldFixed64Repeated,
        fieldSfixed32: this.fieldSfixed32,
        fieldSfixed32Repeated: this.fieldSfixed32Repeated,
        fieldSfixed64: this.fieldSfixed64,
        fieldSfixed64Repeated: this.fieldSfixed64Repeated,
        fieldBool: this.fieldBool,
        fieldBoolRepeated: this.fieldBoolRepeated,
        fieldString: this.fieldString,
        fieldStringRepeated: this.fieldStringRepeated,
        fieldBytes: this.fieldBytes,
        fieldBytesRepeated: this.fieldBytesRepeated,
        fieldEnum: (this.fieldEnum !== undefined
          ? EnumType_Enum[this.fieldEnum]!
          : undefined) as EnumType | undefined,
        fieldEnumRepeated: this.fieldEnumRepeated?.map(
          (e) => Role_Enum[e]! as Role
        ),
        message: this.message?.asInterface(),
        messageRepeated: this.messageRepeated?.map((o) => o.asInterface()),
        timestamp: this.timestamp?.asInterface(),
        timestampRepeated: this.timestampRepeated?.map((o) => o.asInterface()),
        otherPkgMessage: this.otherPkgMessage?.asInterface(),
        otherPkgMessageRepeated: this.otherPkgMessageRepeated?.map((o) =>
          o.asInterface()
        ),
        fieldInt64: this.fieldInt64,
        fieldInt64Repeated: this.fieldInt64Repeated,
      }
    }
  }

  @protobufjs.Type.d('CustomOptionsTest')
  export class CustomOptionsTest
    extends protobufjs.Message<CustomOptionsTest>
    implements ConvertibleTo<ICustomOptionsTest> {
    @protobufjs.Field.d(1, Common_Extra.ExtraPkgMessage)
    public requiredField!: Common_Extra.ExtraPkgMessage

    @protobufjs.Field.d(2, 'int32')
    public typicalOptionalField?: number

    @protobufjs.Field.d(3, 'int32')
    public customOptionalField?: number

    public asInterface(): ICustomOptionsTest {
      return this
    }
  }

  @protobufjs.Type.d('RequiredPropertiesTest')
  export class RequiredPropertiesTest
    extends protobufjs.Message<RequiredPropertiesTest>
    implements ConvertibleTo<IRequiredPropertiesTest> {
    @protobufjs.Field.d(1, 'int32')
    public requiredField!: number

    @protobufjs.Field.d(2, 'int32')
    public customRequiredField!: number

    @protobufjs.Field.d(3, 'int32')
    public optionalField?: number

    public asInterface(): IRequiredPropertiesTest {
      return this
    }
  }
}
