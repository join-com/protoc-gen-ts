// GENERATED CODE -- DO NOT EDIT!
import * as protobufjs from 'protobufjs/minimal';
// @ts-ignore ignored as it's generated and it's difficult to predict if logger is needed
import { logger } from '@join-com/gcloud-logger-trace';
export namespace Common {
  export interface IOtherPkgMessage {
    /** @deprecated */
    firstName?: string;
    latsName?: string;
  }

  export class OtherPkgMessage implements IOtherPkgMessage {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new OtherPkgMessage();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            logger.warn('field firstName is deprecated in OtherPkgMessage');
            message.firstName = reader.string();
            break;
          case 2:
            message.latsName = reader.string();
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    /** @deprecated */
    public firstName?: string;
    public latsName?: string;
    constructor(attrs?: IOtherPkgMessage) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.firstName != null) {
        writer.uint32(10).string(this.firstName);
      }
      if (this.latsName != null) {
        writer.uint32(18).string(this.latsName);
      }
      return writer;
    }
  }
}
