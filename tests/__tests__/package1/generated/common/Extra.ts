// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.1.0.d41be8f.1643383265
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from '../google/protobuf/Timestamp'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Common {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IExtraPkgMessage {
    /**
     * @deprecated
     */
    firstName?: string
    lastName?: string
    birthDate?: Date
  }

  @protobufjs.Type.d('common_ExtraPkgMessage')
  export class ExtraPkgMessage extends protobufjs.Message<ExtraPkgMessage> implements ConvertibleTo<IExtraPkgMessage> {
    /**
     * @deprecated
     */
    @protobufjs.Field.d(1, 'string', 'optional')
    public firstName?: string

    @protobufjs.Field.d(2, 'string', 'optional')
    public lastName?: string

    @protobufjs.Field.d(3, GoogleProtobuf.Timestamp, 'optional')
    public birthDate?: GoogleProtobuf.Timestamp

    public asInterface(): IExtraPkgMessage {
      const message = {
        ...this,
        birthDate:
          this.birthDate != null
            ? new Date((this.birthDate.seconds ?? 0) * 1000 + (this.birthDate.nanos ?? 0) / 1000000)
            : undefined,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IExtraPkgMessage] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IExtraPkgMessage]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IExtraPkgMessage): ExtraPkgMessage {
      const patchedValue = {
        ...value,
        birthDate:
          value.birthDate != null
            ? GoogleProtobuf.Timestamp.fromInterface({
                seconds: Math.floor(value.birthDate.getTime() / 1000),
                nanos: value.birthDate.getMilliseconds() * 1000000,
              })
            : undefined,
      }

      return ExtraPkgMessage.fromObject(patchedValue)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IExtraPkgMessage {
      return ExtraPkgMessage.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IExtraPkgMessage, writer?: protobufjs.Writer): protobufjs.Writer {
      const transformedMessage = ExtraPkgMessage.fromInterface(message)
      return ExtraPkgMessage.encode(transformedMessage, writer)
    }
  }
}
