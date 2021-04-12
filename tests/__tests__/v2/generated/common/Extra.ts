// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'
import * as protobufjs from 'protobufjs/light'

export namespace Common {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IExtraPkgMessage {
    firstName?: string
    latsName?: string
  }

  @protobufjs.Type.d('ExtraPkgMessage')
  export class ExtraPkgMessage
    extends protobufjs.Message<ExtraPkgMessage>
    implements ConvertibleTo<IExtraPkgMessage>, IExtraPkgMessage {
    @protobufjs.Field.d(1, 'string')
    public firstName?: string

    @protobufjs.Field.d(2, 'string')
    public latsName?: string

    public asInterface(): IExtraPkgMessage {
      return this
    }

    public static fromInterface(value: IExtraPkgMessage): ExtraPkgMessage {
      return ExtraPkgMessage.fromObject(value)
    }

    public static decodePatched(
      reader: protobufjs.Reader | Uint8Array
    ): IExtraPkgMessage {
      return ExtraPkgMessage.decode(reader)
    }

    public static encodePatched(
      message: IExtraPkgMessage,
      writer?: protobufjs.Writer
    ): protobufjs.Writer {
      return ExtraPkgMessage.encode(message, writer)
    }
  }
}
