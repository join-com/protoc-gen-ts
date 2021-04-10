import { Foo } from './generated/Test'
import { Foo as LegacyFoo } from '../legacy/generated/Test'

describe('(v2) legacy compatibility', () => {
  const request: Foo.IRequest = { id: 42 }
  const nested: Foo.INested = { title: 'The Test' }

  it('encode generates same buffers as legacy generator', () => {
    const requestBuffer = Foo.Request.encodePatched(request).finish()
    const legacyRequestBuffer = new LegacyFoo.Request(request).encode().finish()
    expect(requestBuffer).toStrictEqual(legacyRequestBuffer)

    const nestedBuffer = Foo.Nested.encodePatched(nested).finish()
    const legacyNestedBuffer = new LegacyFoo.Nested(nested).encode().finish()
    expect(nestedBuffer).toStrictEqual(legacyNestedBuffer)

    // TODO: Fix ITimestamp -> Date conversion
    // TODO: Check why serialization is different
  })
})
