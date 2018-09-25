import * as grpc from 'grpc'
import { Foo } from './generated/Test'

describe('grpc', () => {
  let server: grpc.Server
  let client: grpc.Client
  beforeEach(() => {
    server = new grpc.Server()
    server.addService(Foo.usersServiceDefinition, {
      find: (call, callback) => {
        console.log(call)
        callback(null, { id: call.request.id })
      }
    })
    const port = server.bind(
      '0.0.0.0:0',
      grpc.ServerCredentials.createInsecure()
    )
    server.start()
    const UsersClient = grpc.makeGenericClientConstructor(
      Foo.usersServiceDefinition,
      'Users',
      {}
    )
    client = new UsersClient(
      `0.0.0.0:${port}`,
      grpc.credentials.createInsecure()
    )
  })
  afterEach(() => {
    client.close()
    server.forceShutdown()
  })

  it('should ', done => {
    client['find']({ id: 11 }, (err, response) => {
      console.log(response.id)
      done()
    })
  })
})
