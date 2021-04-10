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

    // TODO: Check why serialization is different
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
