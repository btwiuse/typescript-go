
currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::--p project
//// [/home/src/workspaces/solution/project/index.ts] new file
import { x } from "../utils";
//// [/home/src/workspaces/solution/project/tsconfig.json] new file
{
		"references": [
			{ "path": "../utils" },
		],
	}),
},
//// [/home/src/workspaces/solution/src/utils/index.ts] new file
export const x = 10;
//// [/home/src/workspaces/solution/src/utils/tsconfig.json] new file
{
	"compilerOptions": {
		"composite": true,
		"noEmit": true,
	},
})

ExitStatus:: 2

CompilerOptions::{
    "allowJs": null,
    "allowArbitraryExtensions": null,
    "allowSyntheticDefaultImports": null,
    "allowImportingTsExtensions": null,
    "allowNonTsExtensions": null,
    "allowUmdGlobalAccess": null,
    "allowUnreachableCode": null,
    "allowUnusedLabels": null,
    "assumeChangesOnlyAffectDirectDependencies": null,
    "alwaysStrict": null,
    "baseUrl": "",
    "build": null,
    "checkJs": null,
    "customConditions": null,
    "composite": null,
    "emitDeclarationOnly": null,
    "emitBOM": null,
    "emitDecoratorMetadata": null,
    "downlevelIteration": null,
    "declaration": null,
    "declarationDir": "",
    "declarationMap": null,
    "disableSizeLimit": null,
    "disableSourceOfProjectReferenceRedirect": null,
    "disableSolutionSearching": null,
    "disableReferencedProjectLoad": null,
    "esModuleInterop": null,
    "exactOptionalPropertyTypes": null,
    "experimentalDecorators": null,
    "forceConsistentCasingInFileNames": null,
    "isolatedModules": null,
    "isolatedDeclarations": null,
    "ignoreDeprecations": "",
    "importHelpers": null,
    "inlineSourceMap": null,
    "inlineSources": null,
    "init": null,
    "incremental": null,
    "jsx": 0,
    "jsxFactory": "",
    "jsxFragmentFactory": "",
    "jsxImportSource": "",
    "keyofStringsOnly": null,
    "lib": null,
    "locale": "",
    "mapRoot": "",
    "module": 0,
    "moduleResolution": 0,
    "moduleSuffixes": null,
    "moduleDetectionKind": 0,
    "newLine": 0,
    "noEmit": null,
    "noCheck": null,
    "noErrorTruncation": null,
    "noFallthroughCasesInSwitch": null,
    "noImplicitAny": null,
    "noImplicitThis": null,
    "noImplicitReturns": null,
    "noEmitHelpers": null,
    "noLib": null,
    "noPropertyAccessFromIndexSignature": null,
    "noUncheckedIndexedAccess": null,
    "noEmitOnError": null,
    "noUnusedLocals": null,
    "noUnusedParameters": null,
    "noResolve": null,
    "noImplicitOverride": null,
    "noUncheckedSideEffectImports": null,
    "out": "",
    "outDir": "",
    "outFile": "",
    "paths": null,
    "preserveConstEnums": null,
    "preserveSymlinks": null,
    "project": "/home/src/workspaces/solution/project",
    "resolveJsonModule": null,
    "resolvePackageJsonExports": null,
    "resolvePackageJsonImports": null,
    "removeComments": null,
    "rewriteRelativeImportExtensions": null,
    "reactNamespace": "",
    "rootDir": "",
    "rootDirs": null,
    "skipLibCheck": null,
    "strict": null,
    "strictBindCallApply": null,
    "strictBuiltinIteratorReturn": null,
    "strictFunctionTypes": null,
    "strictNullChecks": null,
    "strictPropertyInitialization": null,
    "stripInternal": null,
    "skipDefaultLibCheck": null,
    "sourceMap": null,
    "sourceRoot": "",
    "suppressOutputPathCheck": null,
    "target": 0,
    "traceResolution": null,
    "tsBuildInfoFile": "",
    "typeRoots": null,
    "types": null,
    "useDefineForClassFields": null,
    "useUnknownInCatchVariables": null,
    "verbatimModuleSyntax": null,
    "maxNodeModuleJsDepth": null,
    "configFilePath": "",
    "noDtsResolution": null,
    "pathsBasePath": "",
    "diagnostics": null,
    "extendedDiagnostics": null,
    "generateCpuProfile": "",
    "generateTrace": "",
    "listEmittedFiles": null,
    "listFiles": null,
    "explainFiles": null,
    "listFilesOnly": null,
    "noEmitForJsFiles": null,
    "preserveWatchOutput": null,
    "pretty": null,
    "version": null,
    "watch": null,
    "showConfig": null,
    "tscBuild": null
}
Output::
project/index.ts(1,19): error TS2307: Cannot find module '../utils' or its corresponding type declarations.


Found 1 error in project/index.ts[90m:1[0m

//// [/home/src/workspaces/solution/project/index.js] new file

//// [/home/src/workspaces/solution/project/index.ts] no change
//// [/home/src/workspaces/solution/project/tsconfig.json] no change
//// [/home/src/workspaces/solution/src/utils/index.ts] no change
//// [/home/src/workspaces/solution/src/utils/tsconfig.json] no change

