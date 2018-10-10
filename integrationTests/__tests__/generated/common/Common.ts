// GENERATED CODE -- DO NOT EDIT!
import * as protobufjs from 'protobufjs/minimal';
export namespace Common {
  export interface OtherPkgMessage {
    firstName?: string;
    latsName?: string;
  }

  export class OtherPkgMessageMsg implements OtherPkgMessage {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new OtherPkgMessageMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
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
    public firstName?: string;
    public latsName?: string;
    constructor(attrs?: OtherPkgMessage) {
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
