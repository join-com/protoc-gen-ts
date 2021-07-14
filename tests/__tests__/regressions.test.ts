import { Regressions } from './generated/Regressions'

describe('regressions', () => {
  it('01. decodes missing nested fields as undefined when there are no enums', () => {
    // Background: before this test was written, calling .asInterface() for nested objects was
    //             done only when there were enum fields at some level of the object.

    const originalA: Regressions.IReg01Outer = {}
    const bufferA = Regressions.Reg01Outer.encodePatched(originalA).finish()
    const reconstructedA = Regressions.Reg01Outer.decodePatched(bufferA)

    const originalB: Regressions.IReg01Outer = { inner: {} }
    const bufferB = Regressions.Reg01Outer.encodePatched(originalB).finish()
    const reconstructedB = Regressions.Reg01Outer.decodePatched(bufferB)

    expect(reconstructedA.inner?.value).toBeUndefined()
    expect(reconstructedB.inner?.value).toBeUndefined()
  })
})
