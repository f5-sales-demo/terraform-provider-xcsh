---
page_title: "endpoint_policy_content"
subcategory: ""
description: "endpoint_policy_content for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1906, "body_sha256": "sha256:06da8a5e7c631aa4a8e315fafb38e49acc0cf50b25fe2b3d8afad29213739e3b", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints"], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:reference", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# endpoint_policy_content

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
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

- [mobile_config_fetch_paths](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--mobile_config_fetch_paths.md): complete subsection reference.

- [protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints.md): complete subsection reference.

- [protected_web_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints.md): complete subsection reference.

<a id="schema-endpoint_policy_content--telemetry_prefix"></a>

### telemetry_prefix property

Type: `"string"`. Computed.

Defines a set of headers used to detect signals based on telemetry prefix.

## Next pages

- [endpoint_policy_content.mobile_config_fetch_paths](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--mobile_config_fetch_paths.md)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints.md)
- [endpoint_policy_content.protected_web_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
