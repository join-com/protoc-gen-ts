// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.d41be8f.1643383265
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as protobufjs from 'protobufjs/light'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GoogleProtobuf {
  const registerCommonClass = <T extends protobufjs.Message<T>>(typeName: string): protobufjs.TypeDecorator<T> => {
    const registeredType = protobufjs.util.decorateRoot.get(typeName)
    if (registeredType == null) {
      return protobufjs.Type.d(typeName)
    }
    return (ctor: protobufjs.Constructor<T>): void => {
      Object.defineProperty(ctor, '$type', { value: registeredType, enumerable: false })
    }
  }
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IEmpty {}

  @registerCommonClass('google_protobuf_Empty')
  export class Empty extends protobufjs.Message<Empty> implements ConvertibleTo<IEmpty>, IEmpty {
    public asInterface(): IEmpty {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IEmpty] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IEmpty]
        }
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
