// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'
import * as protobufjs from 'protobufjs/light'

// eslint-disable-next-line @typescript-eslint/no-namespace
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
    implements ConvertibleTo<ITimestamp>, ITimestamp {
    @protobufjs.Field.d(1, 'int64')
    public seconds?: number

    @protobufjs.Field.d(2, 'int32')
    public nanos?: number

    public asInterface(): ITimestamp {
      return this
    }

    public static fromInterface(value: ITimestamp): Timestamp {
      return Timestamp.fromObject(value)
    }

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): ITimestamp {
      return Timestamp.decode(reader)
    }

    public static encodePatched(
      message: ITimestamp,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return Timestamp.encode(message, writer)
    }
  }
}
