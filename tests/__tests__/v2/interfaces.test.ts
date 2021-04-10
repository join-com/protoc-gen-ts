import { Common } from './generated/common/Common'
import { Foo } from './generated/Test'

type TypesMap = {
  // Booleans
  fieldBool?: boolean
  fieldBoolRepeated?: boolean[]

  // 32-bit integers
  fieldInt32?: number
  fieldInt32Repeated?: number[]
  fieldUint32?: number
  fieldUint32Repeated?: number[]
  fieldSint32?: number
  fieldSint32Repeated?: number[]

  // 64-bit integers
  fieldInt64?: number
  fieldInt64Repeated?: number[]
  fieldUint64?: number
  fieldUint64Repeated?: number[]
  fieldSint64?: number
  fieldSint64Repeated?: number[]

  // 32-bit float point numbers
  fieldFloat?: number
  fieldFloatRepeated?: number[]

  // 64-bit float point numbers
  fieldDouble?: number
  fieldDoubleRepeated?: number[]

  // 32-bit fixed point numbers
  fieldFixed32?: number
  fieldFixed32Repeated?: number[]
  fieldSfixed32?: number
  fieldSfixed32Repeated?: number[]

  // 64-bit fixed point numbers
  fieldFixed64?: number
  fieldFixed64Repeated?: number[]
  fieldSfixed64?: number
  fieldSfixed64Repeated?: number[]

  // strings
  fieldString?: string
  fieldStringRepeated?: string[]

  // byte buffers
  fieldBytes?: Uint8Array
  fieldBytesRepeated?: Uint8Array[]

  // enums
  fieldEnum?: Foo.EnumType
  fieldEnumRepeated?: Foo.Role[]

  // messages/interfaces (same namespace)
  message?: Foo.INested
  messageRepeated?: Foo.INested[]

  // message/interfaces (different namespace)
  otherPkgMessage?: Common.IOtherPkgMessage
  otherPkgMessageRepeated?: Common.IOtherPkgMessage[]

  // timestamps
  timestamp?: Date
  timestampRepeated?: Date[]
}

