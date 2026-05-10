package app

// freeModeEntry maps a model to its required agentId for codebuff free mode.
type freeModeEntry struct {
	AgentID string
	Model   string
}

// freeModeModels is the allowlist of (agentId, model) pairs that codebuff
// permits in cost_mode=free. Derived from codebuff source:
// common/src/constants/free-agents.ts → FREE_MODE_AGENT_MODELS
// common/src/constants/freebuff-models.ts → SUPPORTED_FREEBUFF_MODELS
//
// Last synced: 2026-05-11 (codebuff v1.0.674)
var freeModeModels = map[string]freeModeEntry{
	// Root orchestrators
	"minimax/minimax-m2.7":      {AgentID: "base2-free", Model: "minimax/minimax-m2.7"},
	"z-ai/glm-5.1":             {AgentID: "base2-free", Model: "z-ai/glm-5.1"},
	"deepseek/deepseek-v4-pro":  {AgentID: "base2-free-deepseek", Model: "deepseek/deepseek-v4-pro"},
	"moonshotai/kimi-k2.6":     {AgentID: "base2-free-kimi", Model: "moonshotai/kimi-k2.6"},

	// Code reviewers
	"code-reviewer-minimax/minimax-m2.7":     {AgentID: "code-reviewer-minimax", Model: "minimax/minimax-m2.7"},
	"code-reviewer-minimax/glm-5.1":          {AgentID: "code-reviewer-minimax", Model: "z-ai/glm-5.1"},
	"code-reviewer-kimi/kimi-k2.6":           {AgentID: "code-reviewer-kimi", Model: "moonshotai/kimi-k2.6"},
	"code-reviewer-deepseek/deepseek-v4-pro": {AgentID: "code-reviewer-deepseek", Model: "deepseek/deepseek-v4-pro"},
	// Legacy code-reviewer-lite (kept for older clients)
	"code-reviewer-lite/minimax-m2.7":     {AgentID: "code-reviewer-lite", Model: "minimax/minimax-m2.7"},
	"code-reviewer-lite/kimi-k2.6":        {AgentID: "code-reviewer-lite", Model: "moonshotai/kimi-k2.6"},
	"code-reviewer-lite/deepseek-v4-pro":  {AgentID: "code-reviewer-lite", Model: "deepseek/deepseek-v4-pro"},

	// File exploration agents
	"google/gemini-2.5-flash-lite":         {AgentID: "file-picker", Model: "google/gemini-2.5-flash-lite"},
	"google/gemini-3.1-flash-lite-preview": {AgentID: "file-picker-max", Model: "google/gemini-3.1-flash-lite-preview"},

	// Research agents
	"researcher-web/gemini-3.1-flash-lite-preview":  {AgentID: "researcher-web", Model: "google/gemini-3.1-flash-lite-preview"},
	"researcher-docs/gemini-3.1-flash-lite-preview": {AgentID: "researcher-docs", Model: "google/gemini-3.1-flash-lite-preview"},

	// Browser automation
	"browser-use/gemini-3.1-flash-lite-preview": {AgentID: "browser-use", Model: "google/gemini-3.1-flash-lite-preview"},

	// Command execution
	"basher/gemini-3.1-flash-lite-preview": {AgentID: "basher", Model: "google/gemini-3.1-flash-lite-preview"},
}

// freeModeDefaultModel is the fallback model when no model is specified in free mode.
const freeModeDefaultModel = "minimax/minimax-m2.7"

// resolveFreeModeAgent looks up the model in the free-mode allowlist.
// Returns (agentId, canonicalModel, true) if allowed, or ("", "", false)
// if the model is not available in free mode.
func resolveFreeModeAgent(model string) (agentID, canonicalModel string, ok bool) {
	if entry, found := freeModeModels[model]; found {
		return entry.AgentID, entry.Model, true
	}
	return "", "", false
}
