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
    fieldFloat: 9999.1223,
    fieldFloatRepeated: [9999.1223],
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
    fieldBoolRepeated: [true, false, true]
  }

  beforeEach(() => {
    const user = new Foo.TestMsg(values)
    buffer = user.encode().finish()
    decoded = PbTest.toObject(PbTest.decode(buffer), {
      enums: String,
      longs: Number
    })
  })

  describe.each([
    'fieldInt32',
    'fieldInt64',
    'fieldDouble',
    // 'fieldFloat', -- failed, needs to be checked with protobufjs
    'fieldUint32',
    'fieldSint32',
    'fieldFixed32',
    'fieldFixed64',
    'fieldSfixed32',
    'fieldSfixed64',
    'fieldBool'
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
})

// describe('decode', () => {
//   it('decode correctly', () => {
//     const user = UserPbMessage.fromObject({
//       id: 123,
//       name: 'Joe Doe',
//       type: Foo.Type.ADMIN,
//       roles: [Foo.Role.EDIT, Foo.Role.VIEW],
//       favoriteBook: { title: 'Clean Architecture' },
//       readBooks: [
//         { title: 'Clean Architecture' },
//         { title: 'Clean Code', isbn: '112344' }
//       ]
//     })
//     const buffer = UserPbMessage.encode(user).finish()
//     const decoded = Foo.UserMsg.decode(buffer)
//     expect(decoded.id).toEqual(123)
//     expect(decoded.name).toEqual('Joe Doe')
//     expect(decoded.type).toEqual('ADMIN')
//     expect(decoded.roles).toEqual(['EDIT', 'VIEW'])
//     expect(decoded.favoriteBook.title).toEqual('Clean Architecture')
//     expect(
//       decoded.readBooks.map(({ title, isbn }) => ({
//         title,
//         isbn
//       }))
//     ).toEqual([
//       { title: 'Clean Architecture' },
//       { title: 'Clean Code', isbn: '112344' }
//     ])
//   })
// })
