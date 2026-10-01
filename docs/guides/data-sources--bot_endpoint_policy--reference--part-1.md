---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1632, "body_sha256": "sha256:a22a0c50738d9539be77122fa8a8f9cc9815c9d44baa64388a13c584b0a572f0", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:reference", "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:cookies", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content"], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:reference", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:fundamentals", "path": "docs/guides/data-sources--bot_endpoint_policy--reference--part-1.md", "projection_part": 1, "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations.

- [cookies](data-sources--bot_endpoint_policy--properties--cookies.md): complete subsection reference.

<a id="schema-deployment_mode"></a>

### deployment_mode property

Type: `"string"`. Computed.

\[Enum: REVERSE\_PROXY|API\_MODE\] Deployment Mode By default, the mode will be Reverse Proxy You
need to submit an XC support ticket to request for API mode. Possible values are \`REVERSE\_PROXY\`,
\`API\_MODE\`. Defaults to \`REVERSE\_PROXY\`.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description.

- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels.

<a id="schema-latest_version"></a>

### latest_version property

Type: `"string"`. Computed.

The version number to Endpoint Policy Version for the latest version.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the BotEndpointPolicy to look up.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace of the BotEndpointPolicy.
