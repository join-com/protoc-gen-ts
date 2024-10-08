// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.3.1.cbb8b0b.1728390468
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as protobufjs from 'protobufjs/light'

import { root } from '../../root'

protobufjs.roots['decorated'] = root

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GoogleProtobuf {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IEmpty {}

  @protobufjs.Type.d('google_protobuf_Empty')
  export class Empty extends protobufjs.Message<Empty> implements ConvertibleTo<IEmpty>, IEmpty {
    public asInterface(): IEmpty {
      const message = {
        ...this,
      }
      return message
    }

    public static fromInterface(this: void, value: IEmpty): Empty {
      return Empty.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IEmpty {
      return Empty.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IEmpty, writer?: protobufjs.Writer): protobufjs.Writer {
      return Empty.encode(message, writer)
    }
  }
}
