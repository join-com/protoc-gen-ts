// GENERATED CODE -- DO NOT EDIT!

// import * as nodeTrace from '@join-com/node-trace'
import * as joinGRPC from '@join-com/grpc'
import { grpc } from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from './google/protobuf/Timestamp'
import { Common as Common_Common } from './common/Common'
import { Common as Common_Extra } from './common/Extra'

// eslint-disable-next-line @typescript-eslint/no-namespace
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
    timestamp?: Date
    timestampRepeated?: Date[]
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
    implements ConvertibleTo<IRequest>, IRequest {
    @protobufjs.Field.d(1, 'int32')
    public id?: number

    public asInterface(): IRequest {
      return this
    }

    public static fromInterface(value: IRequest): Request {
      return Request.fromObject(value)
    }

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): IRequest {
      return Request.decode(reader)
    }

    public static encodePatched(
      message: IRequest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return Request.encode(message, writer)
    }
  }

  /**
   * @deprecated
   */
  @protobufjs.Type.d('Nested')
  export class Nested
    extends protobufjs.Message<Nested>
    implements ConvertibleTo<INested>, INested {
    @protobufjs.Field.d(1, 'string')
    public title?: string

    public asInterface(): INested {
      return this
    }

    public static fromInterface(value: INested): Nested {
      return Nested.fromObject(value)
    }

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): INested {
      return Nested.decode(reader)
    }

    public static encodePatched(
      message: INested,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return Nested.encode(message, writer)
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
        ...this,
        fieldEnum: (this.fieldEnum != null
          ? EnumType_Enum[this.fieldEnum]!
          : undefined) as EnumType | undefined,
        fieldEnumRepeated: this.fieldEnumRepeated?.map(
          (e) => Role_Enum[e]! as Role
        ),
        timestamp:
          this.timestamp != null
            ? new Date(
                (this.timestamp.seconds ?? 0) * 1000 +
                  (this.timestamp.nanos ?? 0) / 1000000
              )
            : undefined,
        timestampRepeated: this.timestampRepeated?.map(
          (ts) => new Date((ts.seconds ?? 0) * 1000 + (ts.nanos ?? 0) / 1000000)
        ),
      }
    }

    public static fromInterface(value: ITest): Test {
      const patchedValue = {
        ...value,
        fieldEnum: (value.fieldEnum != null
          ? EnumType_Enum[value.fieldEnum]!
          : undefined) as EnumType_Enum | undefined,
        fieldEnumRepeated: value.fieldEnumRepeated?.map((e) => Role_Enum[e]!),
        timestamp:
          value.timestamp != null
            ? GoogleProtobuf.Timestamp.fromInterface({
                seconds: Math.floor(value.timestamp.getTime() / 1000),
                nanos: value.timestamp.getMilliseconds() * 1000000,
              })
            : undefined,
        timestampRepeated: value.timestampRepeated?.map((d) =>
          GoogleProtobuf.Timestamp.fromInterface({
            seconds: Math.floor(d.getTime() / 1000),
            nanos: d.getMilliseconds() * 1000000,
          })
        ),
      }

      return Test.fromObject(patchedValue)
    }

    public static decodePatched(reader: protobufjs.Reader | Uint8Array): ITest {
      return Test.decode(reader).asInterface()
    }

    public static encodePatched(
      message: ITest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      const transformedMessage = Test.fromInterface(message)
      return Test.encode(transformedMessage, writer)
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
      return {
        ...this,
        requiredField: this.requiredField.asInterface(),
      }
    }

    public static fromInterface(value: ICustomOptionsTest): CustomOptionsTest {
      const patchedValue = {
        ...value,
        requiredField: Common_Extra.ExtraPkgMessage.fromInterface(
          value.requiredField
        ),
      }

      return CustomOptionsTest.fromObject(patchedValue)
    }

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): ICustomOptionsTest {
      return CustomOptionsTest.decode(reader).asInterface()
    }

    public static encodePatched(
      message: ICustomOptionsTest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      const transformedMessage = CustomOptionsTest.fromInterface(message)
      return CustomOptionsTest.encode(transformedMessage, writer)
    }
  }

  @protobufjs.Type.d('RequiredPropertiesTest')
  export class RequiredPropertiesTest
    extends protobufjs.Message<RequiredPropertiesTest>
    implements ConvertibleTo<IRequiredPropertiesTest>, IRequiredPropertiesTest {
    @protobufjs.Field.d(1, 'int32')
    public requiredField!: number

    @protobufjs.Field.d(2, 'int32')
    public customRequiredField!: number

    @protobufjs.Field.d(3, 'int32')
    public optionalField?: number

    public asInterface(): IRequiredPropertiesTest {
      return this
    }

    public static fromInterface(
      value: IRequiredPropertiesTest
    ): RequiredPropertiesTest {
      return RequiredPropertiesTest.fromObject(value)
    }

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): IRequiredPropertiesTest {
      return RequiredPropertiesTest.decode(reader)
    }

    public static encodePatched(
      message: IRequiredPropertiesTest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return RequiredPropertiesTest.encode(message, writer)
    }
  }

  export interface IUsersServiceImplementation
    extends grpc.UntypedServiceImplementation {
    Find: grpc.handleUnaryCall<IRequest, Common_Common.IOtherPkgMessage>
    FindClientStream: grpc.handleClientStreamingCall<
      IRequest,
      Common_Common.IOtherPkgMessage
    >
    FindServerStream: grpc.handleServerStreamingCall<
      IRequest,
      Common_Common.IOtherPkgMessage
    >
    FindBidiStream: grpc.handleClientStreamingCall<
      IRequest,
      Common_Common.IOtherPkgMessage
    >
  }

  export const usersServiceDefinition: grpc.ServiceDefinition<IUsersServiceImplementation> = {
    Find: {
      path: '/foo.Users/Find',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IRequest) =>
        Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(
          response
        ).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
    FindClientStream: {
      path: '/foo.Users/FindClientStream',
      requestStream: true,
      responseStream: false,
      requestSerialize: (request: IRequest) =>
        Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(
          response
        ).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
    FindServerStream: {
      path: '/foo.Users/FindServerStream',
      requestStream: false,
      responseStream: true,
      requestSerialize: (request: IRequest) =>
        Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(
          response
        ).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
    FindBidiStream: {
      path: '/foo.Users/FindBidiStream',
      requestStream: true,
      responseStream: true,
      requestSerialize: (request: IRequest) =>
        Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(
          response
        ).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
  }
}
