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

  it('is able to deal with timestamps', () => {
    const testObj: Foo.ITest = {
      timestamp: new Date('2021-05-05 18:44:00'),
    }

    const buffer = Foo.Test.encodePatched(testObj).finish()
    const reconstructed = Foo.Test.decodePatched(buffer)

    expect(reconstructed.timestamp).not.toBeUndefined()
    expect(reconstructed.timestamp).not.toBeNull()
    expect(testObj.timestamp?.getTime()).toEqual(
      reconstructed.timestamp?.getTime()
    )
  })

  it('undefined values are recovered as undefined', () => {
    const original: Foo.ITest = {}
    const buffer = Foo.Test.encodePatched(original).finish()
    const reconstructed = Foo.Test.decodePatched(buffer)

    expect(reconstructed.fieldEnum).toBeUndefined()
    expect(reconstructed.fieldString).toBeUndefined()
    expect(reconstructed.fieldInt32).toBeUndefined()
  })

  it('nested undefined values are recovered as undefined', () => {
    const originalA: Foo.IBigWrapper = {}
    const bufferA = Foo.BigWrapper.encodePatched(originalA).finish()
    const reconstructedA = Foo.BigWrapper.decodePatched(bufferA)

    const originalB: Foo.IBigWrapper = { nestedTest: {} }
    const bufferB = Foo.BigWrapper.encodePatched(originalB).finish()
    const reconstructedB = Foo.BigWrapper.decodePatched(bufferB)

    const originalC: Foo.IBigWrapper = {
      nestedTest: {
        fieldEnum: undefined,
        fieldString: undefined,
        fieldInt32: undefined,
      },
    }
    const bufferC = Foo.BigWrapper.encodePatched(originalC).finish()
    const reconstructedC = Foo.BigWrapper.decodePatched(bufferC)

    expect(reconstructedA.nestedTest?.fieldEnum).toBeUndefined()
    expect(reconstructedA.nestedTest?.fieldString).toBeUndefined()
    expect(reconstructedA.nestedTest?.fieldInt32).toBeUndefined()

    expect(reconstructedB.nestedTest?.fieldEnum).toBeUndefined()
    expect(reconstructedB.nestedTest?.fieldString).toBeUndefined()
    expect(reconstructedB.nestedTest?.fieldInt32).toBeUndefined()

    expect(reconstructedC.nestedTest?.fieldEnum).toBeUndefined()
    expect(reconstructedC.nestedTest?.fieldString).toBeUndefined()
    expect(reconstructedC.nestedTest?.fieldInt32).toBeUndefined()
  })

  it('zero values are recovered as zeros', () => {
    const original: Foo.ITest = {
      fieldInt32: 0,
      fieldEnum: 'UNKNOWN',
      fieldString: '',
    }
    const buffer = Foo.Test.encodePatched(original).finish()
    const reconstructed = Foo.Test.decodePatched(buffer)

    expect(reconstructed.fieldInt32).toBe(0)
    expect(reconstructed.fieldEnum).toBe('UNKNOWN')
    expect(reconstructed.fieldString).toBe('')
  })

  it('enum values are recovered', () => {
    const unknownValue: Foo.ITest = { fieldEnum: 'UNKNOWN' }
    const unknownBuffer = Foo.Test.encodePatched(unknownValue).finish()
    const reconstructedUnknown = Foo.Test.decodePatched(unknownBuffer)

    const adminValue: Foo.ITest = { fieldEnum: 'ADMIN' }
    const adminBuffer = Foo.Test.encodePatched(adminValue).finish()
    const reconstructedAdmin = Foo.Test.decodePatched(adminBuffer)

    const userValue: Foo.ITest = { fieldEnum: 'USER' }
    const userBuffer = Foo.Test.encodePatched(userValue).finish()
    const reconstructedUser = Foo.Test.decodePatched(userBuffer)

    expect(reconstructedUnknown.fieldEnum).toBe('UNKNOWN')
    expect(reconstructedAdmin.fieldEnum).toBe('ADMIN')
    expect(reconstructedUser.fieldEnum).toBe('USER')
  })
})
