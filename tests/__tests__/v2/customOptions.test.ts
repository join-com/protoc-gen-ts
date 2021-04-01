import { Foo } from './generated/Test'

describe('(v2) custom options', () => {
  it('[typescript_required] causes the field to not extend undefined', () => {
    type IsFieldRequired = Foo.ICustomOptionsTest['requiredField'] extends undefined
      ? false
      : true
    const isFieldRequired: IsFieldRequired = true
    expect(isFieldRequired).toBe(true)
  })
})
