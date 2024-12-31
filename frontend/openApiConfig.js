
//eslint-disable @typescript-eslint/no-var-requires

import path from 'path'

const projectRoot = path.resolve('../')
const schemaFile = path.resolve(projectRoot, './etc/openapi/root.yml')

/**
 * `yarn install` 無しで code generation を行えるように js の config file を利用
 * @seet {@url https://github.com/reduxjs/redux-toolkit/issues/1775}
 *
 * @type {import("@rtk-query/codegen-openapi").ConfigFile}
 */
const config = {
  schemaFile,

  apiImport: 'emptyApi',
  apiFile: './emptyApi.ts',

  exportName: 'generatedApi',
  outputFile: './generatedApi.ts',

  /**
   * User API 以外を除外
   */
  filterEndpoints: /^(get|put|patch|post|option|delete)User/,

  /**
   * openapi の path の tags を provideTags や invlidateTags に利用しない
   */
  tag: false,

  /**
   * endpointName を利用した hooks を export するかどうか
   */
  hooks: {
    queries: true,
    mutations: true,
    lazyQueries: true,
  },

  /**
   * path , query , body の parameter を flat な object にするかどうか
   */
  flattenArg: false,

  /**
   * optional な type に `| undefined` と明示的に `undefined` を指定
   */
  unionUndefined: true,

  /**
   * mutation ↔︎ query を切り替えたい時に使う
   * 以下コマンドでre-generateを行う
   * npx @rtk-query/codegen-openapi ./src/app/store/services/musubuRev2Api/openapi-config.js
   */
  endpointOverrides: [
    {
      pattern: 'postUserPurchaseDetailWithClientSearch',
      type: 'query',
    },
    {
      pattern: 'postUserPurchaseDetailWithClients',
      type: 'query',
    },
    {
      pattern: 'postUserMailRecipientCount',
      type: 'query',
    },
  ],
}

export default config
