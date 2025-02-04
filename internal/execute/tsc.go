package execute

import (
	"fmt"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/compiler/diagnostics"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/tsoptions"
	"github.com/microsoft/typescript-go/internal/tspath"
)

type cbType = func(p any) any

func CommandLine(sys System, cb cbType, commandLineArgs []string) ExitStatus {
	parsedCommandLine := tsoptions.ParseCommandLine(commandLineArgs, sys)
	return executeCommandLineWorker(sys, cb, parsedCommandLine)
}

func executeCommandLineWorker(sys System, cb cbType, commandLine *tsoptions.ParsedCommandLine) ExitStatus {
	configFileName := ""
	reportDiagnostic := createDiagnosticReporter(sys, commandLine.CompilerOptions().Pretty)
	// if commandLine.Options().Locale != nil

	if len(commandLine.Errors) > 0 {
		for _, e := range commandLine.Errors {
			reportDiagnostic(e)
		}
		return ExitStatusDiagnosticsPresent_OutputsSkipped
	}

	// if commandLine.Options().Init != nil
	// if commandLine.Options().Version != nil
	// if commandLine.Options().Help != nil || commandLine.Options().All != nil
	// if commandLine.Options().Watch != nil  && commandLine.Options().ListFilesOnly

	if commandLine.CompilerOptions().Project != "" {
		if len(commandLine.FileNames()) != 0 {
			reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Option_project_cannot_be_mixed_with_source_files_on_a_command_line))
			return ExitStatusDiagnosticsPresent_OutputsSkipped
		}

		fileOrDirectory := tspath.NormalizePath(commandLine.CompilerOptions().Project)
		if fileOrDirectory != "" || sys.FS().DirectoryExists(fileOrDirectory) {
			configFileName = tspath.CombinePaths(fileOrDirectory, "tsconfig.json")
			if !sys.FS().FileExists(configFileName) {
				reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Cannot_find_a_tsconfig_json_file_at_the_current_directory_Colon_0, configFileName))
				return ExitStatusDiagnosticsPresent_OutputsSkipped
			}
		} else {
			configFileName = fileOrDirectory
			if !sys.FS().FileExists(configFileName) {
				reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.The_specified_path_does_not_exist_Colon_0, fileOrDirectory))
				return ExitStatusDiagnosticsPresent_OutputsSkipped
			}
		}
	} else if len(commandLine.FileNames()) == 0 {
		searchPath := tspath.NormalizePath(sys.GetCurrentDirectory())
		configFileName = findConfigFile(searchPath, sys.FS().FileExists, "tsconfig.json")
	}

	if configFileName == "" && len(commandLine.FileNames()) == 0 {
		if commandLine.CompilerOptions().ShowConfig.IsTrue() {
			reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Cannot_find_a_tsconfig_json_file_at_the_current_directory_Colon_0, tspath.NormalizePath(sys.GetCurrentDirectory())))
		} else {
			// print version
			// print help
		}
		return ExitStatusDiagnosticsPresent_OutputsSkipped
	}

	// !!! convert to options with absolute paths is usualy done here, but for ease of implementation, it's done in `tsoptions.ParseCommandLine`
	compilerOptionsFromCommandLine := commandLine.CompilerOptions()

	if configFileName != "" {
		extendedConfigCache := map[string]*tsoptions.ExtendedConfigCacheEntry{}
		configParseResult, errors := getParsedCommandLineOfConfigFile(configFileName, compilerOptionsFromCommandLine, sys, extendedConfigCache)
		if len(errors) != 0 {
			// these are unrecoverable errors--exit to report them as diagnotics
			for _, e := range errors {
				reportDiagnostic(e)
			}
			return ExitStatusDiagnosticsPresent_OutputsGenerated
		}
		// if commandLineOptions.ShowConfig
		// updateReportDiagnostic
		if isWatchSet(configParseResult.CompilerOptions()) {
			// todo watch
			// return ExitStatusDiagnosticsPresent_OutputsSkipped
			return ExitStatusNotImplementedWatch
		} else if isIncrementalCompilation(configParseResult.CompilerOptions()) {
			// todo performIncrementalCompilation
			return ExitStatusNotImplementedIncremental
		} else {
			performCompilation(
				sys,
				cb,
				configParseResult,
				reportDiagnostic,
			)
		}
	} else {
		if compilerOptionsFromCommandLine.ShowConfig.IsTrue() {
			// write show config
			return ExitStatusNotImplementedIncremental
		}
		// todo update reportDiagnostic
		if isWatchSet(compilerOptionsFromCommandLine) {
			// todo watch
			// return ExitStatusDiagnosticsPresent_OutputsSkipped
			return ExitStatusNotImplementedWatch
		} else if isIncrementalCompilation(compilerOptionsFromCommandLine) {
			// todo incremental
			return ExitStatusNotImplementedIncremental
		} else {
			commandLine.SetCompilerOptions(compilerOptionsFromCommandLine)
			performCompilation(
				sys,
				cb,
				commandLine,
				reportDiagnostic,
			)
		}
	}

	return ExitStatusSuccess
}

