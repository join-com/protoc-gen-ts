/**
 * The real test is here: we want to ensure that multiple imports having google Timestamp field do not throw
 * `duplicate name 'Timestamp' in Root` error
 */
import { Companies } from './package2'
import { Foo } from './package1'
import { Jobs } from './package3'

describe('regressions 03', () => {
  it('dummy test', () => {
    // This test has code only to ensure that the linters don't complain because of unused imports.
    const _ts1 = Companies.Company.fromInterface({ id: 1 })
    expect(_ts1.asInterface()).toEqual({ id: 1 })

    const _ts2 = Jobs.Job.fromInterface({ id: 2 })
    expect(_ts2.asInterface()).toEqual({ id: 2 })

    const _ts3 = Foo.Test.fromInterface({ fieldInt32: 3 })
    expect(_ts3.asInterface()).toEqual({ fieldInt32: 3 })
  })
})
