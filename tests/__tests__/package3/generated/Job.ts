// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.7f9219a.1644779393
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as joinGRPC from '@join-com/grpc'
import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from './google/protobuf/Timestamp'

import { grpc } from '@join-com/grpc'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Jobs {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IJob {
    id?: number
    title?: string
    createdAt?: Date
  }

  export interface IFindByIdRequest {
    id?: number
  }

  export interface IFindByIdResponse {
    job?: IJob
  }

  @protobufjs.Type.d('jobs_FindByIdRequest')
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

  @protobufjs.Type.d('jobs_Job')
  export class Job extends protobufjs.Message<Job> implements ConvertibleTo<IJob> {
    @protobufjs.Field.d(1, 'int32', 'optional')
    public id?: number

    @protobufjs.Field.d(2, 'string', 'optional')
    public title?: string

    @protobufjs.Field.d(3, GoogleProtobuf.Timestamp, 'optional')
    public createdAt?: GoogleProtobuf.Timestamp

    public asInterface(): IJob {
      const message = {
        ...this,
        createdAt:
          this.createdAt != null
            ? new Date((this.createdAt.seconds ?? 0) * 1000 + (this.createdAt.nanos ?? 0) / 1000000)
            : undefined,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IJob] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IJob]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IJob): Job {
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

      return Job.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IJob {
      return Job.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IJob, writer?: protobufjs.Writer): protobufjs.Writer {
      const transformedMessage = Job.fromInterface(message)
      return Job.encode(transformedMessage, writer)
    }
  }

  @protobufjs.Type.d('jobs_FindByIdResponse')
  export class FindByIdResponse
    extends protobufjs.Message<FindByIdResponse>
    implements ConvertibleTo<IFindByIdResponse>
  {
    @protobufjs.Field.d(1, Job, 'optional')
    public job?: Job

    public asInterface(): IFindByIdResponse {
      const message = {
        ...this,
        job: this.job?.asInterface(),
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
        job: value.job != null ? Job.fromInterface(value.job) : undefined,
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

  export interface IJobsServiceImplementation {
    FindById: grpc.handleUnaryCall<IFindByIdRequest, IFindByIdResponse>
  }

  export const jobsServiceDefinition: grpc.ServiceDefinition<IJobsServiceImplementation> = {
    FindById: {
      path: '/jobs.Jobs/FindById',
      requestStream: false,
      responseStream: false,
      requestSerialize: (request: IFindByIdRequest) => FindByIdRequest.encodePatched(request).finish() as Buffer,
      requestDeserialize: FindByIdRequest.decodePatched,
      responseSerialize: (response: IFindByIdResponse) => FindByIdResponse.encodePatched(response).finish() as Buffer,
      responseDeserialize: FindByIdResponse.decodePatched,
    },
  }

  export abstract class AbstractJobsService extends joinGRPC.Service<IJobsServiceImplementation> {
    constructor(protected readonly logger?: joinGRPC.INoDebugLogger, trace?: joinGRPC.IServiceTrace) {
      super(
        jobsServiceDefinition,
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

  export interface IJobsClient extends joinGRPC.IExtendedClient<IJobsServiceImplementation, 'jobs.Jobs'> {
    findById(
      request: IFindByIdRequest,
      metadata?: Record<string, string>,
      options?: grpc.CallOptions,
    ): joinGRPC.IUnaryRequest<IFindByIdResponse>
  }

  export class JobsClient extends joinGRPC.Client<IJobsServiceImplementation, 'jobs.Jobs'> implements IJobsClient {
    constructor(config: joinGRPC.ISimplifiedClientConfig<IJobsServiceImplementation>) {
      super(
        {
          ...config,
          serviceDefinition: jobsServiceDefinition,
          credentials: config?.credentials ?? grpc.credentials.createInsecure(),
        },
        'jobs.Jobs',
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
