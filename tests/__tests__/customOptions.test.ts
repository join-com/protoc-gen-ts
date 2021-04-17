import { Common } from './generated/common/Extra'
import { Foo } from './generated/Test'

describe('(v2) custom options', () => {
  it('[typescript_required] causes the field to not extend undefined', () => {
    type IsFieldRequired = Foo.ICustomOptionsTest['requiredField'] extends Common.IExtraPkgMessage
      ? true
      : false
    const isFieldRequired: IsFieldRequired = true
    expect(isFieldRequired).toBe(true)
  })

  it('[typescript_required_fields] causes message fields to not extend undefined', () => {
    type IsFieldRequired = Foo.IRequiredPropertiesTest['requiredField'] extends number
      ? true
      : false
    const isFieldRequired: IsFieldRequired = true
    expect(isFieldRequired).toBe(true)
  })

  it('[typescript_optional] makes field optional for [typescript_required_fields] messages', () => {
    type IsFieldOptional = number | undefined extends Foo.IRequiredPropertiesTest['optionalField']
      ? true
      : false
    const isFieldOptional: IsFieldOptional = true
    expect(isFieldOptional).toBe(true)
  })
})
