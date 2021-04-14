// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'
import * as protobufjs from 'protobufjs/light'

import { GoogleProtobuf } from '../google/protobuf/Timestamp'

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Common {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IExtraPkgMessage {
    firstName?: string
    lastName?: string
    birthDate?: Date
  }

  @protobufjs.Type.d('ExtraPkgMessage')
  export class ExtraPkgMessage
    extends protobufjs.Message<ExtraPkgMessage>
    implements ConvertibleTo<IExtraPkgMessage> {
    @protobufjs.Field.d(1, 'string')
    public firstName?: string

    @protobufjs.Field.d(2, 'string')
    public lastName?: string

    @protobufjs.Field.d(3, GoogleProtobuf.Timestamp)
    public birthDate?: GoogleProtobuf.Timestamp

    public asInterface(): IExtraPkgMessage {
      return {
        ...this,
        birthDate:
          this.birthDate != null
            ? new Date(
                (this.birthDate.seconds ?? 0) * 1000 +
                  (this.birthDate.nanos ?? 0) / 1000000
              )
            : undefined,
      }
    }

    public static fromInterface(value: IExtraPkgMessage): ExtraPkgMessage {
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

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): IExtraPkgMessage {
      return ExtraPkgMessage.decode(reader).asInterface()
    }

    public static encodePatched(
      message: IExtraPkgMessage,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      const transformedMessage = ExtraPkgMessage.fromInterface(message)
      return ExtraPkgMessage.encode(transformedMessage, writer)
    }
  }
}
