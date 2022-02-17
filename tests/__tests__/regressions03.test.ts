// The real test is here: we want to ensure that redundant imports don't cause an error.
import { Root, roots } from 'protobufjs'

// Required to isolate registered protobuf types between packages.
// By default a single Root is used to register types which may cause a `duplicate name X in Root` error.
roots['decorated'] = new Root()

import { GoogleProtobuf as gp1 } from './generated/google/protobuf/Timestamp'
import { GoogleProtobuf as gp2 } from './generated/google/protobuf/Empty'

// Required to isolate registered protobuf types between packages.
// By default a single Root is used to register types which may cause a `duplicate name X in Root` error.
roots['decorated'] = new Root()

import { GoogleProtobuf as gp3 } from './generatedRedundant/google/protobuf/Timestamp'
import { GoogleProtobuf as gp4 } from './generatedRedundant/google/protobuf/Empty'

describe('regressions 03', () => {
  it('dummy test', () => {
    // This test has code only to ensure that the linters don't complain because of unused imports.
    const _ts1 = gp1.Timestamp.fromInterface({ seconds: 1, nanos: 2 })
    const _ts3 = gp3.Timestamp.fromInterface({ seconds: 1, nanos: 2 })
    expect(_ts1.asInterface().seconds).toEqual(_ts3.asInterface().seconds)
    expect(_ts1.asInterface().nanos).toEqual(_ts3.asInterface().nanos)

    const _ts2 = gp2.Empty.fromInterface({})
    const _ts4 = gp4.Empty.fromInterface({ seconds: 0, nanos: 0 })
    expect(_ts2.asInterface()).toEqual({})
    expect(_ts4.asInterface()).toEqual({})
  })
})
