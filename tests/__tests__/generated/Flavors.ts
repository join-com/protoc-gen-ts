// GENERATED CODE -- DO NOT EDIT!
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

  @protobufjs.Type.d('UserProfile')
  export class UserProfile
    extends protobufjs.Message<UserProfile>
    implements ConvertibleTo<IUserProfile>, IUserProfile {
    @protobufjs.Field.d(1, 'int32')
    public id!: UserId

    @protobufjs.Field.d(2, 'string')
    public username!: string

    @protobufjs.Field.d(3, 'string', 'repeated')
    public emails!: Email[]

    public asInterface(): IUserProfile {
      return this
    }

    public static fromInterface(this: void, value: IUserProfile): UserProfile {
      return UserProfile.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): IUserProfile {
      const message = UserProfile.decode(reader)
      for (const fieldName of [
        'id',
        'username',
        'emails',
      ] as (keyof UserProfile)[]) {
        if (message[fieldName] == null) {
          throw new Error(
            `Required field ${fieldName} in UserProfile is null or undefined`
          )
        }
      }
      return message
    }

    public static encodePatched(
      this: void,
      message: IUserProfile,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return UserProfile.encode(message, writer)
    }
  }

  @protobufjs.Type.d('UserRequest')
  export class UserRequest
    extends protobufjs.Message<UserRequest>
    implements ConvertibleTo<IUserRequest>, IUserRequest {
    @protobufjs.Field.d(1, 'int32')
    public userId?: UserId

    public asInterface(): IUserRequest {
      return this
    }

    public static fromInterface(this: void, value: IUserRequest): UserRequest {
      return UserRequest.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): IUserRequest {
      return UserRequest.decode(reader)
    }

    public static encodePatched(
      this: void,
      message: IUserRequest,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
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
      requestSerialize: (request: IUserRequest) =>
        UserRequest.encodePatched(request).finish() as Buffer,
      requestDeserialize: UserRequest.decodePatched,
      responseSerialize: (response: IUserProfile) =>
        UserProfile.encodePatched(response).finish() as Buffer,
      responseDeserialize: UserProfile.decodePatched,
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
          Find: (call) => this.Find(call),
        },
        logger,
        trace
      )
    }

    protected abstract Find(
      call: grpc.ServerUnaryCall<IUserRequest, IUserProfile>
    ): Promise<IUserProfile>
  }

  export interface IUsersClient
    extends joinGRPC.IExtendedClient<
      IUsersServiceImplementation,
      'flavors.Users'
    > {
    Find(
      request: IUserRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IUnaryRequest<IUserProfile>
  }

  export class UsersClient
    extends joinGRPC.Client<IUsersServiceImplementation, 'flavors.Users'>
    implements IUsersClient {
    constructor(
      public readonly config: joinGRPC.IClientConfig<IUsersServiceImplementation>
    ) {
      super(config, 'flavors.Users')
    }

    public Find(
      request: IUserRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions
    ): joinGRPC.IUnaryRequest<IUserProfile> {
      return this.makeUnaryRequest('Find', request, metadata, options)
    }
  }
}
