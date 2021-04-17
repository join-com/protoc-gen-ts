import { Foo } from './generated/Test'
import { Foo as LegacyFoo } from '../legacy/generated/Test'

describe('(v2) legacy compatibility', () => {
  const request: Foo.IRequest = { id: 42 }
  const nested: Foo.INested = { title: 'The Test' }
  const complex: Foo.ITest = {
    fieldEnum: 'ADMIN',
    fieldEnumRepeated: ['VIEW', 'EDIT'],
    message: nested,
    messageRepeated: [nested, nested],
    fieldBytes: Buffer.from('Ubik'),
  }

  it('encode generates same buffers as legacy generator for simple types', () => {
    const requestBuffer = Foo.Request.encodePatched(request).finish()
    const legacyRequestBuffer = new LegacyFoo.Request(request).encode().finish()
    expect(requestBuffer).toStrictEqual(legacyRequestBuffer)

    const nestedBuffer = Foo.Nested.encodePatched(nested).finish()
    const legacyNestedBuffer = new LegacyFoo.Nested(nested).encode().finish()
    expect(nestedBuffer).toStrictEqual(legacyNestedBuffer)
  })

  it('can decode legacy buffers', () => {
    const legacyComplexBuffer = new LegacyFoo.Test(complex).encode().finish()
    const reconstructedComplex = Foo.Test.decodePatched(legacyComplexBuffer)

    expect(reconstructedComplex).toMatchObject(complex)
  })

  it('can encode valid buffers for legacy decoders', () => {
    const complexBuffer = Foo.Test.encodePatched(complex).finish()
    const reconstructedComplex = LegacyFoo.Test.decode(complexBuffer)

    const { fieldBytes, ...reconstructedMinusBytes } = reconstructedComplex
    const { fieldBytes: originalFieldBytes, ...complexMinusBytes } = complex

    // We compare the byte fields separately, we'll see why in a few lines
    expect(reconstructedMinusBytes).toMatchObject(complexMinusBytes)

    // Our new class uses Buffer (that subclasses Uint8Array),
    // while the old one uses "plain" Uint8Array.
    // In any case, they are both compatible, and represent the same values.
    expect(Buffer.from(fieldBytes ?? '')).toEqual(originalFieldBytes)
    expect(new Uint8Array(originalFieldBytes ?? [])).toEqual(fieldBytes)
  })

  it('generates interfaces compatible with the legacy ones', () => {
    type ITest_extends_LegacyITest = Foo.ITest extends LegacyFoo.ITest
      ? true
      : false
    type LegacyITest_extends_ITest = LegacyFoo.ITest extends Foo.ITest
      ? true
      : false

    type GeneratesCompatibleInterfaces = ITest_extends_LegacyITest extends true
      ? LegacyITest_extends_ITest extends true
        ? true
        : false
      : false

    const generatesCompatibleInterfaces: GeneratesCompatibleInterfaces = true
    expect(generatesCompatibleInterfaces).toBe(true)
  })
})
