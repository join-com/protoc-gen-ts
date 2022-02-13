import { Common } from './package1/generated/common/Extra'
import { Flavors } from './package1/generated/Flavors'
import { Foo } from './package1/generated/Test'
import { WithFlavor } from '@coderspirit/nominal'

describe('(v2) custom options', () => {
  type TestUserId = WithFlavor<number, 'UserId'>
  it('[typescript_required] causes the field to not extend undefined', () => {
    type IsFieldRequired = Foo.ICustomOptionsTest['requiredField'] extends Common.IExtraPkgMessage ? true : false
    const isFieldRequired: IsFieldRequired = true
    expect(isFieldRequired).toBe(true)
  })

  it('[typescript_required_fields] causes message fields to not extend undefined', () => {
    type IsFieldRequired = Foo.IRequiredPropertiesTest['requiredField'] extends number ? true : false
    const isFieldRequired: IsFieldRequired = true
    expect(isFieldRequired).toBe(true)
  })

  it('[typescript_optional] makes field optional for [typescript_required_fields] messages', () => {
    type IsFieldOptional = number | undefined extends Foo.IRequiredPropertiesTest['optionalField'] ? true : false
    const isFieldOptional: IsFieldOptional = true
    expect(isFieldOptional).toBe(true)
  })

  it('[typescript_flavor] generates flavor type', () => {
    type TestUserId_extends_UserId = TestUserId extends Flavors.UserId ? true : false
    type UserId_extends_TestUserId = Flavors.UserId extends TestUserId ? true : false

    type UserIdHasFlavor = TestUserId_extends_UserId extends true
      ? UserId_extends_TestUserId extends true
        ? true
        : false
      : false

    const _UserIdHasFlavor: UserIdHasFlavor = true
    expect(_UserIdHasFlavor).toBe(true)
  })

  it('[typescript_flavor] refines field types', () => {
    type UserProfileId = Flavors.UserProfile['id']

    type TestUserId_extends_UserProfileId = TestUserId extends UserProfileId ? true : false
    type UserProfileId_extends_TestUserId = UserProfileId extends TestUserId ? true : false

    type UserProfileIdHasFlavor = TestUserId_extends_UserProfileId extends true
      ? UserProfileId_extends_TestUserId extends true
        ? true
        : false
      : false

    const _UserProfileIdHasFlavor: UserProfileIdHasFlavor = true
    expect(_UserProfileIdHasFlavor).toBe(true)
  })
})
