// eslint-disable-next-line node/no-unpublished-import
import * as ts from 'typescript'
// eslint-disable-next-line node/no-unpublished-import
import { InterfaceDeclaration, Project } from 'ts-morph'
import { Regressions } from './generated/Regressions'
// eslint-disable-next-line node/no-unpublished-import
import { parse as parseComment } from 'comment-parser'
import { join as pathJoin } from 'path'

describe('regressions', () => {
  it('01. decodes missing nested fields as undefined when there are no enums', () => {
    // Background: before this test was written, calling .asInterface() for nested objects was
    //             done only when there were enum fields at some level of the object.

    const originalA: Regressions.IReg01Outer = {}
    const bufferA = Regressions.Reg01Outer.encodePatched(originalA).finish()
    const reconstructedA = Regressions.Reg01Outer.decodePatched(bufferA)

    const originalB: Regressions.IReg01Outer = { inner: {} }
    const bufferB = Regressions.Reg01Outer.encodePatched(originalB).finish()
    const reconstructedB = Regressions.Reg01Outer.decodePatched(bufferB)

    expect(reconstructedA.inner?.value).toBeUndefined()
    expect(reconstructedB.inner?.value).toBeUndefined()
  })

  it('02. adds @deprecated jsdoc annotation to fields marked as deprecated in proto files', () => {
    // Background: before this test was written, these annotations were only added when the message
    //             itself was deprecated as well.

    const tsConfigPath = pathJoin(__dirname, '..', 'tsconfig.json')
    const project = new Project({ tsConfigFilePath: tsConfigPath })

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const regressionsFile = project.getSourceFile(
      pathJoin(__dirname, 'generated', 'Regressions.ts')
    )!
    expect(regressionsFile).toBeDefined()

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const moduleDeclaration = regressionsFile.getFirstChildByKind(
      ts.SyntaxKind.ModuleDeclaration
    )!
    expect(moduleDeclaration).toBeDefined()
    expect(moduleDeclaration.getName()).toBe('Regressions')

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const withDeprecatedFieldInterface = moduleDeclaration.getInterface(
      'IWithDeprecatedField'
    )!
    expect(withDeprecatedFieldInterface).toBeDefined()
    verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(
      withDeprecatedFieldInterface
    )
    verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(
      withDeprecatedFieldInterface
    )

    const deprecatedWithDeprecatedFieldInterface =
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      moduleDeclaration.getInterface('IDeprecatedWithDeprecatedField')!
    expect(deprecatedWithDeprecatedFieldInterface).toBeDefined()
    verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(
      deprecatedWithDeprecatedFieldInterface
    )
    verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(
      deprecatedWithDeprecatedFieldInterface
    )
  })
})

function verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(
  generatedInterface: InterfaceDeclaration
): void {
  const deprecatedField =
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    generatedInterface.getProperty('deprecated')!
  expect(deprecatedField).toBeDefined()

  const commentRanges = deprecatedField.getLeadingCommentRanges()
  expect(commentRanges).toHaveLength(1)

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const parsedComment = parseComment(commentRanges[0]!.getText())
  expect(parsedComment).toHaveLength(1)

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const commentBlock = parsedComment[0]!
  expect(commentBlock.tags).toHaveLength(1)

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const tagSpec = commentBlock.tags[0]!
  expect(tagSpec.tag).toBe('deprecated')
}

function verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(
  generatedInterface: InterfaceDeclaration
): void {
  const notDeprecatedField =
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    generatedInterface.getProperty('notDeprecated')!
  expect(notDeprecatedField).toBeDefined()

  const commentRanges = notDeprecatedField.getLeadingCommentRanges()
  expect(commentRanges).toHaveLength(0)
}
