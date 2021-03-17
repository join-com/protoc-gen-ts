import { Foo } from './generated/Test'
import { loadSync } from 'protobufjs'
import * as path from 'path'

const baseValues = {
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
  fieldUint64: 114119,
  fieldUint64Repeated: [161119],
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
  fieldEnum: 'UNKNOWN',
  fieldEnumRepeated: ['EDIT', 'VIEW'],
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
  timestamp: { seconds: 1414841073, nanos: 123000000 },
  timestampRepeated: [
    { seconds: 1393412400, nanos: 234000000 },
    { seconds: 1369562410, nanos: 221000000 }
  ]
}

describe('decode', () => {
  const values = baseValues

  const root = loadSync(path.join(__dirname, 'proto', 'test.proto'))
  const PbTest = root.lookupType('foo.Test')
  let buffer: Uint8Array
  let decoded: any

  beforeEach(() => {
    const message = PbTest.fromObject(values)
    buffer = PbTest.encode(message).finish()
    decoded = Foo.Test.decode(buffer)
  })

  describe.each([
    'fieldInt32',
    'fieldInt64',
    'fieldDouble',
    'fieldUint32',
    'fieldUint64',
    'fieldSint32',
    'fieldFixed32',
    'fieldFixed64',
    'fieldSfixed32',
    'fieldSfixed64',
    'fieldBool',
    'fieldString',
    'fieldEnum',
    'fieldBytes',
    'message'
  ])('%s', fieldName => {
    it(`decodes ${fieldName}`, () => {
      expect(decoded[fieldName]).toBeDefined()
      expect(decoded[fieldName]).toEqual((values as any)[fieldName])
    })

    it(`decodes ${fieldName}Repeated`, () => {
      const name = `${fieldName}Repeated`
      expect(decoded[name]).toBeDefined()
      expect(decoded[name]).toEqual((values as any)[name])
    })
  })

  describe('timestamp', () => {
    it('encodes timestamp', () => {
      expect(decoded.timestamp).toBeDefined()
      expect(decoded.timestamp).toEqual(new Date('2014-11-01T11:24:33.123Z'))
    })

    it('encodes timestampRepeated', () => {
      expect(decoded.timestampRepeated).toBeDefined()
      expect(decoded.timestampRepeated).toEqual([
        new Date('2014-02-26T11:00:00.234Z'),
        new Date('2013-05-26T10:00:10.221Z')
      ])
    })
  })

  describe('fieldFloat', () => {
    it('encodes fieldFloat', () => {
      expect(decoded.fieldFloat).toBeDefined()
      expect(parseFloat(decoded.fieldFloat)).toEqual(
        Math.fround(values.fieldFloat)
      )
    })

    it('encodes fieldFloatRepeated', () => {
      expect(decoded.fieldFloatRepeated).toBeDefined()
      expect(decoded.fieldFloatRepeated).toEqual(
        values.fieldFloatRepeated.map(Math.fround)
      )
    })
  })
})
