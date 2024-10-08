// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.3.1.cbb8b0b.1728390468
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as joinGRPC from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { Common as Common_Common } from './common/Common'
import { Common as Common_Extra } from './common/Extra'
import { GoogleProtobuf as GoogleProtobuf_Empty } from './google/protobuf/Empty'
import { GoogleProtobuf as GoogleProtobuf_Timestamp } from './google/protobuf/Timestamp'
import { grpc } from '@join-com/grpc'
import { root } from './root'

protobufjs.roots['decorated'] = root

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
    /**
     * @deprecated
     */
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

  export interface IBigWrapper {
    nestedTest?: ITest
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
    repeatedOptionalField?: number[]
    requiredEnumField: EnumType
  }

  export interface IOneRequiredFieldTest {
    singleRequiredField: number
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

    public static fromInterface(this: void, value: ICustomOptionsTest): CustomOptionsTest {
      const patchedValue = {
        ...value,
        requiredField: Common_Extra.ExtraPkgMessage.fromInterface(value.requiredField),
      }

      return CustomOptionsTest.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): ICustomOptionsTest {
      const message = CustomOptionsTest.decode(reader).asInterface()
      for (const fieldName of ['requiredField'] as (keyof ICustomOptionsTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(`Required field ${fieldName} in CustomOptionsTest is null or undefined`)
        }
      }
      return message
    }

    public static encodePatched(
      this: void,
      message: ICustomOptionsTest,
      writer?: protobufjs.Writer,
    ): protobufjs.Writer {
      for (const fieldName of ['requiredField'] as (keyof ICustomOptionsTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(`Required field ${fieldName} in CustomOptionsTest is null or undefined`)
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
  export class Nested extends protobufjs.Message<Nested> implements ConvertibleTo<INested>, INested {
    @protobufjs.Field.d(1, 'string', 'optional')
    public title?: string

    public asInterface(): INested {
      const message = {
        ...this,
      }
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

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): INested {
      return Nested.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: INested, writer?: protobufjs.Writer): protobufjs.Writer {
      return Nested.encode(message, writer)
    }
  }

  @protobufjs.Type.d('foo_OneRequiredFieldTest')
  export class OneRequiredFieldTest
    extends protobufjs.Message<OneRequiredFieldTest>
    implements ConvertibleTo<IOneRequiredFieldTest>, IOneRequiredFieldTest
  {
    @protobufjs.Field.d(1, 'int32')
    public singleRequiredField!: number

    public asInterface(): IOneRequiredFieldTest {
      const message = {
        ...this,
      }
      return message
    }

    public static fromInterface(this: void, value: IOneRequiredFieldTest): OneRequiredFieldTest {
      return OneRequiredFieldTest.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IOneRequiredFieldTest {
      const message = OneRequiredFieldTest.decode(reader).asInterface()
      for (const fieldName of ['singleRequiredField'] as (keyof IOneRequiredFieldTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(`Required field ${fieldName} in OneRequiredFieldTest is null or undefined`)
        }
      }
      return message
    }

    public static encodePatched(
      this: void,
      message: IOneRequiredFieldTest,
      writer?: protobufjs.Writer,
    ): protobufjs.Writer {
      for (const fieldName of ['singleRequiredField'] as (keyof IOneRequiredFieldTest)[]) {
        if (message[fieldName] == null) {
          throw new Error(`Required field ${fieldName} in OneRequiredFieldTest is null or undefined`)
        }
      }
      return OneRequiredFieldTest.encode(message, writer)
    }
  }

  @protobufjs.Type.d('foo_Request')
  export class Request extends protobufjs.Message<Request> implements ConvertibleTo<IRequest>, IRequest {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    public asInterface(): IRequest {
      const message = {
        ...this,
      }
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

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IRequest {
      return Request.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IRequest, writer?: protobufjs.Writer): protobufjs.Writer {
      return Request.encode(message, writer)
    }
  }

  @protobufjs.Type.d('foo_RequiredPropertiesTest')
  export class RequiredPropertiesTest
    extends protobufjs.Message<RequiredPropertiesTest>
    implements ConvertibleTo<IRequiredPropertiesTest>
  {
    @protobufjs.Field.d(1, 'int32')
    public requiredField!: number

    @protobufjs.Field.d(2, 'int32')
    public customRequiredField!: number

    @protobufjs.Field.d(3, 'int32', 'optional')
    public optionalField?: number

    @protobufjs.Field.d(4, 'int32', 'repeated')
    public repeatedOptionalField?: number[]

    @protobufjs.Field.d(5, EnumType_Enum)
    public requiredEnumField!: EnumType_Enum

    public asInterface(): IRequiredPropertiesTest {
      const message = {
        ...this,
        requiredEnumField: EnumType_Enum[this.requiredEnumField]! as EnumType,
      }
      for (const fieldName of Object.keys(message)) {
        const field = message[fieldName as keyof IRequiredPropertiesTest]
        if (field == null || (Array.isArray(field) && (field as any[]).length === 0)) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IRequiredPropertiesTest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IRequiredPropertiesTest): RequiredPropertiesTest {
      const patchedValue = {
        ...value,
        requiredEnumField: EnumType_Enum[value.requiredEnumField]!,
      }

      return RequiredPropertiesTest.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IRequiredPropertiesTest {
      const message = RequiredPropertiesTest.decode(reader).asInterface()
      for (const fieldName of [
        'requiredField',
        'customRequiredField',
        'requiredEnumField',
      ] as (keyof IRequiredPropertiesTest)[]) {
        const field = message[fieldName]
        if (field == null || (Array.isArray(field) && (field as any[]).length === 0)) {
          throw new Error(`Required field ${fieldName} in RequiredPropertiesTest is null or undefined`)
        }
      }
      return message
    }

    public static encodePatched(
      this: void,
      message: IRequiredPropertiesTest,
      writer?: protobufjs.Writer,
    ): protobufjs.Writer {
      for (const fieldName of [
        'requiredField',
        'customRequiredField',
        'requiredEnumField',
      ] as (keyof IRequiredPropertiesTest)[]) {
        const field = message[fieldName]
        if (field == null || (Array.isArray(field) && (field as any[]).length === 0)) {
          throw new Error(`Required field ${fieldName} in RequiredPropertiesTest is null or undefined`)
        }
      }
      const transformedMessage = RequiredPropertiesTest.fromInterface(message)
      return RequiredPropertiesTest.encode(transformedMessage, writer)
    }
  }

  @protobufjs.Type.d('foo_Test')
  export class Test extends protobufjs.Message<Test> implements ConvertibleTo<ITest> {
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

    /**
     * @deprecated
     */
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

    @protobufjs.Field.d(35, GoogleProtobuf_Timestamp.Timestamp, 'optional')
    public timestamp?: GoogleProtobuf_Timestamp.Timestamp

    @protobufjs.Field.d(36, GoogleProtobuf_Timestamp.Timestamp, 'repeated')
    public timestampRepeated?: GoogleProtobuf_Timestamp.Timestamp[]

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
        fieldEnum: (this.fieldEnum != null ? EnumType_Enum[this.fieldEnum]! : undefined) as EnumType | undefined,
        fieldEnumRepeated: this.fieldEnumRepeated?.map(e => Role_Enum[e]! as Role),
        message: this.message?.asInterface(),
        messageRepeated: this.messageRepeated?.map(o => o.asInterface()),
        timestamp:
          this.timestamp != null
            ? new Date((this.timestamp.seconds ?? 0) * 1000 + (this.timestamp.nanos ?? 0) / 1000000)
            : undefined,
        timestampRepeated: this.timestampRepeated?.map(
          ts => new Date((ts.seconds ?? 0) * 1000 + (ts.nanos ?? 0) / 1000000),
        ),
        otherPkgMessage: this.otherPkgMessage?.asInterface(),
        otherPkgMessageRepeated: this.otherPkgMessageRepeated?.map(o => o.asInterface()),
      }
      for (const fieldName of Object.keys(message)) {
        const field = message[fieldName as keyof ITest]
        if (field == null || (Array.isArray(field) && (field as any[]).length === 0)) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof ITest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: ITest): Test {
      const patchedValue = {
        ...value,
        fieldEnum: (value.fieldEnum != null ? EnumType_Enum[value.fieldEnum]! : undefined) as EnumType_Enum | undefined,
        fieldEnumRepeated: value.fieldEnumRepeated?.map(e => Role_Enum[e]!),
        timestamp:
          value.timestamp != null
            ? GoogleProtobuf_Timestamp.Timestamp.fromInterface({
                seconds: Math.floor(value.timestamp.getTime() / 1000),
                nanos: value.timestamp.getMilliseconds() * 1000000,
              })
            : undefined,
        timestampRepeated: value.timestampRepeated?.map(d =>
          GoogleProtobuf_Timestamp.Timestamp.fromInterface({
            seconds: Math.floor(d.getTime() / 1000),
            nanos: d.getMilliseconds() * 1000000,
          }),
        ),
      }

      return Test.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): ITest {
      return Test.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: ITest, writer?: protobufjs.Writer): protobufjs.Writer {
      const transformedMessage = Test.fromInterface(message)
      return Test.encode(transformedMessage, writer)
    }
  }

  @protobufjs.Type.d('foo_BigWrapper')
  export class BigWrapper extends protobufjs.Message<BigWrapper> implements ConvertibleTo<IBigWrapper> {
    @protobufjs.Field.d(1, Test, 'optional')
    public nestedTest?: Test

    public asInterface(): IBigWrapper {
      const message = {
        ...this,
        nestedTest: this.nestedTest?.asInterface(),
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IBigWrapper] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IBigWrapper]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IBigWrapper): BigWrapper {
      const patchedValue = {
        ...value,
        nestedTest: value.nestedTest != null ? Test.fromInterface(value.nestedTest) : undefined,
      }

      return BigWrapper.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IBigWrapper {
      return BigWrapper.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IBigWrapper, writer?: protobufjs.Writer): protobufjs.Writer {
      const transformedMessage = BigWrapper.fromInterface(message)
      return BigWrapper.encode(transformedMessage, writer)
    }
  }

  export interface IUsersServiceImplementation {
    Find: grpc.handleUnaryCall<IRequest, Common_Common.IOtherPkgMessage>
    FindBidiStream: grpc.handleBidiStreamingCall<IRequest, Common_Common.IOtherPkgMessage>
    FindClientStream: grpc.handleClientStreamingCall<IRequest, Common_Common.IOtherPkgMessage>
    FindServerStream: grpc.handleServerStreamingCall<IRequest, Common_Common.IOtherPkgMessage>
  }

  export const usersServiceDefinition: grpc.ServiceDefinition<IUsersServiceImplementation> = {
    Find: {
      path: '/foo.Users/Find',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IRequest) => Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(response).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
    FindBidiStream: {
      path: '/foo.Users/FindBidiStream',
      requestStream: true,
      responseStream: true,
      requestSerialize: (request: IRequest) => Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(response).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
    FindClientStream: {
      path: '/foo.Users/FindClientStream',
      requestStream: true,
      responseStream: false,
      requestSerialize: (request: IRequest) => Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(response).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
    FindServerStream: {
      path: '/foo.Users/FindServerStream',
      requestStream: false,
      responseStream: true,
      requestSerialize: (request: IRequest) => Request.encodePatched(request).finish() as Buffer,
      requestDeserialize: Request.decodePatched,
      responseSerialize: (response: Common_Common.IOtherPkgMessage) =>
        Common_Common.OtherPkgMessage.encodePatched(response).finish() as Buffer,
      responseDeserialize: Common_Common.OtherPkgMessage.decodePatched,
    },
  }

  export abstract class AbstractUsersService extends joinGRPC.Service<IUsersServiceImplementation> {
    constructor(
      protected readonly logger?: joinGRPC.INoDebugLogger,
      protected readonly errorHandler?: joinGRPC.IServiceErrorHandler,
    ) {
      super(
        usersServiceDefinition,
        {
          find: call => this.Find(call),
          findBidiStream: call => this.FindBidiStream(call),
          findClientStream: call => this.FindClientStream(call),
          findServerStream: call => this.FindServerStream(call),
        },
        logger,
        errorHandler,
      )
    }

    public abstract Find(
      call: grpc.ServerUnaryCall<IRequest, Common_Common.IOtherPkgMessage>,
    ): Promise<Common_Common.IOtherPkgMessage>
    public abstract FindBidiStream(
      call: grpc.ServerDuplexStream<IRequest, Common_Common.IOtherPkgMessage>,
    ): Promise<void>
    public abstract FindClientStream(
      call: grpc.ServerReadableStream<IRequest, Common_Common.IOtherPkgMessage>,
    ): Promise<Common_Common.IOtherPkgMessage>
    public abstract FindServerStream(
      call: grpc.ServerWritableStream<IRequest, Common_Common.IOtherPkgMessage>,
    ): Promise<void>
  }

  export interface ISimpleTestServiceImplementation {
    ForwardParameter: grpc.handleUnaryCall<IBigWrapper, IBigWrapper>
    GetEmptyResult: grpc.handleUnaryCall<GoogleProtobuf_Empty.IEmpty, IBigWrapper>
  }

  export const simpleTestServiceDefinition: grpc.ServiceDefinition<ISimpleTestServiceImplementation> = {
    ForwardParameter: {
      path: '/foo.SimpleTest/ForwardParameter',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IBigWrapper) => BigWrapper.encodePatched(request).finish() as Buffer,
      requestDeserialize: BigWrapper.decodePatched,
      responseSerialize: (response: IBigWrapper) => BigWrapper.encodePatched(response).finish() as Buffer,
      responseDeserialize: BigWrapper.decodePatched,
    },
    GetEmptyResult: {
      path: '/foo.SimpleTest/GetEmptyResult',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: GoogleProtobuf_Empty.IEmpty) =>
        GoogleProtobuf_Empty.Empty.encodePatched(request).finish() as Buffer,
      requestDeserialize: GoogleProtobuf_Empty.Empty.decodePatched,
      responseSerialize: (response: IBigWrapper) => BigWrapper.encodePatched(response).finish() as Buffer,
      responseDeserialize: BigWrapper.decodePatched,
    },
  }

  export abstract class AbstractSimpleTestService extends joinGRPC.Service<ISimpleTestServiceImplementation> {
    constructor(
      protected readonly logger?: joinGRPC.INoDebugLogger,
      protected readonly errorHandler?: joinGRPC.IServiceErrorHandler,
    ) {
      super(
        simpleTestServiceDefinition,
        {
          forwardParameter: call => this.ForwardParameter(call),
          getEmptyResult: call => this.GetEmptyResult(call),
        },
        logger,
        errorHandler,
      )
    }

    public abstract ForwardParameter(call: grpc.ServerUnaryCall<IBigWrapper, IBigWrapper>): Promise<IBigWrapper>
    public abstract GetEmptyResult(
      call: grpc.ServerUnaryCall<GoogleProtobuf_Empty.IEmpty, IBigWrapper>,
    ): Promise<IBigWrapper>
  }

  export interface IUsersClient extends joinGRPC.IExtendedClient<IUsersServiceImplementation, 'foo.Users'> {
    /**
     * @deprecated
     */
    find(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<Common_Common.IOtherPkgMessage>

    findBidiStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IBidiStreamRequest<IRequest, Common_Common.IOtherPkgMessage>

    findClientStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IClientStreamRequest<IRequest, Common_Common.IOtherPkgMessage>

    findServerStream(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IServerStreamRequest<Common_Common.IOtherPkgMessage>
  }

  export class UsersClient extends joinGRPC.Client<IUsersServiceImplementation, 'foo.Users'> implements IUsersClient {
    constructor(config: joinGRPC.ISimplifiedClientConfig<IUsersServiceImplementation>) {
      super(
        {
          ...config,
          serviceDefinition: usersServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'foo.Users',
      )
    }

    /**
     * @deprecated
     */
    public find(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<Common_Common.IOtherPkgMessage> {
      this.logger?.warn("using deprecated service method 'UsersClient.find'")
      return this.makeUnaryRequest('Find', request, metadata, options)
    }

    public findBidiStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IBidiStreamRequest<IRequest, Common_Common.IOtherPkgMessage> {
      return this.makeBidiStreamRequest('FindBidiStream', metadata, options)
    }

    public findClientStream(
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IClientStreamRequest<IRequest, Common_Common.IOtherPkgMessage> {
      return this.makeClientStreamRequest('FindClientStream', metadata, options)
    }

    public findServerStream(
      request: IRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IServerStreamRequest<Common_Common.IOtherPkgMessage> {
      return this.makeServerStreamRequest('FindServerStream', request, metadata, options)
    }
  }

  export interface ISimpleTestClient
    extends joinGRPC.IExtendedClient<ISimpleTestServiceImplementation, 'foo.SimpleTest'> {
    forwardParameter(
      request: IBigWrapper,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IBigWrapper>

    getEmptyResult(
      request: GoogleProtobuf_Empty.IEmpty,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IBigWrapper>
  }

  export class SimpleTestClient
    extends joinGRPC.Client<ISimpleTestServiceImplementation, 'foo.SimpleTest'>
    implements ISimpleTestClient
  {
    constructor(config: joinGRPC.ISimplifiedClientConfig<ISimpleTestServiceImplementation>) {
      super(
        {
          ...config,
          serviceDefinition: simpleTestServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'foo.SimpleTest',
      )
    }

    public forwardParameter(
      request: IBigWrapper,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IBigWrapper> {
      return this.makeUnaryRequest('ForwardParameter', request, metadata, options)
    }

    public getEmptyResult(
      request: GoogleProtobuf_Empty.IEmpty,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IBigWrapper> {
      return this.makeUnaryRequest('GetEmptyResult', request, metadata, options)
    }
  }
}
