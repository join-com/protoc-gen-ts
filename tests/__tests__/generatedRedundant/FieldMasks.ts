// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.3.0.9b92372.1728131327
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as joinGRPC from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from './google/protobuf/FieldMask'
import { grpc } from '@join-com/grpc'
import { root } from './root'

protobufjs.roots['decorated'] = root

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace FieldMasks {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  interface IFieldMask<T> {
    paths?: (keyof T)[] | string[]
  }

  export interface IGetRequest {
    id?: number
    readMask?: IFieldMask<IUser>
  }

  export interface IUser {
    id?: number
    name?: string
    email?: string
  }

  @protobufjs.Type.d('field_masks_GetRequest')
  export class GetRequest extends protobufjs.Message<GetRequest> implements ConvertibleTo<IGetRequest>, IGetRequest {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    @protobufjs.Field.d(2, GoogleProtobuf.FieldMask, 'optional')
    public readMask?: GoogleProtobuf.FieldMask

    public asInterface(): IGetRequest {
      const message = {
        ...this,
        readMask: this.readMask?.asInterface(),
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IGetRequest] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IGetRequest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IGetRequest): GetRequest {
      return GetRequest.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IGetRequest {
      return GetRequest.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IGetRequest, writer?: protobufjs.Writer): protobufjs.Writer {
      return GetRequest.encode(message, writer)
    }
  }

  @protobufjs.Type.d('field_masks_User')
  export class User extends protobufjs.Message<User> implements ConvertibleTo<IUser>, IUser {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    @protobufjs.Field.d(2, 'string', 'optional')
    public name?: string

    @protobufjs.Field.d(3, 'string', 'optional')
    public email?: string

    public asInterface(): IUser {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IUser] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IUser]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IUser): User {
      return User.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IUser {
      return User.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IUser, writer?: protobufjs.Writer): protobufjs.Writer {
      return User.encode(message, writer)
    }
  }

  export interface IUsersServiceImplementation {
    Get: grpc.handleUnaryCall<IGetRequest, IUser>
  }

  export const usersServiceDefinition: grpc.ServiceDefinition<IUsersServiceImplementation> = {
    Get: {
      path: '/field_masks.Users/Get',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IGetRequest) => GetRequest.encodePatched(request).finish() as Buffer,
      requestDeserialize: GetRequest.decodePatched,
      responseSerialize: (response: IUser) => User.encodePatched(response).finish() as Buffer,
      responseDeserialize: User.decodePatched,
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
          get: call => this.Get(call),
        },
        logger,
        errorHandler,
      )
    }

    public abstract Get(call: grpc.ServerUnaryCall<IGetRequest, IUser>): Promise<IUser>
  }

  export interface IUsersClient extends joinGRPC.IExtendedClient<IUsersServiceImplementation, 'field_masks.Users'> {
    get(
      request: IGetRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IUser>
  }

  export class UsersClient
    extends joinGRPC.Client<IUsersServiceImplementation, 'field_masks.Users'>
    implements IUsersClient
  {
    constructor(config: joinGRPC.ISimplifiedClientConfig<IUsersServiceImplementation>) {
      super(
        {
          ...config,
          serviceDefinition: usersServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'field_masks.Users',
      )
    }

    public get(
      request: IGetRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IUser> {
      return this.makeUnaryRequest('Get', request, metadata, options)
    }
  }
}
