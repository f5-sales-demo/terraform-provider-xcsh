---
page_title: "endpoint_policy_content.mobile_config_fetch_paths"
subcategory: ""
description: "endpoint_policy_content.mobile_config_fetch_paths for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1575, "body_sha256": "sha256:5df0108be3f3f499cb9c21f620ba8bccb09bd68201ca7063ba5f31004fdad247", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:mobile_config_fetch_paths", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content--mobile_config_fetch_paths.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content", "mobile_config_fetch_paths"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/mobile_config_fetch_paths/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.mobile_config_fetch_paths for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.mobile_config_fetch_paths

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
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

- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
