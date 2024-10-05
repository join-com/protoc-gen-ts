// GENERATED CODE -- DO NOT EDIT!
// GENERATOR VERSION: 2.3.0.9b92372.1728131327
/* eslint-disable @typescript-eslint/no-non-null-assertion */

import * as protobufjs from 'protobufjs/light'

import { root } from './root'

protobufjs.roots['decorated'] = root

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Regressions {
  interface ConvertibleTo<T> {
    asInterface(): T
  }

  export interface IReg01Inner {
    value?: string
  }

  export interface IReg01Outer {
    inner?: IReg01Inner
  }

  export interface IMessageWithDeprecatedField {
    notDeprecated?: string
    /**
     * @deprecated
     */
    deprecated?: string
  }

  /**
   * @deprecated
   */
  export interface IDeprecatedMessageWithDeprecatedField {
    notDeprecated?: string
    /**
     * @deprecated
     */
    deprecated?: string
  }

  /**
   * @deprecated
   */
  @protobufjs.Type.d('regressions_DeprecatedMessageWithDeprecatedField')
  export class DeprecatedMessageWithDeprecatedField
    extends protobufjs.Message<DeprecatedMessageWithDeprecatedField>
    implements ConvertibleTo<IDeprecatedMessageWithDeprecatedField>, IDeprecatedMessageWithDeprecatedField
  {
    @protobufjs.Field.d(1, 'string', 'optional')
    public notDeprecated?: string

    /**
     * @deprecated
     */
    @protobufjs.Field.d(2, 'string', 'optional')
    public deprecated?: string

    public asInterface(): IDeprecatedMessageWithDeprecatedField {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IDeprecatedMessageWithDeprecatedField] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IDeprecatedMessageWithDeprecatedField]
        }
      }
      return message
    }

    public static fromInterface(
      this: void,
      value: IDeprecatedMessageWithDeprecatedField,
    ): DeprecatedMessageWithDeprecatedField {
      return DeprecatedMessageWithDeprecatedField.fromObject(value)
    }

    public static decodePatched(
      this: void,
      reader: protobufjs.Reader | Uint8Array,
    ): IDeprecatedMessageWithDeprecatedField {
      return DeprecatedMessageWithDeprecatedField.decode(reader).asInterface()
    }

    public static encodePatched(
      this: void,
      message: IDeprecatedMessageWithDeprecatedField,
      writer?: protobufjs.Writer,
    ): protobufjs.Writer {
      return DeprecatedMessageWithDeprecatedField.encode(message, writer)
    }
  }

  @protobufjs.Type.d('regressions_MessageWithDeprecatedField')
  export class MessageWithDeprecatedField
    extends protobufjs.Message<MessageWithDeprecatedField>
    implements ConvertibleTo<IMessageWithDeprecatedField>, IMessageWithDeprecatedField
  {
    @protobufjs.Field.d(1, 'string', 'optional')
    public notDeprecated?: string

    /**
     * @deprecated
     */
    @protobufjs.Field.d(2, 'string', 'optional')
    public deprecated?: string

    public asInterface(): IMessageWithDeprecatedField {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IMessageWithDeprecatedField] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IMessageWithDeprecatedField]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IMessageWithDeprecatedField): MessageWithDeprecatedField {
      return MessageWithDeprecatedField.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IMessageWithDeprecatedField {
      return MessageWithDeprecatedField.decode(reader).asInterface()
    }

    public static encodePatched(
      this: void,
      message: IMessageWithDeprecatedField,
      writer?: protobufjs.Writer,
    ): protobufjs.Writer {
      return MessageWithDeprecatedField.encode(message, writer)
    }
  }

  @protobufjs.Type.d('regressions_Reg01Inner')
  export class Reg01Inner extends protobufjs.Message<Reg01Inner> implements ConvertibleTo<IReg01Inner>, IReg01Inner {
    @protobufjs.Field.d(1, 'string', 'optional')
    public value?: string

    public asInterface(): IReg01Inner {
      const message = {
        ...this,
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IReg01Inner] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IReg01Inner]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IReg01Inner): Reg01Inner {
      return Reg01Inner.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IReg01Inner {
      return Reg01Inner.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IReg01Inner, writer?: protobufjs.Writer): protobufjs.Writer {
      return Reg01Inner.encode(message, writer)
    }
  }

  @protobufjs.Type.d('regressions_Reg01Outer')
  export class Reg01Outer extends protobufjs.Message<Reg01Outer> implements ConvertibleTo<IReg01Outer>, IReg01Outer {
    @protobufjs.Field.d(1, Reg01Inner, 'optional')
    public inner?: Reg01Inner

    public asInterface(): IReg01Outer {
      const message = {
        ...this,
        inner: this.inner?.asInterface(),
      }
      for (const fieldName of Object.keys(message)) {
        if (message[fieldName as keyof IReg01Outer] == null) {
          // We remove the key to avoid problems with code making too many assumptions
          delete message[fieldName as keyof IReg01Outer]
        }
      }
      return message
    }

    public static fromInterface(this: void, value: IReg01Outer): Reg01Outer {
      return Reg01Outer.fromObject(value)
    }

    public static decodePatched(this: void, reader: protobufjs.Reader | Uint8Array): IReg01Outer {
      return Reg01Outer.decode(reader).asInterface()
    }

    public static encodePatched(this: void, message: IReg01Outer, writer?: protobufjs.Writer): protobufjs.Writer {
      return Reg01Outer.encode(message, writer)
    }
  }
}
