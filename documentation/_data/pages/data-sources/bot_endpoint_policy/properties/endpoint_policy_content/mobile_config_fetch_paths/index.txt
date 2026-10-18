---
page_title: "endpoint_policy_content.mobile_config_fetch_paths"
subcategory: ""
description: "Android and iOS mobile SDK config fetch paths."
xcsh_docs: {"aliases": ["endpoint policy content mobile config fetch paths"], "body_bytes": 1542, "body_sha256": "sha256:560b6c296aaebbbe429fb44efbffbc27a6458fd5185ce7d57689e10ba015838b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/mobile_config_fetch_paths/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2332331230121203-3130100112333110-3231022010203120-0110033322132112-2312300200230211-0123020233213122-3101122011333212-3102322032033300", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content mobile config fetch paths path android"], "anchor": "schema-endpoint_policy_content--mobile_config_fetch_paths--path_android", "description": "Android mobile client will fetch F5 Client mobile configuration SDK from this path. This path must not conflict with any other website/mobile/application paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths", "path_android"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content mobile config fetch paths path ios"], "anchor": "schema-endpoint_policy_content--mobile_config_fetch_paths--path_ios", "description": "IOS mobile client will fetch F5 Client mobile configuration SDK from this path. This path must not conflict with any other website/mobile/application paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths", "path_ios"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content mobile config fetch paths unavailable text"], "anchor": "schema-endpoint_policy_content--mobile_config_fetch_paths--unavailable_text", "description": "Certain mobile policies rely on older mobile components.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths", "unavailable_text"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/mobile_config_fetch_paths/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Android and iOS mobile SDK config fetch paths.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": [], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.mobile_config_fetch_paths

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- endpoint_policy_content.mobile_config_fetch_paths

<a id="section"></a>

Type: `"single"`. Computed.

Android and iOS mobile SDK config fetch paths.

## Direct properties

<a id="schema-endpoint_policy_content--mobile_config_fetch_paths--path_android"></a>

### path_android property

Type: `"string"`. Computed.

Android mobile client will fetch F5 Client mobile configuration SDK from this path. This path must
not conflict with any other website/mobile/application paths.

<a id="schema-endpoint_policy_content--mobile_config_fetch_paths--path_ios"></a>

### path_ios property

Type: `"string"`. Computed.

IOS mobile client will fetch F5 Client mobile configuration SDK from this path. This path must not
conflict with any other website/mobile/application paths.

<a id="schema-endpoint_policy_content--mobile_config_fetch_paths--unavailable_text"></a>

### unavailable_text property

Type: `"string"`. Computed.

Certain mobile policies rely on older mobile components.
