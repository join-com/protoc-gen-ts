// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable @typescript-eslint/no-non-null-assertion */

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

  @protobufjs.Type.d('google.protobuf.Timestamp')
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

    public static fromInterface(this: void, value: ITimestamp): Timestamp {
      return Timestamp.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): ITimestamp {
      return Timestamp.decode(reader)
    }

    public static encodePatched(
      this: void,
      message: ITimestamp,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return Timestamp.encode(message, writer)
    }
  }
}
