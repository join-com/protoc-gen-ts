import { Root, roots } from 'protobufjs'

roots['decorated'] = new Root()

export { Foo } from './generated/Test'
