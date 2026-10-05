-- OpenCode Go models that are served only over the OpenAI Responses API.
--
-- maki picks one wire format per models.dev provider, and models.dev marks
-- `opencode-go` as `@ai-sdk/openai-compatible`, so every model it knows about
-- is posted to /chat/completions. The gateway serves the models below on
-- /responses only and answers the rest of the catalogue there with:
--
--   API error (400): {"type":"error","error":{"type":
--   "ModelProtocolUnsupported","message":"Model does not support this
--   protocol."}}
--
-- A plugin cannot claim the `opencode-go` slug (models.dev owns it), so this
-- companion provider points the responses codec at the same gateway. Use
-- `opencode-go-responses/<model>` where you would have used
-- `opencode-go/<model>`. Everything else stays on the built-in provider.
--
-- Upstream issue: https://github.com/tontinton/maki/issues/1091
-- Endpoint table: https://opencode.ai/v2/docs/console/go/#endpoints

local BASE_URL = "https://opencode.ai/zen/go/v1"

-- id, tier, family, context_window, max_output_tokens, vision,
-- pricing = { input, output, cache_write, cache_read } in USD per 1M tokens.
-- Values mirror models.dev; the gateway charges more above a context
-- threshold, which maki's flat per-model pricing cannot express, so the
-- quoted baseline (<= threshold) rate is used.
local RESPONSES_MODELS = {
  { "grok-4.7", "strong", "generic", 500000, 131072, true, { 2.0, 6.0, 0.0, 0.5 } },
  { "grok-4.6", "strong", "generic", 500000, 131072, true, { 2.0, 6.0, 0.0, 0.5 } },
  { "gpt-6-luna", "weak", "gpt", 1050000, 128000, true, { 0.1, 0.5, 0.125, 0.01 } },
  { "gpt-5.6-luna", "weak", "gpt", 1050000, 128000, true, { 0.2, 1.2, 0.25, 0.02 } },
}

local function pricing(p)
  return { input = p[1], output = p[2], cache_write = p[3], cache_read = p[4] }
end

local model_rows = {}
local listed_rows = {}
for i, m in ipairs(RESPONSES_MODELS) do
  local id, tier, family, context, out, vision, price = m[1], m[2], m[3], m[4], m[5], m[6], m[7]
  model_rows[i] = {
    prefixes = { id },
    tier = tier,
    family = family,
    context_window = context,
    max_output_tokens = out,
    supports_vision = vision,
    pricing = pricing(price),
    default = i == 1,
  }
  listed_rows[i] = {
    id = id,
    tier = tier,
    context_window = context,
    max_output_tokens = out,
    supports_vision = vision,
    pricing = pricing(price),
  }
end

-- Without this hook maki lists the gateway's whole catalogue under this
-- slug, and the chat-completions models would be posted to /responses and
-- fail. The hook pins the picker to the responses-only models.
local function list_models()
  return listed_rows
end

-- OpenCode Go requires a stable `x-opencode-session` on every request, but
-- maki writes it only for its built-in `opencode`/`opencode-go` slugs, and the
-- responses codec ignores the `openai.session_id` option. So this provider adds
-- the header itself. The auth hook is the only hook whose headers reach a
-- responses request; maki gives it no session id and runs it once per plugin
-- load, so this is one id per process, not per conversation.
local SLUG = "opencode-go-responses"
local SESSION_HEADER = "x-opencode-session"
local SESSION_ID = string.format("maki-%d-%d", os.time(), math.random(0, 2147483647))

local function has_header(headers, name)
  for k in pairs(headers) do
    if k:lower() == name then
      return true
    end
  end
  return false
end

-- Keeps the resolved credentials (`ctx.headers` already carries the bearer,
-- whether it came from `api_key_env` or `maki auth login`) and adds the session
-- header. Omitting `base_url` keeps the resolved origin.
local function auth(ctx)
  local headers = {}
  for k, v in pairs(ctx.headers or {}) do
    headers[k] = v
  end
  if not has_header(headers, "authorization") then
    local key = maki.uv.os_getenv("OPENCODE_API_KEY")
    if key and key ~= "" then
      headers.Authorization = "Bearer " .. key
    end
  end
  headers[SESSION_HEADER] = SESSION_ID
  return { headers = headers }
end

maki.provider.register({
  slug = SLUG,
  display_name = "OpenCode Go (Responses)",
  codec = "openai-responses",
  base_url = BASE_URL,
  api_key_env = "OPENCODE_API_KEY",
  login_url = "https://opencode.ai/console",
  default_model = "grok-4.7",
  models = model_rows,
  list_models = list_models,
  auth = auth,
})