describe('(v2) interfaces', () => {
  // Existence if ITest is implicitly tested by importing it

  it('generates correct types for interface fields', () => {
    // We don't directly compare Foo.ITest against TypesMap because all properties are optional

    type Match<A, B> = A extends B ? (B extends A ? true : never) : never
    type IsOk<K extends keyof TypesMap & keyof Foo.ITest> = Match<
      Foo.ITest[K],
      TypesMap[K]
    > extends true
      ? Match<NonNullable<Foo.ITest[K]>, NonNullable<TypesMap[K]>>
      : never

    // booleans
    const tBool_Ok: IsOk<'fieldBool'> = true
    expect(tBool_Ok).toBe(true)

    const tBoolR_Ok: IsOk<'fieldBoolRepeated'> = true
    expect(tBoolR_Ok).toBe(true)

    // 32-bit integers
    const tInt32_Ok: IsOk<'fieldInt32'> = true
    expect(tInt32_Ok).toBe(true)

    const tInt32R_Ok: IsOk<'fieldInt32Repeated'> = true
    expect(tInt32R_Ok).toBe(true)

    const tUint32_Ok: IsOk<'fieldUint32'> = true
    expect(tUint32_Ok).toBe(true)

    const tUint32R_Ok: IsOk<'fieldUint32Repeated'> = true
    expect(tUint32R_Ok).toBe(true)

    const tSint32_Ok: IsOk<'fieldSint32'> = true
    expect(tSint32_Ok).toBe(true)

    const tSint32R_Ok: IsOk<'fieldSint32Repeated'> = true
    expect(tSint32R_Ok).toBe(true)

    // 64-bit integers
    const tInt64_Ok: IsOk<'fieldInt64'> = true
    expect(tInt64_Ok).toBe(true)

    const tInt64R_Ok: IsOk<'fieldInt64Repeated'> = true
    expect(tInt64R_Ok).toBe(true)

    const tUint64_Ok: IsOk<'fieldUint64'> = true
    expect(tUint64_Ok).toBe(true)

    const tUint64R_Ok: IsOk<'fieldUint64Repeated'> = true
    expect(tUint64R_Ok).toBe(true)

    const tSint64_Ok: IsOk<'fieldSint64'> = true
    expect(tSint64_Ok).toBe(true)

    const tSint64R_Ok: IsOk<'fieldSint64Repeated'> = true
    expect(tSint64R_Ok).toBe(true)

    // 32-bit float point numbers
    const tFloat_Ok: IsOk<'fieldFloat'> = true
    expect(tFloat_Ok).toBe(true)

    const tFloatR_Ok: IsOk<'fieldFloatRepeated'> = true
    expect(tFloatR_Ok).toBe(true)

    // 64-bit float point numbers
    const tDouble_Ok: IsOk<'fieldDouble'> = true
    expect(tDouble_Ok).toBe(true)

    const tDoubleR_Ok: IsOk<'fieldDoubleRepeated'> = true
    expect(tDoubleR_Ok).toBe(true)

    // 32-bit fixed point numbers
    const tFixed32_Ok: IsOk<'fieldFixed32'> = true
    expect(tFixed32_Ok).toBe(true)

    const tFixed32R_Ok: IsOk<'fieldFixed32Repeated'> = true
    expect(tFixed32R_Ok).toBe(true)

    const tSfixed32_Ok: IsOk<'fieldSfixed32'> = true
    expect(tSfixed32_Ok).toBe(true)

    const tSfixed32R_Ok: IsOk<'fieldSfixed32Repeated'> = true
    expect(tSfixed32R_Ok).toBe(true)

    // 64-bit fixed point numbers
    const tFixed64_Ok: IsOk<'fieldFixed64'> = true
    expect(tFixed64_Ok).toBe(true)

    const tFixed64R_Ok: IsOk<'fieldFixed64Repeated'> = true
    expect(tFixed64R_Ok).toBe(true)

    const tSfixed64_Ok: IsOk<'fieldSfixed64'> = true
    expect(tSfixed64_Ok).toBe(true)

    const tSfixed64R_Ok: IsOk<'fieldSfixed64Repeated'> = true
    expect(tSfixed64R_Ok).toBe(true)

    // strings
    const tString_Ok: IsOk<'fieldString'> = true
    expect(tString_Ok).toBe(true)

    const tStringR_Ok: IsOk<'fieldStringRepeated'> = true
    expect(tStringR_Ok).toBe(true)

    // byte buffers
    const tBytes_Ok: IsOk<'fieldBytes'> = true
    expect(tBytes_Ok).toBe(true)

    const tBytesR_Ok: IsOk<'fieldBytesRepeated'> = true
    expect(tBytesR_Ok).toBe(true)

    // enums
    const tEnum_Ok: IsOk<'fieldEnum'> = true
    expect(tEnum_Ok).toBe(true)

    const tEnumR_Ok: IsOk<'fieldEnumRepeated'> = true
    expect(tEnumR_Ok).toBe(true)

    // messages/interfaces (same namespace)
    const tMessage_Ok: IsOk<'message'> = true
    expect(tMessage_Ok).toBe(true)

    const tMessageR_Ok: IsOk<'messageRepeated'> = true
    expect(tMessageR_Ok).toBe(true)

    // message/interfaces (different namespace)
    const tOtherPkgMessage_Ok: IsOk<'otherPkgMessage'> = true
    expect(tOtherPkgMessage_Ok).toBe(true)

    const tOtherPkgMessageR_Ok: IsOk<'otherPkgMessageRepeated'> = true
    expect(tOtherPkgMessageR_Ok).toBe(true)

    // timestamps
    // const tTimestamp_Ok: IsOk<'timestamp'> = true
    // expect(tTimestamp_Ok).toBe(true)

    // const tTimestampR_Ok: IsOk<'timestampRepeated'> = true
    // expect(tTimestampR_Ok).toBe(true)
  })
})
