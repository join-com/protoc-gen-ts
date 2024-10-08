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

  export interface IFieldMask {
    paths?: string[]
  }

  @protobufjs.Type.d('google_protobuf_FieldMask')
  export class FieldMask extends protobufjs.Message<FieldMask> implements ConvertibleTo<IFieldMask>, IFieldMask {
    @protobufjs.Field.d(1, 'string', 'repeated')
    public paths?: string[]

    public asInterface(): IFieldMask {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        const field = message[fieldName as keyof IFieldMask]
        if (field == null || (Array.isArray(field) && (field as any[]).length === 0)) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IFieldMask]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IFieldMask): FieldMask {
      return FieldMask.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IFieldMask {
      return FieldMask.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IFieldMask, writer?: protobufjs.Writer): protobufjs.Writer {
      return FieldMask.encode(message, writer)
    }
  }
}
