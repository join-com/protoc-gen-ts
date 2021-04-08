// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'
import * as protobufjs from 'protobufjs/light'

export namespace GoogleProtobuf {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface ITimestamp {
    seconds?: number
    nanos?: number
  }

  @protobufjs.Type.d('Timestamp')
  export class Timestamp
    extends protobufjs.Message<Timestamp>
    implements ConvertibleTo<ITimestamp> {
    @protobufjs.Field.d(1, 'int64')
    public seconds?: number

    @protobufjs.Field.d(2, 'int32')
    public nanos?: number

    public asInterface(): ITimestamp {
      return this
    }
  }
}
