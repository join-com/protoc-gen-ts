import { Foo } from './generated/Test'

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
})
