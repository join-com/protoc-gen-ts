// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as protobufjs from 'protobufjs/light'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Common {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IOtherPkgMessage {
    firstName?: string
    latsName?: string
  }

  @protobufjs.Type.d('common_OtherPkgMessage')
  export class OtherPkgMessage
    extends protobufjs.Message<OtherPkgMessage>
    implements ConvertibleTo<IOtherPkgMessage>, IOtherPkgMessage {
    @protobufjs.Field.d(1, 'string')
    public firstName?: string

    @protobufjs.Field.d(2, 'string')
    public latsName?: string

    public asInterface(): IOtherPkgMessage {
      return this
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
      return OtherPkgMessage.decode(reader)
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
