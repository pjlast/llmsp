package types

import "github.com/sourcegraph/go-lsp"

type MemoryFileMap map[lsp.DocumentURI]string

type LLMSPSettings struct {
	Sourcegraph *SourcegraphSettings `json:"sourcegraph"`
}

type SourcegraphSettings struct {
	URL            string   `json:"url"`
	AccessToken    string   `json:"accessToken"`
	RepoEmbeddings []string `json:"repos"`
}

type LLMSPConfig struct {
	Settings SourcegraphSettings `json:"sourcegraph"`
}

type ConfigurationSettings struct {
	LLMSP LLMSPSettings `json:"llmsp"`
}

type TextDocumentEdit struct {
	TextDocument lsp.VersionedTextDocumentIdentifier `json:"textDocument"`
	Edits        []lsp.TextEdit                      `json:"edits"`
}

type WorkspaceEdit struct {
	DocumentChanges []TextDocumentEdit `json:"documentChanges"`
}

type ApplyWorkspaceEditParams struct {
	Edit WorkspaceEdit `json:"edit"`
}

type DidChangeConfigurationParams struct {
	Settings ConfigurationSettings `json:"settings"`
}

type CompletionList struct {
	IsIncomplete bool                 `json:"isIncomplete"`
	Items        []lsp.CompletionItem `json:"items"`
}

type CompletionOptions struct {
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
	WorkDoneProgress  bool     `json:"workDoneProgress,omitempty"`
}

type CompletionParams struct {
	lsp.TextDocumentPositionParams
	Context            lsp.CompletionContext `json:"context,omitempty"`
	PartialResultToken int                `json:"partialResultToken,omitempty"`
	WorkDoneToken      int                `json:"workDoneToken,omitempty"`
}

type ProgressParams[T any] struct {
	Token string `json:"token"`
	Value T      `json:"value"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
}

type ServerCapabilities struct {
	TextDocumentSync                 *lsp.TextDocumentSyncOptionsOrKind   `json:"textDocumentSync,omitempty"`
	HoverProvider                    bool                                 `json:"hoverProvider,omitempty"`
	CompletionProvider               *CompletionOptions                   `json:"completionProvider,omitempty"`
	SignatureHelpProvider            *lsp.SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
	DefinitionProvider               bool                                 `json:"definitionProvider,omitempty"`
	TypeDefinitionProvider           bool                                 `json:"typeDefinitionProvider,omitempty"`
	ReferencesProvider               bool                                 `json:"referencesProvider,omitempty"`
	DocumentHighlightProvider        bool                                 `json:"documentHighlightProvider,omitempty"`
	DocumentSymbolProvider           bool                                 `json:"documentSymbolProvider,omitempty"`
	WorkspaceSymbolProvider          bool                                 `json:"workspaceSymbolProvider,omitempty"`
	ImplementationProvider           bool                                 `json:"implementationProvider,omitempty"`
	CodeActionProvider               bool                                 `json:"codeActionProvider,omitempty"`
	CodeLensProvider                 *lsp.CodeLensOptions                 `json:"codeLensProvider,omitempty"`
	DocumentFormattingProvider       bool                                 `json:"documentFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider  bool                                 `json:"documentRangeFormattingProvider,omitempty"`
	DocumentOnTypeFormattingProvider *lsp.DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
	RenameProvider                   bool                                 `json:"renameProvider,omitempty"`
	ExecuteCommandProvider           *lsp.ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
	SemanticHighlighting             *lsp.SemanticHighlightingOptions     `json:"semanticHighlighting,omitempty"`

	// XWorkspaceReferencesProvider indicates the server provides support for
	// xworkspace/references. This is a Sourcegraph extension.
	XWorkspaceReferencesProvider bool `json:"xworkspaceReferencesProvider,omitempty"`

	// XDefinitionProvider indicates the server provides support for
	// textDocument/xdefinition. This is a Sourcegraph extension.
	XDefinitionProvider bool `json:"xdefinitionProvider,omitempty"`

	// XWorkspaceSymbolByProperties indicates the server provides support for
	// querying symbols by properties with WorkspaceSymbolParams.symbol. This
	// is a Sourcegraph extension.
	XWorkspaceSymbolByProperties bool `json:"xworkspaceSymbolByProperties,omitempty"`

	Experimental interface{} `json:"experimental,omitempty"`
}

type CompletionItem struct {
  lsp.CompletionItem
  Preselect bool `json:"preselect,omitempty"`
}
