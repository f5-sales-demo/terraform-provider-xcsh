---
page_title: "endpoint_policy_content.mobile_config_fetch_paths"
subcategory: ""
description: "Android and iOS mobile SDK config fetch paths."
xcsh_docs: {"aliases": ["endpoint policy content mobile config fetch paths"], "body_bytes": 1832, "body_sha256": "sha256:3b5c4d8d5772a30bfbd71e53318518e84df59a7f033fd5446c8409f82b578c59", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/mobile_config_fetch_paths/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2332331230121203-3130100112333110-3231022010203120-0110033322132112-2312300200230211-0123020233213122-3101122011333212-3102322032033300", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths"], "schema_version": 1, "sections": [{"aliases": ["path android"], "anchor": "schema-endpoint_policy_content--mobile_config_fetch_paths--path_android", "description": "Android mobile client will fetch F5 Client mobile configuration SDK from this path. This path must not conflict with any other website/mobile/application paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths", "path_android"], "syntax": "attribute", "type": "string"}, {"aliases": ["path ios"], "anchor": "schema-endpoint_policy_content--mobile_config_fetch_paths--path_ios", "description": "IOS mobile client will fetch F5 Client mobile configuration SDK from this path. This path must not conflict with any other website/mobile/application paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths", "path_ios"], "syntax": "attribute", "type": "string"}, {"aliases": ["unavailable text"], "anchor": "schema-endpoint_policy_content--mobile_config_fetch_paths--unavailable_text", "description": "Certain mobile policies rely on older mobile components.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths", "unavailable_text"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/mobile_config_fetch_paths/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Android and iOS mobile SDK config fetch paths.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
