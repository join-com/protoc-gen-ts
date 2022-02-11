// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.f1fa0f2.1644588877
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as joinGRPC from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { WithFlavor } from '@coderspirit/nominal'

import { grpc } from '@join-com/grpc'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Flavors {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export type Email = WithFlavor<string, 'Email'>
  export type UserId = WithFlavor<number, 'UserId'>

  export interface IUserRequest {
    userId?: UserId
  }

  export interface IUserProfile {
    id: UserId
    username: string
    emails: Email[]
  }

  @protobufjs.Type.d('flavors_UserProfile')
  export class UserProfile
    extends protobufjs.Message<UserProfile>
    implements ConvertibleTo<IUserProfile>, IUserProfile
  {
    @protobufjs.Field.d(1, 'int32')
    public id!: UserId

    @protobufjs.Field.d(2, 'string')
    public username!: string

    @protobufjs.Field.d(3, 'string', 'repeated')
    public emails!: Email[]

    public asInterface(): IUserProfile {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        const field = message[fieldName as keyof IUserProfile]
        if (field == null || (Array.isArray(field) && field.length === 0)) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IUserProfile]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IUserProfile): UserProfile {
      return UserProfile.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IUserProfile {
      const message = UserProfile.decode(reader).asInterface()
      for (const fieldName of ['id', 'username', 'emails'] as (keyof IUserProfile)[]) {
        const field = message[fieldName]
        if (field == null || (Array.isArray(field) && field.length === 0)) {
          throw new Error(`Required field ${fieldName} in UserProfile is null or undefined`)
        }
      }
      return message
    }

    public static encodePatched(this: void, message: IUserProfile, writer?: protobufjs.Writer): protobufjs.Writer {
      for (const fieldName of ['id', 'username', 'emails'] as (keyof IUserProfile)[]) {
        const field = message[fieldName]
        if (field == null || (Array.isArray(field) && field.length === 0)) {
          throw new Error(`Required field ${fieldName} in UserProfile is null or undefined`)
        }
      }
      return UserProfile.encode(message, writer)
    }
  }

  @protobufjs.Type.d('flavors_UserRequest')
  export class UserRequest
    extends protobufjs.Message<UserRequest>
    implements ConvertibleTo<IUserRequest>, IUserRequest
  {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public userId?: UserId

    public asInterface(): IUserRequest {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IUserRequest] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IUserRequest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IUserRequest): UserRequest {
      return UserRequest.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IUserRequest {
      return UserRequest.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IUserRequest, writer?: protobufjs.Writer): protobufjs.Writer {
      return UserRequest.encode(message, writer)
    }
  }

  export interface IUsersServiceImplementation {
    Find: grpc.handleUnaryCall<IUserRequest, IUserProfile>
  }

  export const usersServiceDefinition: grpc.ServiceDefinition<IUsersServiceImplementation> = {
    Find: {
      path: '/flavors.Users/Find',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IUserRequest) => UserRequest.encodePatched(request).finish() as Buffer,
      requestDeserialize: UserRequest.decodePatched,
      responseSerialize: (response: IUserProfile) => UserProfile.encodePatched(response).finish() as Buffer,
      responseDeserialize: UserProfile.decodePatched,
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
        },
        logger,
        errorHandler,
      )
    }

    public abstract Find(call: grpc.ServerUnaryCall<IUserRequest, IUserProfile>): Promise<IUserProfile>
  }

  export interface IUsersClient extends joinGRPC.IExtendedClient<IUsersServiceImplementation, 'flavors.Users'> {
    find(
      request: IUserRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IUserProfile>
  }

  export class UsersClient
    extends joinGRPC.Client<IUsersServiceImplementation, 'flavors.Users'>
    implements IUsersClient
  {
    constructor(config: joinGRPC.ISimplifiedClientConfig<IUsersServiceImplementation>) {
      super(
        {
          ...config,
          serviceDefinition: usersServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'flavors.Users',
      )
    }

    public find(
      request: IUserRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IUserProfile> {
      return this.makeUnaryRequest('Find', request, metadata, options)
    }
  }
}
