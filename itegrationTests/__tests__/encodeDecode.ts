// // import { execSync } from 'child_process'
// import { Foo } from './generated/Test'
// import { loadSync } from 'protobufjs'
// import * as path from 'path'

// let UserPbMessage: any
// beforeAll(() => {
//   //execSync('yarn proto:build')
//   const root = loadSync(path.join(__dirname, 'proto', 'test.proto'))
//   UserPbMessage = root.lookupType('foo.User')
// })

// describe('encode', () => {
//   it('encodes correctly', () => {
//     const user = new Foo.UserMsg({
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
//     const buffer = user.encode().finish()
//     const decoded = UserPbMessage.toObject(UserPbMessage.decode(buffer), {
//       enums: String
//     })
//     expect(decoded.id).toEqual(123)
//     expect(decoded.name).toEqual('Joe Doe')
//     expect(decoded.type).toEqual('ADMIN')
//     expect(decoded.roles).toEqual(['EDIT', 'VIEW'])
//     expect(decoded.favoriteBook).toEqual({ title: 'Clean Architecture' })
//     expect(decoded.readBooks).toEqual([
//       { title: 'Clean Architecture' },
//       { title: 'Clean Code', isbn: '112344' }
//     ])
//   })
// })

// describe('decode', () => {
//   it('encodes correctly', () => {
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
