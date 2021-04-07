// GENERATED CODE -- DO NOT EDIT!

// import * as joinGRPC from '@join-com/grpc'
// import * as nodeTrace from '@join-com/node-trace'
import * as protobufjs from 'protobufjs/light'

export namespace Common {
  export interface IExtraPkgMessage {
    firstName?: string
    latsName?: string
  }

  @protobufjs.Type.d('ExtraPkgMessage')
  export class ExtraPkgMessage
    extends protobufjs.Message<ExtraPkgMessage>
    implements IExtraPkgMessage {
    @protobufjs.Field.d(1, 'string')
    public firstName?: string

    @protobufjs.Field.d(2, 'string')
    public latsName?: string
  }
}
