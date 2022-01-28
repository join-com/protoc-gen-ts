// The real test is here: we want to ensure that redundant imports don't cause an error.
import { GoogleProtobuf as gp1 } from './generated/google/protobuf/Timestamp'
import { GoogleProtobuf as gp2 } from './generatedRedundant/google/protobuf/Timestamp'

import { GoogleProtobuf as gp3 } from './generated/google/protobuf/Empty'
import { GoogleProtobuf as gp4 } from './generatedRedundant/google/protobuf/Empty'

describe('regressions 03', () => {
  it('dummy test', () => {
    // This test has code only to ensure that the linters don't complain because of unused imports.
    const _ts1 = gp1.Timestamp.fromInterface({ seconds: 0, nanos: 0 })
    const _ts2 = gp2.Timestamp.fromInterface({ seconds: 0, nanos: 0 })

    expect(_ts1.asInterface().seconds).toEqual(_ts2.asInterface().seconds)
    expect(_ts1.asInterface().nanos).toEqual(_ts2.asInterface().nanos)

    const _ts3 = gp3.Empty.fromInterface({})
    const _ts4 = gp4.Empty.fromInterface({ seconds: 0, nanos: 0 })

    expect(_ts3.asInterface()).toEqual({})
    expect(_ts4.asInterface()).toEqual({})
  })
})
