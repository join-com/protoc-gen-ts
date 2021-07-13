// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as joinGRPC from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { Common as Common_Common } from './common/Common'
import { Common as Common_Extra } from './common/Extra'
import { GoogleProtobuf } from './google/protobuf/Timestamp'

import { grpc } from '@join-com/grpc'

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

  @protobufjs.Type.d('foo_CustomOptionsTest')
  export class CustomOptionsTest
    extends protobufjs.Message<CustomOptionsTest>
    implements ConvertibleTo<ICustomOptionsTest>
  {
    @protobufjs.Field.d(1, Common_Extra.ExtraPkgMessage)
    public requiredField!: Common_Extra.ExtraPkgMessage

    @protobufjs.Field.d(2, 'int32', 'optional')
    public typicalOptionalField?: number

    @protobufjs.Field.d(3, 'int32', 'optional')
    public customOptionalField?: number

    public asInterface(): ICustomOptionsTest {
      const message = {
        ...this,
        requiredField: this.requiredField.asInterface(),
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof ICustomOptionsTest] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof ICustomOptionsTest]
        }
      }
      return message
    }

    public static fromInterface(
      this: void,
      value: ICustomOptionsTest
    ): CustomOptionsTest {
      const patchedValue = {
        ...value,
        requiredField: Common_Extra.ExtraPkgMessage.fromInterface(
          value.requiredField
        ),
      }

      return CustomOptionsTest.fromObject(patchedValue)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): ICustomOptionsTest {
      const message = CustomOptionsTest.decode(reader).asInterface()
      for (const fieldName of [
        'requiredField',
      ] as (keyof ICustomOptionsTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(
            `Required field ${fieldName} in CustomOptionsTest is null or undefined`
          )
        }
      }
      return message
    }

    public static encodePatched(
      this: void,
      message: ICustomOptionsTest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      for (const fieldName of [
        'requiredField',
      ] as (keyof ICustomOptionsTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(
            `Required field ${fieldName} in CustomOptionsTest is null or undefined`
          )
        }
      }
      const transformedMessage = CustomOptionsTest.fromInterface(message)
      return CustomOptionsTest.encode(transformedMessage, writer)
    }
  }

  /**
   * @deprecated
   */
  @protobufjs.Type.d('foo_Nested')
  export class Nested
    extends protobufjs.Message<Nested>
    implements ConvertibleTo<INested>, INested
  {
    @protobufjs.Field.d(1, 'string', 'optional')
    public title?: string

    public asInterface(): INested {
      const message = { ...this }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof INested] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof INested]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: INested): Nested {
      return Nested.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): INested {
      return Nested.decode(reader).asInterface()
    }

    public static encodePatched(
      this: void,
      message: INested,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return Nested.encode(message, writer)
    }
  }

  @protobufjs.Type.d('foo_Request')
  export class Request
    extends protobufjs.Message<Request>
    implements ConvertibleTo<IRequest>, IRequest
  {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    public asInterface(): IRequest {
      const message = { ...this }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IRequest] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IRequest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IRequest): Request {
      return Request.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): IRequest {
      return Request.decode(reader).asInterface()
    }

    public static encodePatched(
      this: void,
      message: IRequest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return Request.encode(message, writer)
    }
  }

  @protobufjs.Type.d('foo_RequiredPropertiesTest')
  export class RequiredPropertiesTest
    extends protobufjs.Message<RequiredPropertiesTest>
    implements ConvertibleTo<IRequiredPropertiesTest>, IRequiredPropertiesTest
  {
    @protobufjs.Field.d(1, 'int32')
    public requiredField!: number

    @protobufjs.Field.d(2, 'int32')
    public customRequiredField!: number

    @protobufjs.Field.d(3, 'int32', 'optional')
    public optionalField?: number

    public asInterface(): IRequiredPropertiesTest {
      const message = { ...this }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IRequiredPropertiesTest] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IRequiredPropertiesTest]
        }
      }
      return message
    }

    public static fromInterface(
      this: void,
      value: IRequiredPropertiesTest
    ): RequiredPropertiesTest {
      return RequiredPropertiesTest.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): IRequiredPropertiesTest {
      const message = RequiredPropertiesTest.decode(reader).asInterface()
      for (const fieldName of [
        'requiredField',
        'customRequiredField',
      ] as (keyof IRequiredPropertiesTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(
            `Required field ${fieldName} in RequiredPropertiesTest is null or undefined`
          )
        }
      }
      return message
    }

    public static encodePatched(
      this: void,
      message: IRequiredPropertiesTest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      for (const fieldName of [
        'requiredField',
        'customRequiredField',
      ] as (keyof IRequiredPropertiesTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(
            `Required field ${fieldName} in RequiredPropertiesTest is null or undefined`
          )
        }
      }
      return RequiredPropertiesTest.encode(message, writer)
    }
  }

  @protobufjs.Type.d('foo_Test')
  export class Test
    extends protobufjs.Message<Test>
    implements ConvertibleTo<ITest>
  {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public fieldInt32?: number

    @protobufjs.Field.d(2, 'int32', 'repeated')
    public fieldInt32Repeated?: number[]

    @protobufjs.Field.d(3, 'double', 'optional')
    public fieldDouble?: number

    @protobufjs.Field.d(4, 'double', 'repeated')
    public fieldDoubleRepeated?: number[]

    @protobufjs.Field.d(5, 'float', 'optional')
    public fieldFloat?: number

    @protobufjs.Field.d(6, 'float', 'repeated')
    public fieldFloatRepeated?: number[]

    @protobufjs.Field.d(7, 'uint32', 'optional')
    public fieldUint32?: number

    @protobufjs.Field.d(8, 'uint32', 'repeated')
    public fieldUint32Repeated?: number[]

    @protobufjs.Field.d(9, 'uint64', 'optional')
    public fieldUint64?: number

    @protobufjs.Field.d(10, 'uint64', 'repeated')
    public fieldUint64Repeated?: number[]

    @protobufjs.Field.d(11, 'sint32', 'optional')
    public fieldSint32?: number

    @protobufjs.Field.d(12, 'sint32', 'repeated')
    public fieldSint32Repeated?: number[]

    @protobufjs.Field.d(13, 'sint64', 'optional')
    public fieldSint64?: number

    @protobufjs.Field.d(14, 'sint64', 'repeated')
    public fieldSint64Repeated?: number[]

    @protobufjs.Field.d(15, 'fixed32', 'optional')
    public fieldFixed32?: number

    @protobufjs.Field.d(16, 'fixed32', 'repeated')
    public fieldFixed32Repeated?: number[]

    @protobufjs.Field.d(17, 'fixed64', 'optional')
    public fieldFixed64?: number

    @protobufjs.Field.d(18, 'fixed64', 'repeated')
    public fieldFixed64Repeated?: number[]

    @protobufjs.Field.d(19, 'sfixed32', 'optional')
    public fieldSfixed32?: number

    @protobufjs.Field.d(20, 'sfixed32', 'repeated')
    public fieldSfixed32Repeated?: number[]

    @protobufjs.Field.d(21, 'sfixed64', 'optional')
    public fieldSfixed64?: number

    @protobufjs.Field.d(22, 'sfixed64', 'repeated')
    public fieldSfixed64Repeated?: number[]

    @protobufjs.Field.d(23, 'bool', 'optional')
    public fieldBool?: boolean

    @protobufjs.Field.d(24, 'bool', 'repeated')
    public fieldBoolRepeated?: boolean[]

    @protobufjs.Field.d(25, 'string', 'optional')
    public fieldString?: string

    @protobufjs.Field.d(26, 'string', 'repeated')
    public fieldStringRepeated?: string[]

    @protobufjs.Field.d(27, 'bytes', 'optional')
    public fieldBytes?: Uint8Array

    @protobufjs.Field.d(28, 'bytes', 'repeated')
    public fieldBytesRepeated?: Uint8Array[]

    @protobufjs.Field.d(29, EnumType_Enum, 'optional', '__JoinGrpcUndefined__')
    public fieldEnum?: EnumType_Enum

    @protobufjs.Field.d(30, Role_Enum, 'repeated')
    public fieldEnumRepeated?: Role_Enum[]

    @protobufjs.Field.d(33, Nested, 'optional')
    public message?: Nested

    @protobufjs.Field.d(34, Nested, 'repeated')
    public messageRepeated?: Nested[]

    @protobufjs.Field.d(35, GoogleProtobuf.Timestamp, 'optional')
    public timestamp?: GoogleProtobuf.Timestamp

    @protobufjs.Field.d(36, GoogleProtobuf.Timestamp, 'repeated')
    public timestampRepeated?: GoogleProtobuf.Timestamp[]

    @protobufjs.Field.d(37, Common_Common.OtherPkgMessage, 'optional')
    public otherPkgMessage?: Common_Common.OtherPkgMessage

    @protobufjs.Field.d(38, Common_Common.OtherPkgMessage, 'repeated')
    public otherPkgMessageRepeated?: Common_Common.OtherPkgMessage[]

    @protobufjs.Field.d(39, 'int64', 'optional')
    public fieldInt64?: number

    @protobufjs.Field.d(40, 'int64', 'repeated')
    public fieldInt64Repeated?: number[]

    public asInterface(): ITest {
      const message = {
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
      for (const fieldName of Object.keys(message)) {
        const field = message[fieldName as keyof ITest]
        if (field == null || (Array.isArray(field) && field.length === 0)) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof ITest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: ITest): Test {
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

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): ITest {
      return Test.decode(reader).asInterface()
    }

    public static encodePatched(
      this: void,
      message: ITest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      const transformedMessage = Test.fromInterface(message)
      return Test.encode(transformedMessage, writer)
    }
  }

  export interface IUsersServiceImplementation {
    Find: grpc.handleUnaryCall<IRequest, Common_Common.IOtherPkgMessage>
    FindBidiStream: grpc.handleBidiStreamingCall<
      IRequest,
      Common_Common.IOtherPkgMessage
    >
    FindClientStream: grpc.handleClientStreamingCall<
      IRequest,
      Common_Common.IOtherPkgMessage
    >
    FindServerStream: grpc.handleServerStreamingCall<
      IRequest,
      Common_Common.IOtherPkgMessage
    >
  }

  export const usersServiceDefinition: grpc.ServiceDefinition<IUsersServiceImplementation> =
    {
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
    }

  export abstract class AbstractUsersService extends joinGRPC.Service<IUsersServiceImplementation> {
    constructor(
      protected readonly logger?: joinGRPC.INoDebugLogger,
      trace?: joinGRPC.IServiceTrace
    ) {
      super(
        usersServiceDefinition,
        {
          find: (call) => this.Find(call),
          findBidiStream: (call) => this.FindBidiStream(call),
          findClientStream: (call) => this.FindClientStream(call),
          findServerStream: (call) => this.FindServerStream(call),
        },
        logger,
        trace
      )
    }

    public abstract Find(
      call: grpc.ServerUnaryCall<IRequest, Common_Common.IOtherPkgMessage>
    ): Promise<Common_Common.IOtherPkgMessage>
    public abstract FindBidiStream(
      call: grpc.ServerDuplexStream<IRequest, Common_Common.IOtherPkgMessage>
    ): Promise<void>
    public abstract FindClientStream(
      call: grpc.ServerReadableStream<IRequest, Common_Common.IOtherPkgMessage>
    ): Promise<Common_Common.IOtherPkgMessage>
    public abstract FindServerStream(
      call: grpc.ServerWritableStream<IRequest, Common_Common.IOtherPkgMessage>
    ): Promise<void>
  }

  export interface IUsersClient
    extends joinGRPC.IExtendedClient<IUsersServiceImplementation, 'foo.Users'> {
    /**
     * @deprecated
     */
    find(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IUnaryRequest<Common_Common.IOtherPkgMessage>

    findBidiStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IBidiStreamRequest<IRequest, Common_Common.IOtherPkgMessage>

    findClientStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IClientStreamRequest<IRequest, Common_Common.IOtherPkgMessage>

    findServerStream(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IServerStreamRequest<Common_Common.IOtherPkgMessage>
  }

  export class UsersClient
    extends joinGRPC.Client<IUsersServiceImplementation, 'foo.Users'>
    implements IUsersClient
  {
    constructor(
      config: joinGRPC.ISimplifiedClientConfig<IUsersServiceImplementation>
    ) {
      super(
        {
          ...config,
          serviceDefinition: usersServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'foo.Users'
      )
    }

    /**
     * @deprecated
     */
    public find(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IUnaryRequest<Common_Common.IOtherPkgMessage> {
      this.logger?.warn("using deprecated service method 'UsersClient.find'")
      return this.makeUnaryRequest('Find', request, metadata, options)
    }

    public findBidiStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IBidiStreamRequest<IRequest, Common_Common.IOtherPkgMessage> {
      return this.makeBidiStreamRequest('FindBidiStream', metadata, options)
    }

    public findClientStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IClientStreamRequest<IRequest, Common_Common.IOtherPkgMessage> {
      return this.makeClientStreamRequest('FindClientStream', metadata, options)
    }

    public findServerStream(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IServerStreamRequest<Common_Common.IOtherPkgMessage> {
      return this.makeServerStreamRequest(
        'FindServerStream',
        request,
        metadata,
        options
      )
    }
  }
}
