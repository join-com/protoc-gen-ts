// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.0758bd7.1634553278
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as protobufjs from 'protobufjs/light'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Common {
  const registerGrpcClass = <T extends protobufjs.Message<T>>(
    typeName: string
  ): protobufjs.TypeDecorator<T> => {
    if (protobufjs.util.decorateRoot.get(typeName) != null) {
      // eslint-disable-next-line @typescript-eslint/ban-types
      return (
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        _: protobufjs.Constructor<T>
      ): void => {
        // Do nothing
      }
    }
    return protobufjs.Type.d(typeName)
  }
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IOtherPkgMessage {
    /**
     * @deprecated
     */
    firstName?: string
    latsName?: string
  }

  @registerGrpcClass('common_OtherPkgMessage')
  export class OtherPkgMessage
    extends protobufjs.Message<OtherPkgMessage>
    implements ConvertibleTo<IOtherPkgMessage>, IOtherPkgMessage
  {
    /**
     * @deprecated
     */
    @protobufjs.Field.d(1, 'string', 'optional')
    public firstName?: string

    @protobufjs.Field.d(2, 'string', 'optional')
    public latsName?: string

    public asInterface(): IOtherPkgMessage {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IOtherPkgMessage] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IOtherPkgMessage]
        }
      }
      return message
    }

    public static fromInterface(
      this: void,
      value: IOtherPkgMessage
    ): OtherPkgMessage {
      return OtherPkgMessage.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array
    ): IOtherPkgMessage {
      return OtherPkgMessage.decode(reader).asInterface()
    }

    public static encodePatched(
      this: void,
      message: IOtherPkgMessage,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return OtherPkgMessage.encode(message, writer)
    }
  }
}
