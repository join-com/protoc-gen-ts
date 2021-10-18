// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.c3b10e1.1634550238
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

  @protobufjs.Type.d('google_protobuf_Timestamp')
  export class Timestamp
    extends protobufjs.Message<Timestamp>
    implements ConvertibleTo<ITimestamp>, ITimestamp
  {
    @protobufjs.Field.d(1, 'int64', 'optional')
    public seconds?: number

    @protobufjs.Field.d(2, 'int32', 'optional')
    public nanos?: number

    public asInterface(): ITimestamp {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof ITimestamp] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof ITimestamp]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: ITimestamp): Timestamp {
      return Timestamp.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): ITimestamp {
      return Timestamp.decode(reader).asInterface()
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
