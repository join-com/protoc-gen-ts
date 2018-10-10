// GENERATED CODE -- DO NOT EDIT!
import { GoogleProtobuf } from './google/protobuf/Timestamp';
import { Common } from './common/Common';
import * as protobufjs from 'protobufjs/minimal';

import * as grpc from 'grpc';
import * as grpcts from '@join-com/grpc-ts';

export namespace Foo {
  export enum Type {
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

  export interface Book {
    title?: string;
    isbn?: string;
  }

  export class BookMsg implements Book {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new BookMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            message.title = reader.string();
            break;
          case 2:
            message.isbn = reader.string();
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    public title?: string;
    public isbn?: string;
    constructor(attrs?: Book) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.title != null) {
        writer.uint32(10).string(this.title);
      }
      if (this.isbn != null) {
        writer.uint32(18).string(this.isbn);
      }
      return writer;
    }
  }

  export interface User {
    id?: number;
    name?: string;
    type?: Type;
    roles?: Role[];
    favoriteBook?: Book;
    readBooks?: Book[];
    nicks?: string[];
    types?: Type[];
    date?: Date;
    fullName?: Common.Name;
    dates?: Date[];
  }

  export class UserMsg implements User {
    public static decode(
      inReader: Uint8Array | protobufjs.Reader,
      length?: number
    ) {
      const reader = !(inReader instanceof protobufjs.Reader)
        ? protobufjs.Reader.create(inReader)
        : inReader;
      const end = length === undefined ? reader.len : reader.pos + length;
      const message = new UserMsg();
      while (reader.pos < end) {
        const tag = reader.uint32();
        switch (tag >>> 3) {
          case 1:
            message.id = reader.int32();
            break;
          case 2:
            message.name = reader.string();
            break;
          case 3:
            message.type = (val => {
              switch (val) {
                case 0:
                  return Type.UNKNOWN;
                case 1:
                  return Type.ADMIN;
                case 2:
                  return Type.USER;
                default:
                  return;
              }
            })(reader.int32());
            break;
          case 4:
            if (!(message.roles && message.roles.length)) {
              message.roles = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const roles = (val => {
                  switch (val) {
                    case 0:
                      return Role.VIEW;
                    case 1:
                      return Role.EDIT;
                    default:
                      return;
                  }
                })(reader.int32());
                if (roles) {
                  message.roles.push(roles);
                }
              }
            } else {
              const roles = (val => {
                switch (val) {
                  case 0:
                    return Role.VIEW;
                  case 1:
                    return Role.EDIT;
                  default:
                    return;
                }
              })(reader.int32());
              if (roles) {
                message.roles.push(roles);
              }
            }
            break;
          case 5:
            message.favoriteBook = BookMsg.decode(reader, reader.uint32());
            break;
          case 6:
            if (!(message.readBooks && message.readBooks.length)) {
              message.readBooks = [];
            }
            message.readBooks.push(BookMsg.decode(reader, reader.uint32()));
            break;
          case 7:
            if (!(message.nicks && message.nicks.length)) {
              message.nicks = [];
            }
            message.nicks.push(reader.string());
            break;
          case 8:
            if (!(message.types && message.types.length)) {
              message.types = [];
            }
            if ((tag & 7) === 2) {
              const end2 = reader.uint32() + reader.pos;
              while (reader.pos < end2) {
                const types = (val => {
                  switch (val) {
                    case 0:
                      return Type.UNKNOWN;
                    case 1:
                      return Type.ADMIN;
                    case 2:
                      return Type.USER;
                    default:
                      return;
                  }
                })(reader.int32());
                if (types) {
                  message.types.push(types);
                }
              }
            } else {
              const types = (val => {
                switch (val) {
                  case 0:
                    return Type.UNKNOWN;
                  case 1:
                    return Type.ADMIN;
                  case 2:
                    return Type.USER;
                  default:
                    return;
                }
              })(reader.int32());
              if (types) {
                message.types.push(types);
              }
            }
            break;
          case 9:
            const date = GoogleProtobuf.TimestampMsg.decode(
              reader,
              reader.uint32()
            );
            message.date = new Date(
              (date.seconds || 0) * 1000 + (date.nanos || 0) / 1000000
            );
            break;
          case 10:
            message.fullName = Common.NameMsg.decode(reader, reader.uint32());
            break;
          case 11:
            if (!(message.dates && message.dates.length)) {
              message.dates = [];
            }
            const dates = GoogleProtobuf.TimestampMsg.decode(
              reader,
              reader.uint32()
            );
            message.dates.push(
              new Date(
                (®.seconds || 0) * 1000 + (dates.nanos || 0) / 1000000
              )
            );
            break;
          default:
            reader.skipType(tag & 7);
            break;
        }
      }
      return message;
    }
    public id?: number;
    public name?: string;
    public type?: Type;
    public roles?: Role[];
    public favoriteBook?: Book;
    public readBooks?: Book[];
    public nicks?: string[];
    public types?: Type[];
    public date?: Date;
    public fullName?: Common.Name;
    public dates?: Date[];
    constructor(attrs?: User) {
      Object.assign(this, attrs);
    }
    public encode(writer: protobufjs.Writer = protobufjs.Writer.create()) {
      if (this.id != null) {
        writer.uint32(8).int32(this.id);
      }
      if (this.name != null) {
        writer.uint32(18).string(this.name);
      }
      if (this.type != null) {
        writer.uint32(24).int32(this.type);
      }
      if (this.roles != null) {
        for (const value of this.roles) {
          writer.uint32(32).int32(value);
        }
      }
      if (this.favoriteBook != null) {
        const msg = new BookMsg(this.favoriteBook);
        msg.encode(writer.uint32(42).fork()).ldelim();
      }
      if (this.readBooks != null) {
        for (const value of this.readBooks) {
          const msg = new BookMsg(value);
          msg.encode(writer.uint32(50).fork()).ldelim();
        }
      }
      if (this.nicks != null) {
        for (const value of this.nicks) {
          writer.uint32(58).string(value);
        }
      }
      if (this.types != null) {
        for (const value of this.types) {
          writer.uint32(64).int32(value);
        }
      }
      if (this.date != null) {
        const msg = new GoogleProtobuf.TimestampMsg(this.date);
        msg.encode(writer.uint32(74).fork()).ldelim();
      }
      if (this.fullName != null) {
        const msg = new Common.NameMsg(this.fullName);
        msg.encode(writer.uint32(82).fork()).ldelim();
      }
      if (this.dates != null) {
        for (const value of this.dates) {
          const msg = new GoogleProtobuf.TimestampMsg(value);
          msg.encode(writer.uint32(90).fork()).ldelim();
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
      responseType: UserMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: User) =>
        new UserMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) => UserMsg.decode(argBuf)
    },
    findClientStream: {
      path: '/Users/FindClientStream',
      requestStream: true,
      responseStream: false,
      requestType: RequestMsg,
      responseType: UserMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: User) =>
        new UserMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) => UserMsg.decode(argBuf)
    },
    findServerStream: {
      path: '/Users/FindServerStream',
      requestStream: false,
      responseStream: true,
      requestType: RequestMsg,
      responseType: UserMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: User) =>
        new UserMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) => UserMsg.decode(argBuf)
    },
    findBidiStream: {
      path: '/Users/FindBidiStream',
      requestStream: true,
      responseStream: true,
      requestType: RequestMsg,
      responseType: UserMsg,
      requestSerialize: (args: Request) =>
        new RequestMsg(args).encode().finish() as Buffer,
      requestDeserialize: (argBuf: Buffer) => RequestMsg.decode(argBuf),
      responseSerialize: (args: User) =>
        new UserMsg(args).encode().finish() as Buffer,
      responseDeserialize: (argBuf: Buffer) => UserMsg.decode(argBuf)
    }
  };

  export interface UsersImplementation extends grpcts.Implementations {
    find(call: grpc.ServerUnaryCall<Request>): Promise<User>;
    find(
      call: grpc.ServerUnaryCall<Request>,
      callback: grpc.sendUnaryData<User>
    ): void;
    findClientStream(call: grpc.ServerReadableStream<Request>): Promise<User>;
    findClientStream(
      call: grpc.ServerReadableStream<Request>,
      callback: grpc.sendUnaryData<User>
    ): void;
    findServerStream(call: grpc.ServerWriteableStream<Request>): void;
    findBidiStream(call: grpc.ServerDuplexStream<Request, User>): void;
  }

  export class UsersClient extends grpcts.Client {
    public find(req: Request, metadata?: grpcts.Metadata) {
      return super.makeUnaryRequest<Request, User>('find', req, metadata);
    }
    public findClientStream(metadata?: grpcts.Metadata) {
      return super.makeClientStreamRequest<Request, User>(
        'findClientStream',
        metadata
      );
    }
    public findServerStream(req: Request, metadata?: grpcts.Metadata) {
      return super.makeServerStreamRequest<Request, User>(
        'findServerStream',
        req,
        metadata
      );
    }
    public findBidiStream(metadata?: grpcts.Metadata) {
      return super.makeBidiStreamRequest<Request, User>(
        'findBidiStream',
        metadata
      );
    }
  }
}
