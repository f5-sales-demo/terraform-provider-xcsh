---
page_title: "endpoint_policy_content"
subcategory: ""
description: "Protected Endpoint. Configures Endpoint Policy Content."
xcsh_docs: {"aliases": ["endpoint policy content"], "body_bytes": 1636, "body_sha256": "sha256:a07604c0c99de475f0ccd983faf4d8ec7fd81fc3733cd0e631d1048e82be5795", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:reference", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content js download path"], "anchor": "schema-endpoint_policy_content--js_download_path", "description": "Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any other website/application paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "js_download_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content mobile config fetch paths"], "anchor": "section", "description": "Android and iOS mobile SDK config fetch paths.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content protected mobile endpoints"], "anchor": "section", "description": "Protected Mobile Endpoints. Protected Mobile Endpoints List.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content protected web endpoints"], "anchor": "section", "description": "Protected Web Endpoints. Protected Web Endpoints List.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint_policy_content", "protected_web_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint policy content telemetry prefix"], "anchor": "schema-endpoint_policy_content--telemetry_prefix", "description": "Defines a set of headers used to detect signals based on telemetry prefix.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "telemetry_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Protected Endpoint. Configures Endpoint Policy Content.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
