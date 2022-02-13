// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.d41be8f.1643383265
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as joinGRPC from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from './google/protobuf/Timestamp'

import { grpc } from '@join-com/grpc'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Companies {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface ICompany {
    id?: number
    name?: string
    createdAt?: Date
  }

  export interface IFindByIdRequest {
    id?: number
  }

  export interface IFindByIdResponse {
    company?: ICompany
  }

  @protobufjs.Type.d('companies_Company')
  export class Company extends protobufjs.Message<Company> implements ConvertibleTo<ICompany> {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    @protobufjs.Field.d(2, 'string', 'optional')
    public name?: string

    @protobufjs.Field.d(3, GoogleProtobuf.Timestamp, 'optional')
    public createdAt?: GoogleProtobuf.Timestamp

    public asInterface(): ICompany {
      const message = {
        ...this,
        createdAt:
          this.createdAt != null
            ? new Date((this.createdAt.seconds ?? 0) * 1000 + (this.createdAt.nanos ?? 0) / 1000000)
            : undefined,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof ICompany] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof ICompany]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: ICompany): Company {
      const patchedValue = {
        ...value,
        createdAt:
          value.createdAt != null
            ? GoogleProtobuf.Timestamp.fromInterface({
                seconds: Math.floor(value.createdAt.getTime() / 1000),
                nanos: value.createdAt.getMilliseconds() * 1000000,
              })
            : undefined,
      }

      return Company.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): ICompany {
      return Company.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: ICompany, writer?: protobufjs.Writer): protobufjs.Writer {
      const transformedMessage = Company.fromInterface(message)
      return Company.encode(transformedMessage, writer)
    }
  }

  @protobufjs.Type.d('companies_FindByIdRequest')
  export class FindByIdRequest
    extends protobufjs.Message<FindByIdRequest>
    implements ConvertibleTo<IFindByIdRequest>, IFindByIdRequest
  {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    public asInterface(): IFindByIdRequest {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IFindByIdRequest] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IFindByIdRequest]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IFindByIdRequest): FindByIdRequest {
      return FindByIdRequest.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IFindByIdRequest {
      return FindByIdRequest.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IFindByIdRequest, writer?: protobufjs.Writer): protobufjs.Writer {
      return FindByIdRequest.encode(message, writer)
    }
  }

  @protobufjs.Type.d('companies_FindByIdResponse')
  export class FindByIdResponse
    extends protobufjs.Message<FindByIdResponse>
    implements ConvertibleTo<IFindByIdResponse>
  {
    @protobufjs.Field.d(1, Company, 'optional')
    public company?: Company

    public asInterface(): IFindByIdResponse {
      const message = {
        ...this,
        company: this.company?.asInterface(),
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IFindByIdResponse] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IFindByIdResponse]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IFindByIdResponse): FindByIdResponse {
      const patchedValue = {
        ...value,
        company: value.company != null ? Company.fromInterface(value.company) : undefined,
      }

      return FindByIdResponse.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IFindByIdResponse {
      return FindByIdResponse.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IFindByIdResponse, writer?: protobufjs.Writer): protobufjs.Writer {
      const transformedMessage = FindByIdResponse.fromInterface(message)
      return FindByIdResponse.encode(transformedMessage, writer)
    }
  }

  export interface ICompaniesServiceImplementation {
    FindById: grpc.handleUnaryCall<IFindByIdRequest, IFindByIdResponse>
  }

  export const companiesServiceDefinition: grpc.ServiceDefinition<ICompaniesServiceImplementation> = {
    FindById: {
      path: '/companies.Companies/FindById',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IFindByIdRequest) => FindByIdRequest.encodePatched(request).finish() as Buffer,
      requestDeserialize: FindByIdRequest.decodePatched,
      responseSerialize: (response: IFindByIdResponse) => FindByIdResponse.encodePatched(response).finish() as Buffer,
      responseDeserialize: FindByIdResponse.decodePatched,
    },
  }

  export abstract class AbstractCompaniesService extends joinGRPC.Service<ICompaniesServiceImplementation> {
    constructor(protected readonly logger?: joinGRPC.INoDebugLogger, trace?: joinGRPC.IServiceTrace) {
      super(
        companiesServiceDefinition,
        {
          findById: call => this.FindById(call),
        },
        logger,
        trace,
      )
    }

    public abstract FindById(
      call: grpc.ServerUnaryCall<IFindByIdRequest, IFindByIdResponse>,
    ): Promise<IFindByIdResponse>
  }

  export interface ICompaniesClient
    extends joinGRPC.IExtendedClient<ICompaniesServiceImplementation, 'companies.Companies'> {
    findById(
      request: IFindByIdRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IFindByIdResponse>
  }

  export class CompaniesClient
    extends joinGRPC.Client<ICompaniesServiceImplementation, 'companies.Companies'>
    implements ICompaniesClient
  {
    constructor(config: joinGRPC.ISimplifiedClientConfig<ICompaniesServiceImplementation>) {
      super(
        {
          ...config,
          serviceDefinition: companiesServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'companies.Companies',
      )
    }

    public findById(
      request: IFindByIdRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IFindByIdResponse> {
      return this.makeUnaryRequest('FindById', request, metadata, options)
    }
  }
}
