import { Server, grpc } from '@join-com/grpc'
import { Foo } from './generated/Test'

class SimpleTestService extends Foo.AbstractSimpleTestService {
  public async ForwardParameter(
    call: grpc.ServerUnaryCall<Foo.IBigWrapper, Foo.IBigWrapper>
  ): Promise<Foo.IBigWrapper> {
    return Promise.resolve(call.request)
  }
  public async GetEmptyResult(): Promise<Foo.IBigWrapper> {
    return Promise.resolve({})
  }
}

describe('(v2) generated services', () => {
  let server: Server
  let client: Foo.ISimpleTestClient

  beforeAll(async () => {
    server = new Server()
    server.addService(new SimpleTestService())
    await server.start('0.0.0.0:0')

    client = new Foo.SimpleTestClient({
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      address: `0.0.0.0:${server.port!}`,
    })
  })

  afterAll(async () => {
    client.close()
    await server.tryShutdown()
  })

  it('service -> client does not alter encode/decode semantics', async () => {
    const emptyResult = await client.getEmptyResult({}).res
    const forwardedResult = await client.forwardParameter({
      nestedTest: {
        fieldString: '',
        fieldInt32: 0,
        fieldEnum: 'UNKNOWN',
      },
    }).res

    expect(emptyResult.nestedTest?.fieldString).toBeUndefined()
    expect(emptyResult.nestedTest?.fieldInt32).toBeUndefined()
    expect(emptyResult.nestedTest?.fieldEnum).toBeUndefined()

    expect(forwardedResult.nestedTest?.fieldString).toEqual('')
    expect(forwardedResult.nestedTest?.fieldInt32).toEqual(0)
    expect(forwardedResult.nestedTest?.fieldEnum).toEqual('UNKNOWN')
  })
})
