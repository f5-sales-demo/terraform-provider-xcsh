---
page_title: "endpoint_policy_content"
subcategory: ""
description: "Protected Endpoint. Configures Endpoint Policy Content."
xcsh_docs: {"aliases": ["endpoint policy content"], "body_bytes": 1636, "body_sha256": "sha256:a07604c0c99de475f0ccd983faf4d8ec7fd81fc3733cd0e631d1048e82be5795", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:reference", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content js download path"], "anchor": "schema-endpoint_policy_content--js_download_path", "description": "Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any other website/application paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "js_download_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content mobile config fetch paths"], "anchor": "section", "description": "Android and iOS mobile SDK config fetch paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content protected mobile endpoints"], "anchor": "section", "description": "Protected Mobile Endpoints. Protected Mobile Endpoints List.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content protected web endpoints"], "anchor": "section", "description": "Protected Web Endpoints. Protected Web Endpoints List.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content telemetry prefix"], "anchor": "schema-endpoint_policy_content--telemetry_prefix", "description": "Defines a set of headers used to detect signals based on telemetry prefix.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "telemetry_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Protected Endpoint. Configures Endpoint Policy Content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- endpoint_policy_content

<a id="section"></a>

Type: `"single"`. Computed.

Protected Endpoint. Configures Endpoint Policy Content.

## Direct properties

<a id="schema-endpoint_policy_content--js_download_path"></a>

### js_download_path property

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

- [mobile_config_fetch_paths](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/mobile_config_fetch_paths/): complete subsection reference.

- [protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/): complete subsection reference.

- [protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/): complete subsection reference.

<a id="schema-endpoint_policy_content--telemetry_prefix"></a>

### telemetry_prefix property

Type: `"string"`. Computed.

Defines a set of headers used to detect signals based on telemetry prefix.
