import { Foo } from './generated/Test'
import { Foo as LegacyFoo } from '../legacy/generated/Test'

describe('(v2) classes', () => {
  const request: Foo.IRequest = { id: 42 }
  const nested: Foo.INested = { title: 'The Test' }
  const complex: Foo.ITest = {
    fieldEnum: 'ADMIN',
    fieldEnumRepeated: ['VIEW', 'EDIT'],
    message: nested,
    messageRepeated: [nested, nested],
    fieldBytes: Buffer.from('Ubik'),
  }

  it('decodePatched(encodePatched(*)) === identity', () => {
    const requestBuffer = Foo.Request.encodePatched(request).finish()
    const reconstructedRequest = Foo.Request.decodePatched(requestBuffer)
    expect(reconstructedRequest).toEqual(request)

    const nestedBuffer = Foo.Nested.encodePatched(nested).finish()
    const reconstructedNested = Foo.Nested.decodePatched(nestedBuffer)
    expect(reconstructedNested).toEqual(nested)

    const complexBuffer = Foo.Test.encodePatched(complex).finish()
    const reconstructedComplex = Foo.Test.decodePatched(complexBuffer)

    // The reconstructed object will have some "extra" fields (with undefined value,
    // and defined in the interface).
    expect(reconstructedComplex).toMatchObject(complex)
  })

  it('encode generates same buffers as legacy generator', () => {
    const requestBuffer = Foo.Request.encodePatched(request).finish()
    const legacyRequestBuffer = new LegacyFoo.Request(request).encode().finish()
    expect(requestBuffer).toStrictEqual(legacyRequestBuffer)

    const nestedBuffer = Foo.Nested.encodePatched(nested).finish()
    const legacyNestedBuffer = new LegacyFoo.Nested(nested).encode().finish()
    expect(nestedBuffer).toStrictEqual(legacyNestedBuffer)

    // TODO: Fix ITimestamp -> Date conversion
    // TODO: Check why serialization is different
    // const complexBuffer = Foo.Test.encodePatched(complex).finish()
    // const legacyComplexBuffer = new LegacyFoo.Test(complex as LegacyFoo.ITest).encode().finish()
    // expect(complexBuffer).toStrictEqual(legacyComplexBuffer)
  })
})
