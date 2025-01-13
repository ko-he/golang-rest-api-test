// eslint-disable @typescript-eslint/no-unsafe-enum-comparison

import { promises as fs } from 'fs'
import path from 'path'

import { parse } from '@typescript-eslint/parser'

/**
 * Api の tagTypes を generatedApi から生成する
 */
const generate = async () => {
  const apiFilePath = path.resolve(__dirname, './generatedApi.ts')

  const tagTypesFilePath = path.resolve(__dirname, './ApiTagTypes.ts')

  const tagTypesFileBodyTemplate = (names: string[]) => {
    const arrayItemCode = names.join("',\n  '")

    return [
      '/* eslint-disable */',
      '// prettier-ignore',
      '',
      `export const ApiTagTypes = [\n  '${arrayItemCode}'\n] as const`,
      '',
      'export type ApiTagType = typeof ApiTagTypes[number]',
    ].join('\n')
  }

  /**
   * ApiArg type は除外する
   */
  const nameMatcher = (name: string) => {
    return !name.endsWith('ApiArg')
  }

  const code = await fs.readFile(apiFilePath, 'utf-8')

  const ast = parse(code, {
    loc: true,
    range: true,
  })

  const names = ast.body
    .filter(
      (
        // type alias と interface のみに filtering
        dec
      ) =>
        dec.type === 'ExportNamedDeclaration' &&
        (dec?.declaration?.type === 'TSTypeAliasDeclaration' ||
          dec?.declaration?.type === 'TSInterfaceDeclaration')
    )
    .map(
      (
        // type alias と interface の name を取り出す
        dec
      ) => {
        if (dec.type === 'ExportNamedDeclaration') {
          switch (dec?.declaration?.type) {
            case 'TSTypeAliasDeclaration':

            case 'TSInterfaceDeclaration': {
              const { name } = dec.declaration.id

              return nameMatcher(name) ? name : ''
            }
          }
        }

        return ''
      }
    )
    .filter(Boolean)

  await fs.writeFile(
    tagTypesFilePath,
    new Uint8Array(Buffer.from(tagTypesFileBodyTemplate(names)))
  )
}

generate()
