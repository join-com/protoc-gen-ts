// GENERATED CODE -- DO NOT EDIT!
import * as protobufjs from 'protobufjs/minimal';
export namespace GoogleProtobuf {
  export interface Timestamp {
    seconds?: number;
    nanos?: number;
  }

  export class TimestampMsg implements Timestamp {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new TimestampMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            const seconds = reader.int64();
            message.seconds = new protobufjs.util.LongBits(
              seconds.low >>> 0,
              seconds.high >>> 0
            ).toNumber();
            break;
          case 2:
            message.nanos = reader.int32();
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    public seconds?: number;
    public nanos?: number;
    constructor(attrs?: Timestamp) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.seconds != null) {
        writer.uint32(8).int64(this.seconds);
      }
      if (this.nanos != null) {
        writer.uint32(16).int32(this.nanos);
      }
      return writer;
    }
  }
}
