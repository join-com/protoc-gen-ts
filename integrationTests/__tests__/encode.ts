// TODO:
// 1. deprecated log
// 2. test changes in messages adding fields, removing
// 3. test enums changes
// 4. test if not sent

import { Foo } from './generated/Test'
import { loadSync } from 'protobufjs'
import * as path from 'path'

let PbTest: any
beforeAll(() => {
  const root = loadSync(path.join(__dirname, 'proto', 'test.proto'))
  PbTest = root.lookupType('foo.Test')
})

describe('encode', () => {
  let buffer: Uint8Array
  let decoded: any
  const values = {
    fieldInt32: 123,
    fieldInt32Repeated: [123, 12312],
    fieldInt64: 12321313,
    fieldInt64Repeated: [1333023, 12000312],
    fieldDouble: 1234.1221231323,
    fieldDoubleRepeated: [1234.1221231323],
    fieldFloat: 9999.1,
    fieldFloatRepeated: [9999.1],
    fieldUint32: 11119,
    fieldUint32Repeated: [11119],
    fieldSint32: 13123,
    fieldSint32Repeated: [13123, 1312312],
    fieldFixed32: 123123,
    fieldFixed32Repeated: [3223, 12312],
    fieldFixed64: 123123,
    fieldFixed64Repeated: [3223, 12312],
    fieldSfixed32: 123123,
    fieldSfixed32Repeated: [3223, 12312],
    fieldSfixed64: 123123,
    fieldSfixed64Repeated: [3223, 12312],
    fieldBool: true,
    fieldBoolRepeated: [true, false, true],
    fieldString: 'foo',
    fieldStringRepeated: ['foo', 'bar'],
    fieldBytes: new Uint8Array([21, 31]),
    fieldBytesRepeated: [new Uint8Array([21, 31]), new Uint8Array([2, 31])],
    fieldEnum: Foo.EnumType.UNKNOWN,
    fieldEnumRepeated: [Foo.Role.EDIT, Foo.Role.VIEW],
    message: {
      title: 'msg'
    },
    messageRepeated: [
      {
        title: 'msg1'
      },
      {
        title: 'msg2'
      }
    ],
    timestamp: new Date('2014-11-01T12:24:33.123'),
    timestampRepeated: [
      new Date('2014-02-26T12:00:00.234'),
      new Date('2013-05-26T12:00:10.221')
    ]
  }

  beforeEach(() => {
    const user = new Foo.TestMsg(values)
    buffer = user.encode().finish()
    decoded = PbTest.toObject(PbTest.decode(buffer), {
      enums: String,
      longs: Number,
      defaults: false
    })
  })

  describe.each([
    'fieldInt32',
    'fieldInt64',
    'fieldDouble',
    'fieldUint32',
    'fieldSint32',
    'fieldFixed32',
    'fieldFixed64',
    'fieldSfixed32',
    'fieldSfixed64',
    'fieldBool',
    'fieldString',
    'fieldEnum',
    // 'fieldBytes' // -- failed
    'message'
  ])('%s', fieldName => {
    it(`encodes ${fieldName}`, () => {
      expect(decoded[fieldName]).toBeDefined()
      expect(decoded[fieldName]).toEqual((values as any)[fieldName])
    })

    it(`encodes ${fieldName}Repeated`, () => {
      const name = `${fieldName}Repeated`
      expect(decoded[name]).toBeDefined()
      expect(decoded[name]).toEqual((values as any)[name])
    })
  })

  describe('timestamp', () => {
    const toDate = (timestamp: { nanos: number; seconds: number }) =>
      new Date(
        (timestamp.seconds || 0) * 1000 + (timestamp.nanos || 0) / 1000000
      )

    it(`encodes timestamp`, () => {
      expect(decoded.timestamp).toBeDefined()
      expect(toDate(decoded.timestamp)).toEqual(values.timestamp)
    })

    it(`encodes timestampRepeated`, () => {
      expect(decoded.timestampRepeated).toBeDefined()
      expect(decoded.timestampRepeated.map(toDate)).toEqual(
        values.timestampRepeated
      )
    })
  })

  describe('fieldFloat', () => {
    it(`encodes fieldFloat`, () => {
      expect(decoded.fieldFloat).toBeDefined()
      expect(parseFloat(decoded.fieldFloat)).toEqual(
        Math.fround(values.fieldFloat)
      )
    })

    it(`encodes fieldFloatRepeated`, () => {
      expect(decoded.fieldFloatRepeated).toBeDefined()
      expect(decoded.fieldFloatRepeated).toEqual(
        values.fieldFloatRepeated.map(Math.fround)
      )
    })
  })
})
