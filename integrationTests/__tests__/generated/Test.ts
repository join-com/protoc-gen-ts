// GENERATED CODE -- DO NOT EDIT!
import { GoogleProtobuf } from './google/protobuf/Timestamp';
import { Common } from './common/Common';
import * as protobufjs from 'protobufjs/minimal';

import * as grpc from 'grpc';
import * as grpcts from '@join-com/grpc-ts';

export namespace Foo {
  export enum EnumType {
    UNKNOWN = 'UNKNOWN',
    ADMIN = 'ADMIN',
    USER = 'USER'
  }

  export enum Role {
    VIEW = 'VIEW',
    EDIT = 'EDIT'
  }

  export interface Request {
    id?: number;
  }

  export class RequestMsg implements Request {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new RequestMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            message.id = reader.int32();
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    public id?: number;
    constructor(attrs?: Request) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.id != null) {
        writer.uint32(8).int32(this.id);
      }
      return writer;
    }
  }

  export interface Nested {
    title?: string;
  }

  export class NestedMsg implements Nested {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new NestedMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            message.title = reader.string();
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    public title?: string;
    constructor(attrs?: Nested) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.title != null) {
        writer.uint32(10).string(this.title);
      }
      return writer;
    }
  }

  export interface Test {
    fieldInt32?: number;
    fieldInt32Repeated?: number[];
    fieldDouble?: number;
    fieldDoubleRepeated?: number[];
    fieldFloat?: number;
    fieldFloatRepeated?: number[];
    fieldUint32?: number;
    fieldUint32Repeated?: number[];
    fieldUint64?: number;
    fieldUint64Repeated?: number[];
    fieldSint32?: number;
    fieldSint32Repeated?: number[];
    fieldSint64?: number;
    fieldSint64Repeated?: number[];
    fieldFixed32?: number;
    fieldFixed32Repeated?: number[];
    fieldFixed64?: number;
    fieldFixed64Repeated?: number[];
    fieldSfixed32?: number;
    fieldSfixed32Repeated?: number[];
    fieldSfixed64?: number;
    fieldSfixed64Repeated?: number[];
    fieldBool?: boolean;
    fieldBoolRepeated?: boolean[];
    fieldString?: string;
    fieldStringRepeated?: string[];
    fieldBytes?: Uint8Array;
    fieldBytesRepeated?: Uint8Array[];
    fieldEnum?: EnumType;
    fieldEnumRepeated?: EnumType[];
    message?: Nested;
    messageRepeated?: Nested[];
    timestamp?: Date;
    timestampRepeated?: Date[];
    otherPkgMessage?: Common.OtherPkgMessage;
    otherPkgMessageRepeated?: Common.OtherPkgMessage[];
    fieldInt64?: number;
    fieldInt64Repeated?: number[];
  }

  export class TestMsg implements Test {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new TestMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            message.fieldInt32 = reader.int32();
            break;
          case 2:
            if (
              !(message.fieldInt32Repeated && message.fieldInt32Repeated.length)
            ) {
              message.fieldInt32Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldInt32Repeated.push(reader.int32());
              }
            } else {
              message.fieldInt32Repeated.push(reader.int32());
            }
            break;
          case 3:
            message.fieldDouble = reader.double();
            break;
          case 4:
            if (
              !(
                message.fieldDoubleRepeated &&
                message.fieldDoubleRepeated.length
              )
            ) {
              message.fieldDoubleRepeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldDoubleRepeated.push(reader.double());
              }
            } else {
              message.fieldDoubleRepeated.push(reader.double());
            }
            break;
          case 5:
            message.fieldFloat = reader.float();
            break;
          case 6:
            if (
              !(message.fieldFloatRepeated && message.fieldFloatRepeated.length)
            ) {
              message.fieldFloatRepeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldFloatRepeated.push(reader.float());
              }
            } else {
              message.fieldFloatRepeated.push(reader.float());
            }
            break;
          case 7:
            message.fieldUint32 = reader.uint32();
            break;
          case 8:
            if (
              !(
                message.fieldUint32Repeated &&
                message.fieldUint32Repeated.length
              )
            ) {
              message.fieldUint32Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldUint32Repeated.push(reader.uint32());
              }
            } else {
              message.fieldUint32Repeated.push(reader.uint32());
            }
            break;
          case 9:
            const fieldUint64 = reader.uint64();
            message.fieldUint64 = new protobufjs.util.LongBits(
              fieldUint64.low >>> 0,
              fieldUint64.high >>> 0
            ).toNumber();
            break;
          case 10:
            if (
              !(
                message.fieldUint64Repeated &&
                message.fieldUint64Repeated.length
              )
            ) {
              message.fieldUint64Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const fieldUint64Repeated = reader.uint64();
                message.fieldUint64Repeated.push(
                  new protobufjs.util.LongBits(
                    fieldUint64Repeated.low >>> 0,
                    fieldUint64Repeated.high >>> 0
                  ).toNumber()
                );
              }
            } else {
              const fieldUint64Repeated = reader.uint64();
              message.fieldUint64Repeated.push(
                new protobufjs.util.LongBits(
                  fieldUint64Repeated.low >>> 0,
                  fieldUint64Repeated.high >>> 0
                ).toNumber()
              );
            }
            break;
          case 11:
            message.fieldSint32 = reader.sint32();
            break;
          case 12:
            if (
              !(
                message.fieldSint32Repeated &&
                message.fieldSint32Repeated.length
              )
            ) {
              message.fieldSint32Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldSint32Repeated.push(reader.sint32());
              }
            } else {
              message.fieldSint32Repeated.push(reader.sint32());
            }
            break;
          case 13:
            const fieldSint64 = reader.sint64();
            message.fieldSint64 = new protobufjs.util.LongBits(
              fieldSint64.low >>> 0,
              fieldSint64.high >>> 0
            ).toNumber();
            break;
          case 14:
            if (
              !(
                message.fieldSint64Repeated &&
                message.fieldSint64Repeated.length
              )
            ) {
              message.fieldSint64Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const fieldSint64Repeated = reader.sint64();
                message.fieldSint64Repeated.push(
                  new protobufjs.util.LongBits(
                    fieldSint64Repeated.low >>> 0,
                    fieldSint64Repeated.high >>> 0
                  ).toNumber()
                );
              }
            } else {
              const fieldSint64Repeated = reader.sint64();
              message.fieldSint64Repeated.push(
                new protobufjs.util.LongBits(
                  fieldSint64Repeated.low >>> 0,
                  fieldSint64Repeated.high >>> 0
                ).toNumber()
              );
            }
            break;
          case 15:
            message.fieldFixed32 = reader.fixed32();
            break;
          case 16:
            if (
              !(
                message.fieldFixed32Repeated &&
                message.fieldFixed32Repeated.length
              )
            ) {
              message.fieldFixed32Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldFixed32Repeated.push(reader.fixed32());
              }
            } else {
              message.fieldFixed32Repeated.push(reader.fixed32());
            }
            break;
          case 17:
            const fieldFixed64 = reader.fixed64();
            message.fieldFixed64 = new protobufjs.util.LongBits(
              fieldFixed64.low >>> 0,
              fieldFixed64.high >>> 0
            ).toNumber();
            break;
          case 18:
            if (
              !(
                message.fieldFixed64Repeated &&
                message.fieldFixed64Repeated.length
              )
            ) {
              message.fieldFixed64Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const fieldFixed64Repeated = reader.fixed64();
                message.fieldFixed64Repeated.push(
                  new protobufjs.util.LongBits(
                    fieldFixed64Repeated.low >>> 0,
                    fieldFixed64Repeated.high >>> 0
                  ).toNumber()
                );
              }
            } else {
              const fieldFixed64Repeated = reader.fixed64();
              message.fieldFixed64Repeated.push(
                new protobufjs.util.LongBits(
                  fieldFixed64Repeated.low >>> 0,
                  fieldFixed64Repeated.high >>> 0
                ).toNumber()
              );
            }
            break;
          case 19:
            message.fieldSfixed32 = reader.sfixed32();
            break;
          case 20:
            if (
              !(
                message.fieldSfixed32Repeated &&
                message.fieldSfixed32Repeated.length
              )
            ) {
              message.fieldSfixed32Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldSfixed32Repeated.push(reader.sfixed32());
              }
            } else {
              message.fieldSfixed32Repeated.push(reader.sfixed32());
            }
            break;
          case 21:
            const fieldSfixed64 = reader.sfixed64();
            message.fieldSfixed64 = new protobufjs.util.LongBits(
              fieldSfixed64.low >>> 0,
              fieldSfixed64.high >>> 0
            ).toNumber();
            break;
          case 22:
            if (
              !(
                message.fieldSfixed64Repeated &&
                message.fieldSfixed64Repeated.length
              )
            ) {
              message.fieldSfixed64Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const fieldSfixed64Repeated = reader.sfixed64();
                message.fieldSfixed64Repeated.push(
                  new protobufjs.util.LongBits(
                    fieldSfixed64Repeated.low >>> 0,
                    fieldSfixed64Repeated.high >>> 0
                  ).toNumber()
                );
              }
            } else {
              const fieldSfixed64Repeated = reader.sfixed64();
              message.fieldSfixed64Repeated.push(
                new protobufjs.util.LongBits(
                  fieldSfixed64Repeated.low >>> 0,
                  fieldSfixed64Repeated.high >>> 0
                ).toNumber()
              );
            }
            break;
          case 23:
            message.fieldBool = reader.bool();
            break;
          case 24:
            if (
              !(message.fieldBoolRepeated && message.fieldBoolRepeated.length)
            ) {
              message.fieldBoolRepeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                message.fieldBoolRepeated.push(reader.bool());
              }
            } else {
              message.fieldBoolRepeated.push(reader.bool());
            }
            break;
          case 25:
            message.fieldString = reader.string();
            break;
          case 26:
            if (
              !(
                message.fieldStringRepeated &&
                message.fieldStringRepeated.length
              )
            ) {
              message.fieldStringRepeated = [];
            }
            message.fieldStringRepeated.push(reader.string());
            break;
          case 27:
            message.fieldBytes = new Uint8Array(reader.bytes());
            break;
          case 28:
            if (
              !(message.fieldBytesRepeated && message.fieldBytesRepeated.length)
            ) {
              message.fieldBytesRepeated = [];
            }
            message.fieldBytesRepeated.push(new Uint8Array(reader.bytes()));
            break;
          case 29:
            message.fieldEnum = (val => {
              switch (val) {
                case 0:
                  return EnumType.UNKNOWN;
                case 1:
                  return EnumType.ADMIN;
                case 2:
                  return EnumType.USER;
                default:
                  return;
              }
            })(reader.int32());
            break;
          case 30:
            if (
              !(message.fieldEnumRepeated && message.fieldEnumRepeated.length)
            ) {
              message.fieldEnumRepeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const fieldEnumRepeated = (val => {
                  switch (val) {
                    case 0:
                      return EnumType.UNKNOWN;
                    case 1:
                      return EnumType.ADMIN;
                    case 2:
                      return EnumType.USER;
                    default:
                      return;
                  }
                })(reader.int32());
                if (fieldEnumRepeated) {
                  message.fieldEnumRepeated.push(fieldEnumRepeated);
                }
              }
            } else {
              const fieldEnumRepeated = (val => {
                switch (val) {
                  case 0:
                    return EnumType.UNKNOWN;
                  case 1:
                    return EnumType.ADMIN;
                  case 2:
                    return EnumType.USER;
                  default:
                    return;
                }
              })(reader.int32());
              if (fieldEnumRepeated) {
                message.fieldEnumRepeated.push(fieldEnumRepeated);
              }
            }
            break;
          case 31:
            message.message = NestedMsg.decode(reader, reader.uint32());
            break;
          case 32:
            if (!(message.messageRepeated && message.messageRepeated.length)) {
              message.messageRepeated = [];
            }
            message.messageRepeated.push(
              NestedMsg.decode(reader, reader.uint32())
            );
            break;
          case 33:
            const timestamp = GoogleProtobuf.TimestampMsg.decode(
              reader,
              reader.uint32()
            );
            message.timestamp = new Date(
              (timestamp.seconds || 0) * 1000 + (timestamp.nanos || 0) / 1000000
            );
            break;
          case 34:
            if (
              !(message.timestampRepeated && message.timestampRepeated.length)
            ) {
              message.timestampRepeated = [];
            }
            const timestampRepeated = GoogleProtobuf.TimestampMsg.decode(
              reader,
              reader.uint32()
            );
            message.timestampRepeated.push(
              new Date(
                (timestampRepeated.seconds || 0) * 1000 +
                  (timestampRepeated.nanos || 0) / 1000000
              )
            );
            break;
          case 35:
            message.otherPkgMessage = Common.OtherPkgMessageMsg.decode(
              reader,
              reader.uint32()
            );
            break;
          case 36:
            if (
              !(
                message.otherPkgMessageRepeated &&
                message.otherPkgMessageRepeated.length
              )
            ) {
              message.otherPkgMessageRepeated = [];
            }
            message.otherPkgMessageRepeated.push(
              Common.OtherPkgMessageMsg.decode(reader, reader.uint32())
            );
            break;
          case 37:
            const fieldInt64 = reader.int64();
            message.fieldInt64 = new protobufjs.util.LongBits(
              fieldInt64.low >>> 0,
              fieldInt64.high >>> 0
            ).toNumber();
            break;
          case 38:
            if (
              !(message.fieldInt64Repeated && message.fieldInt64Repeated.length)
            ) {
              message.fieldInt64Repeated = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const fieldInt64Repeated = reader.int64();
                message.fieldInt64Repeated.push(
                  new protobufjs.util.LongBits(
                    fieldInt64Repeated.low >>> 0,
                    fieldInt64Repeated.high >>> 0
                  ).toNumber()
                );
              }
            } else {
              const fieldInt64Repeated = reader.int64();
              message.fieldInt64Repeated.push(
                new protobufjs.util.LongBits(
                  fieldInt64Repeated.low >>> 0,
                  fieldInt64Repeated.high >>> 0
                ).toNumber()
              );
            }
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    public fieldInt32?: number;
    public fieldInt32Repeated?: number[];
    public fieldDouble?: number;
    public fieldDoubleRepeated?: number[];
    public fieldFloat?: number;
    public fieldFloatRepeated?: number[];
    public fieldUint32?: number;
    public fieldUint32Repeated?: number[];
    public fieldUint64?: number;
    public fieldUint64Repeated?: number[];
    public fieldSint32?: number;
    public fieldSint32Repeated?: number[];
    public fieldSint64?: number;
    public fieldSint64Repeated?: number[];
    public fieldFixed32?: number;
    public fieldFixed32Repeated?: number[];
    public fieldFixed64?: number;
    public fieldFixed64Repeated?: number[];
    public fieldSfixed32?: number;
    public fieldSfixed32Repeated?: number[];
    public fieldSfixed64?: number;
    public fieldSfixed64Repeated?: number[];
    public fieldBool?: boolean;
    public fieldBoolRepeated?: boolean[];
    public fieldString?: string;
    public fieldStringRepeated?: string[];
    public fieldBytes?: Uint8Array;
    public fieldBytesRepeated?: Uint8Array[];
    public fieldEnum?: EnumType;
    public fieldEnumRepeated?: EnumType[];
    public message?: Nested;
    public messageRepeated?: Nested[];
    public timestamp?: Date;
    public timestampRepeated?: Date[];
    public otherPkgMessage?: Common.OtherPkgMessage;
    public otherPkgMessageRepeated?: Common.OtherPkgMessage[];
    public fieldInt64?: number;
    public fieldInt64Repeated?: number[];
    constructor(attrs?: Test) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.fieldInt32 != null) {
        writer.uint32(8).int32(this.fieldInt32);
      }
      if (this.fieldInt32Repeated != null) {
        for (const value of this.fieldInt32Repeated) {
          writer.uint32(16).int32(value);
        }
      }
      if (this.fieldDouble != null) {
        writer.uint32(25).double(this.fieldDouble);
      }
      if (this.fieldDoubleRepeated != null) {
        for (const value of this.fieldDoubleRepeated) {
          writer.uint32(33).double(value);
        }
      }
      if (this.fieldFloat != null) {
        writer.uint32(45).float(this.fieldFloat);
      }
      if (this.fieldFloatRepeated != null) {
        for (const value of this.fieldFloatRepeated) {
          writer.uint32(53).float(value);
        }
      }
      if (this.fieldUint32 != null) {
        writer.uint32(56).uint32(this.fieldUint32);
      }
      if (this.fieldUint32Repeated != null) {
        for (const value of this.fieldUint32Repeated) {
          writer.uint32(64).uint32(value);
        }
      }
      if (this.fieldUint64 != null) {
        writer.uint32(72).uint64(this.fieldUint64);
      }
      if (this.fieldUint64Repeated != null) {
        for (const value of this.fieldUint64Repeated) {
          writer.uint32(80).uint64(value);
        }
      }
      if (this.fieldSint32 != null) {
        writer.uint32(88).sint32(this.fieldSint32);
      }
      if (this.fieldSint32Repeated != null) {
        for (const value of this.fieldSint32Repeated) {
          writer.uint32(96).sint32(value);
        }
      }
      if (this.fieldSint64 != null) {
        writer.uint32(104).sint64(this.fieldSint64);
      }
      if (this.fieldSint64Repeated != null) {
        for (const value of this.fieldSint64Repeated) {
          writer.uint32(112).sint64(value);
        }
      }
      if (this.fieldFixed32 != null) {
        writer.uint32(125).fixed32(this.fieldFixed32);
      }
      if (this.fieldFixed32Repeated != null) {
        for (const value of this.fieldFixed32Repeated) {
          writer.uint32(133).fixed32(value);
        }
      }
      if (this.fieldFixed64 != null) {
        writer.uint32(137).fixed64(this.fieldFixed64);
      }
      if (this.fieldFixed64Repeated != null) {
        for (const value of this.fieldFixed64Repeated) {
          writer.uint32(145).fixed64(value);
        }
      }
      if (this.fieldSfixed32 != null) {
        writer.uint32(157).sfixed32(this.fieldSfixed32);
      }
      if (this.fieldSfixed32Repeated != null) {
        for (const value of this.fieldSfixed32Repeated) {
          writer.uint32(165).sfixed32(value);
        }
      }
      if (this.fieldSfixed64 != null) {
        writer.uint32(169).sfixed64(this.fieldSfixed64);
      }
      if (this.fieldSfixed64Repeated != null) {
        for (const value of this.fieldSfixed64Repeated) {
          writer.uint32(177).sfixed64(value);
        }
      }
      if (this.fieldBool != null) {
        writer.uint32(184).bool(this.fieldBool);
      }
      if (this.fieldBoolRepeated != null) {
        for (const value of this.fieldBoolRepeated) {
          writer.uint32(192).bool(value);
        }
      }
      if (this.fieldString != null) {
        writer.uint32(202).string(this.fieldString);
      }
      if (this.fieldStringRepeated != null) {
        for (const value of this.fieldStringRepeated) {
          writer.uint32(210).string(value);
        }
      }
      if (this.fieldBytes != null) {
        writer.uint32(218).bytes(this.fieldBytes);
      }
      if (this.fieldBytesRepeated != null) {
        for (const value of this.fieldBytesRepeated) {
          writer.uint32(226).bytes(value);
        }
      }
      if (this.fieldEnum != null) {
        const fieldEnum = (val => {
          switch (val) {
            case EnumType.UNKNOWN:
              return 0;
            case EnumType.ADMIN:
              return 1;
            case EnumType.USER:
              return 2;
            default:
              return;
          }
        })(this.fieldEnum);
        if (fieldEnum != null) {
          writer.uint32(232).int32(fieldEnum);
        }
      }
      if (this.fieldEnumRepeated != null) {
        for (const value of this.fieldEnumRepeated) {
          const fieldEnumRepeated = (val => {
            switch (val) {
              case EnumType.UNKNOWN:
                return 0;
              case EnumType.ADMIN:
                return 1;
              case EnumType.USER:
                return 2;
              default:
                return;
            }
          })(value);
          if (fieldEnumRepeated != null) {
            writer.uint32(240).int32(fieldEnumRepeated);
          }
        }
      }
      if (this.message != null) {
        const msg = new NestedMsg(this.message);
        msg.encode(writer.uint32(250).fork()).ldelim();
      }
      if (this.messageRepeated != null) {
        for (const value of this.messageRepeated) {
          if (!value) {
            continue;
          }
          const msg = new NestedMsg(value);
          msg.encode(writer.uint32(258).fork()).ldelim();
        }
      }
      if (this.timestamp != null) {
        const msg = new GoogleProtobuf.TimestampMsg({
          seconds: Math.floor(this.timestamp.getTime() / 1000),
          nanos: this.timestamp.getMilliseconds() * 1000000
        });
        msg.encode(writer.uint32(266).fork()).ldelim();
      }
      if (this.timestampRepeated != null) {
        for (const value of this.timestampRepeated) {
          if (!value) {
            continue;
          }
          const msg = new GoogleProtobuf.TimestampMsg({
            seconds: Math.floor(value.getTime() / 1000),
            nanos: value.getMilliseconds() * 1000000
          });
          msg.encode(writer.uint32(274).fork()).ldelim();
        }
      }
      if (this.otherPkgMessage != null) {
        const msg = new Common.OtherPkgMessageMsg(this.otherPkgMessage);
        msg.encode(writer.uint32(282).fork()).ldelim();
      }
      if (this.otherPkgMessageRepeated != null) {
        for (const value of this.otherPkgMessageRepeated) {
          if (!value) {
            continue;
          }
          const msg = new Common.OtherPkgMessageMsg(value);
          msg.encode(writer.uint32(290).fork()).ldelim();
        }
      }
      if (this.fieldInt64 != null) {
        writer.uint32(296).int64(this.fieldInt64);
      }
      if (this.fieldInt64Repeated != null) {
        for (const value of this.fieldInt64Repeated) {
          writer.uint32(304).int64(value);
        }
      }
      return writer;
    }
  }

  export const usersServiceDefinition = {
    find: {
      path: '/Users/Find',
      requestStream: false,
      responseStream: false,
      requestType: RequestMsg,
      responseType: Common.OtherPkgMessageMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: Common.OtherPkgMessage) =>
        new Common.OtherPkgMessageMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) =>
        Common.OtherPkgMessageMsg.decode(argBuf)
    },
    findClientStream: {
      path: '/Users/FindClientStream',
      requestStream: true,
      responseStream: false,
      requestType: RequestMsg,
      responseType: Common.OtherPkgMessageMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: Common.OtherPkgMessage) =>
        new Common.OtherPkgMessageMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) =>
        Common.OtherPkgMessageMsg.decode(argBuf)
    },
    findServerStream: {
      path: '/Users/FindServerStream',
      requestStream: false,
      responseStream: true,
      requestType: RequestMsg,
      responseType: Common.OtherPkgMessageMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: Common.OtherPkgMessage) =>
        new Common.OtherPkgMessageMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) =>
        Common.OtherPkgMessageMsg.decode(argBuf)
    },
    findBidiStream: {
      path: '/Users/FindBidiStream',
      requestStream: true,
      responseStream: true,
      requestType: RequestMsg,
      responseType: Common.OtherPkgMessageMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: Common.OtherPkgMessage) =>
        new Common.OtherPkgMessageMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) =>
        Common.OtherPkgMessageMsg.decode(argBuf)
    }
  };

  export interface UsersImplementation extends grpcts.Implementations {
    find(call: grpc.ServerUnaryCall<Request>): Promise<Common.OtherPkgMessage>;
    find(
      call: grpc.ServerUnaryCall<Request>,
      callback: grpc.sendUnaryData<Common.OtherPkgMessage>
    ): void;
    findClientStream(
      call: grpc.ServerReadableStream<Request>
    ): Promise<Common.OtherPkgMessage>;
    findClientStream(
      call: grpc.ServerReadableStream<Request>,
      callback: grpc.sendUnaryData<Common.OtherPkgMessage>
    ): void;
    findServerStream(call: grpc.ServerWriteableStream<Request>): void;
    findBidiStream(
      call: grpc.ServerDuplexStream<Request, Common.OtherPkgMessage>
    ): void;
  }

  export class UsersClient extends grpcts.Client {
    public find(req: Request, metadata?: grpcts.Metadata) {
      return super.makeUnaryRequest<Request, Common.OtherPkgMessage>(
        'find',
        req,
        metadata
      );
    }
    public findClientStream(metadata?: grpcts.Metadata) {
      return super.makeClientStreamRequest<Request, Common.OtherPkgMessage>(
        'findClientStream',
        metadata
      );
    }
    public findServerStream(req: Request, metadata?: grpcts.Metadata) {
      return super.makeServerStreamRequest<Request, Common.OtherPkgMessage>(
        'findServerStream',
        req,
        metadata
      );
    }
    public findBidiStream(metadata?: grpcts.Metadata) {
      return super.makeBidiStreamRequest<Request, Common.OtherPkgMessage>(
        'findBidiStream',
        metadata
      );
    }
  }
}
