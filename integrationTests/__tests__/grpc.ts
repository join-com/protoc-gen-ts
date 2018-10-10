import * as grpc from 'grpc'
import { Foo } from './generated/Test'

type handleUnaryCallPromise<RequestType, ResponseType> = (
  call: grpc.ServerUnaryCall<RequestType>
) => Promise<ResponseType>

type handleClientStreamingCallPromise<RequestType, ResponseType> = (
  call: grpc.ServerReadableStream<RequestType>
) => Promise<ResponseType>

type handleCall<RequestType, ResponseType> =
  | grpc.handleCall<RequestType, ResponseType>
  | handleUnaryCallPromise<RequestType, ResponseType>
  | handleClientStreamingCallPromise<RequestType, ResponseType>

interface Implementations {
  [key: string]: handleCall<any, any>
}

interface UsersServerImplementation extends Implementations {
  find(call: grpc.ServerUnaryCall<Foo.Request>): Promise<Foo.User>
  find(
    call: grpc.ServerUnaryCall<Foo.Request>,
    callback: grpc.sendUnaryData<Foo.User>
  ): void
  findClientStream(
    call: grpc.ServerReadableStream<Foo.Request>
  ): Promise<Foo.User>
  findClientStream(
    call: grpc.ServerReadableStream<Foo.Request>,
    callback: grpc.sendUnaryData<Foo.User>
  ): void
  // findServerStream(call: grpc.ServerWriteableStream<Foo.Request>): void
  // findBidiStream(call: grpc.ServerDuplexStream<Foo.Request, Foo.User>): void
}

class Service<I extends Implementations> {
  public readonly definitions: grpc.ServiceDefinition<any>
  public readonly grpcImplementations: grpc.UntypedServiceImplementation
  constructor(
    rawDefinitions: grpc.ServiceDefinition<any>,
    public readonly implementations: I
  ) {
    this.definitions = this.addLogging(rawDefinitions)
    this.grpcImplementations = this.convertToGrpcImplementation(
      this.implementations
    )
  }

  private addLogging(definitions: grpc.ServiceDefinition<any>) {
    const loggedDefinitions = Object.entries<grpc.MethodDefinition<any, any>>(
      definitions
    ).reduce((acc, [name, definition]) => {
      const loggedDefinition: grpc.MethodDefinition<any, any> = {
        ...definition,
        requestDeserialize: argBuf => {
          const request = definition.requestDeserialize(argBuf)
          console.log('Request', definition.path, request)
          return request
        },
        responseSerialize: args => {
          console.log('Response', definition.path, args)
          return definition.responseSerialize(args)
        }
      }
      return { ...acc, [name]: loggedDefinition }
    }, {}) as grpc.ServiceDefinition<any>
    return loggedDefinitions
  }

  private convertToGrpcImplementation(
    implementations: I
  ): grpc.UntypedServiceImplementation {
    return Object.entries(implementations).reduce(
      (acc, [name, implementation]) => {
        let newImplementation: handleCall<any, any>
        if (
          this.definitions[name].responseStream &&
          !this.definitions[name].requestStream
        ) {
          newImplementation = implementation
        } else if (
          this.definitions[name].responseStream &&
          this.definitions[name].requestStream
        ) {
          newImplementation = implementation
        } else if (
          !this.definitions[name].responseStream &&
          this.definitions[name].requestStream
        ) {
          if (implementation.length === 2) {
            newImplementation = implementation
          } else {
            newImplementation = async (
              call: grpc.ServerReadableStream<any>,
              callback: grpc.sendUnaryData<any>
            ) => {
              try {
                const result = await (implementation as handleClientStreamingCallPromise<
                  any,
                  any
                >)(call)
                callback(null, result)
              } catch (e) {
                // TODO: handle Error
                // callback(null, result)
              }
            }
          }
        } else {
          if (implementation.length === 2) {
            newImplementation = implementation
          } else {
            newImplementation = async (
              call: grpc.ServerUnaryCall<any>,
              callback: grpc.sendUnaryData<any>
            ) => {
              try {
                const result = await (implementation as handleUnaryCallPromise<
                  any,
                  any
                >)(call)
                callback(null, result)
              } catch (e) {
                // TODO: handle Error
                // callback(null, result)
              }
            }
          }
        }
        return { ...acc, [name]: newImplementation }
      },
      {}
    ) as grpc.UntypedServiceImplementation
  }
}

// class UserService
//   extends Service<Foo.UsersServerImplementation, UsersServerImplementation>
//   implements UsersServerImplementation {
//   public async find(call: grpc.ServerUnaryCall<Foo.Request>): void {}
// }

// (args: User) =>
//         new UserMsg(args).encode().finish() as Buffer,

// (argBuf: Buffer) =>
//         RequestMsg.decode(argBuf) as Request,

describe('grpc', () => {
  let server: grpc.Server
  let client: grpc.Client
  beforeEach(() => {
    server = new grpc.Server()
    const service = new Service<UsersServerImplementation>(
      Foo.usersServiceDefinition,
      {
        find: async (call: grpc.ServerUnaryCall<Foo.Request>) => ({ id: 11 }),
        findClientStream: async (
          call: grpc.ServerReadableStream<Foo.Request>
        ) => ({ id: 11 })
      }
    )

    const fn = async ({ request }) => ({ id: request.id })
    server.addService<Foo.UsersServerImplementation>(
      Foo.usersServiceDefinition,
      {
        find: (call, callback) => {
          fn(call)
            .then(result => callback(null, result))
            .catch(e => {
              const metadata = new grpc.Metadata()
              metadata.add('error', JSON.stringify(e))
              callback(
                {
                  code: grpc.status.UNKNOWN,
                  message: e.message,
                  name: 'aaa',
                  metadata
                },
                null
              )
            })
        }
      }
    )

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
