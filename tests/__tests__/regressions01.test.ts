// eslint-disable-next-line node/no-unpublished-import
import * as ts from 'typescript'
// eslint-disable-next-line node/no-unpublished-import
import { ClassDeclaration, InterfaceDeclaration, Project } from 'ts-morph'
import { Regressions } from './package1/generated/Regressions'
// eslint-disable-next-line node/no-unpublished-import
import { parse as parseComment } from 'comment-parser'
import { join as pathJoin } from 'path'

describe('regressions 01-02', () => {
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
    const regressionsFile = project.getSourceFile(pathJoin(__dirname, 'package1/generated', 'Regressions.ts'))!
    expect(regressionsFile).toBeDefined()

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const moduleDeclaration = regressionsFile.getFirstChildByKind(ts.SyntaxKind.ModuleDeclaration)!
    expect(moduleDeclaration).toBeDefined()
    expect(moduleDeclaration.getName()).toBe('Regressions')

    // Here is where we verify the fix works as intended.
    // -------------------------------------------------------------------------
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const messageInterfaceWithDeprecatedField = moduleDeclaration.getInterface('IMessageWithDeprecatedField')!
    expect(messageInterfaceWithDeprecatedField).toBeDefined()
    verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(messageInterfaceWithDeprecatedField)
    verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(messageInterfaceWithDeprecatedField)

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const messageClassWithDeprecatedField = moduleDeclaration.getClass('MessageWithDeprecatedField')!
    expect(messageClassWithDeprecatedField).toBeDefined()
    verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(messageClassWithDeprecatedField)
    verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(messageClassWithDeprecatedField)

    // Here we verify that the fix did not introduce any unintended regression.
    // -------------------------------------------------------------------------
    const deprecatedMessageInterfaceWithDeprecatedField =
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      moduleDeclaration.getInterface('IDeprecatedMessageWithDeprecatedField')!
    expect(deprecatedMessageInterfaceWithDeprecatedField).toBeDefined()
    verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(deprecatedMessageInterfaceWithDeprecatedField)
    verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(deprecatedMessageInterfaceWithDeprecatedField)

    const deprecatedMessageClassWithDeprecatedField =
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      moduleDeclaration.getClass('DeprecatedMessageWithDeprecatedField')!
    expect(deprecatedMessageClassWithDeprecatedField).toBeDefined()
    verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(deprecatedMessageClassWithDeprecatedField)
    verifyThatNotDeprecatedFieldDoesNotHaveJSDocAnnotation(deprecatedMessageClassWithDeprecatedField)
  })
})

function verifyThatDeprecatedFieldHasJSDocDeprecationAnnotation(
  generatedInterface: InterfaceDeclaration | ClassDeclaration,
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
  generatedInterface: InterfaceDeclaration | ClassDeclaration,
): void {
  const notDeprecatedField =
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    generatedInterface.getProperty('notDeprecated')!
  expect(notDeprecatedField).toBeDefined()

  const commentRanges = notDeprecatedField.getLeadingCommentRanges()
  expect(commentRanges).toHaveLength(0)
}