func findConfigFile(searchPath string, fileExists func(string) bool, configName string) string {
	result, ok := tspath.ForEachAncestorDirectory(searchPath, func(ancestor string) (string, bool) {
		fullConfigName := tspath.CombinePaths(ancestor, configName)
		if fileExists(fullConfigName) {
			return fullConfigName, true
		}
		return fullConfigName, false
	})
	if !ok {
		return ""
	}
	return result
}

// Reads the config file and reports errors. Exits if the config file cannot be found
func getParsedCommandLineOfConfigFile(configFileName string, options *core.CompilerOptions, sys System, extendedConfigCache map[string]*tsoptions.ExtendedConfigCacheEntry) (*tsoptions.ParsedCommandLine, []*ast.Diagnostic) {
	errors := []*ast.Diagnostic{}
	configFileText, errors := tsoptions.TryReadFile(configFileName, sys.FS().ReadFile, errors)
	if len(errors) > 0 {
		// these are unrecoverable errors--exit to report them as diagnotics
		return nil, errors
	}

	tsConfigSourceFile := tsoptions.NewTsconfigSourceFileFromFilePath(configFileName, configFileText)
	cwd := sys.GetCurrentDirectory()
	tsConfigSourceFile.SourceFile.SetPath(tspath.ToPath(configFileName, cwd, sys.FS().UseCaseSensitiveFileNames()))
	// tsConfigSourceFile.resolvedPath = tsConfigSourceFile.FileName()
	// tsConfigSourceFile.originalFileName = tsConfigSourceFile.FileName()
	return tsoptions.ParseJsonSourceFileConfigFileContent(
		tsConfigSourceFile,
		sys,
		tspath.GetNormalizedAbsolutePath(tspath.GetDirectoryPath(configFileName), cwd),
		options,
		tspath.GetNormalizedAbsolutePath(configFileName, cwd),
		nil,
		nil,
		extendedConfigCache,
	), nil
}

func performCompilation(sys System, cb cbType, config *tsoptions.ParsedCommandLine, reportDiagnostic diagnosticReporter) ExitStatus {
	host := compiler.NewCompilerHost(config.CompilerOptions(), sys.GetCurrentDirectory(), sys.FS())
	// todo: cache, statistics, tracing
	program := compiler.NewProgramFromParsedCommandLine(config, host)
	options := program.Options()
	allDiagnostics := program.GetOptionsDiagnostics()

	// todo: early exit logic and append diagnostics
	diagnostics := program.GetSyntacticDiagnostics(nil)
	if len(diagnostics) == 0 {
		diagnostics = append(diagnostics, program.GetOptionsDiagnostics()...)
		if options.ListFilesOnly.IsFalse() {
			// program.GetBindDiagnostics(nil)
			diagnostics = append(diagnostics, program.GetGlobalDiagnostics()...)
		}
	}
	if len(diagnostics) == 0 {
		diagnostics = append(diagnostics, program.GetSemanticDiagnostics(nil)...)
	}
	// TODO: declaration diagnostics
	if len(diagnostics) == 0 && options.NoEmit == core.TSTrue && (options.Declaration.IsTrue() && options.Composite.IsTrue()) {
		return ExitStatusNotImplemented
		// addRange(allDiagnostics, program.getDeclarationDiagnostics(/*sourceFile*/ undefined, cancellationToken));
	}

	emitResult := &compiler.EmitResult{EmitSkipped: true, Diagnostics: []*ast.Diagnostic{}}
	if options.ListFilesOnly.IsFalse() {
		// todo emit not fully implemented
		emitResult = program.Emit(&compiler.EmitOptions{})
	}
	diagnostics = append(diagnostics, emitResult.Diagnostics...)

	allDiagnostics = append(allDiagnostics, diagnostics...)
	if allDiagnostics != nil {
		allDiagnostics = compiler.SortAndDeduplicateDiagnostics(allDiagnostics)
		for _, diagnostic := range allDiagnostics {
			reportDiagnostic(diagnostic)
		}
	}

	// !!! if (write)
	if sys.Writer() != nil {
		for _, file := range emitResult.EmittedFiles {
			fmt.Fprint(sys.Writer(), "TSFILE: ", tspath.GetNormalizedAbsolutePath(file, sys.GetCurrentDirectory()))
		}
		// todo: listFiles(program, sys.Writer())
	}

	createReportErrorSummary(sys, config.CompilerOptions())(allDiagnostics)

	reportStatistics(sys, program)
	if cb != nil {
		cb(program)
	}

	if emitResult.EmitSkipped && diagnostics != nil && len(diagnostics) > 0 {
		return ExitStatusDiagnosticsPresent_OutputsSkipped
	} else if len(diagnostics) > 0 {
		return ExitStatusDiagnosticsPresent_OutputsGenerated
	}
	return ExitStatusSuccess
}

// func isBuildCommand(args []string) bool {
// 	return len(args) > 0 && args[0] == "build"
// }

func isWatchSet(options *core.CompilerOptions) bool {
	return options.Watch.IsTrue()
}

func isIncrementalCompilation(options *core.CompilerOptions) bool {
	return options.Incremental.IsTrue()
}
