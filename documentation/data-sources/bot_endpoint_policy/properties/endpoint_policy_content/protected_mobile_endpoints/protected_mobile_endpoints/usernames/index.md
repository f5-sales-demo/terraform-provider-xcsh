---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.usernames"
subcategory: ""
description: "Add the field name where you store username on the Request Body."
xcsh_docs: {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints usernames"], "body_bytes": 2307, "body_sha256": "sha256:8d235b4af0d948121e3409f8c0b6c2736cac15fea7ea15b5ed2e8cdc00bd5ed8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:usernames", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/usernames/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1112103210332121-2003303003301322-2112101122203321-3313303211131330-1032021221132230-0331200230120130-0212312031030120-1230310203201012", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "usernames"], "schema_version": 1, "sections": [{"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints usernames encryption type"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--usernames--encryption_type", "description": "Encryption Type for username reporting - PlainText: plaintext - Hashed: hashed. Possible values are `PlainText`, `Hashed`. Defaults to `PlainText`.", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:usernames", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "usernames", "encryption_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint policy content protected mobile endpoints protected mobile endpoints usernames username reporting"], "anchor": "schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--usernames--username_reporting", "description": "Field. Human-readable name for the resource", "document_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:usernames", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "usernames", "username_reporting"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/usernames/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Add the field name where you store username on the Request Body.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.usernames

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.usernames

<a id="section"></a>

Type: `"list"`. Computed.

Add the field name where you store username on the Request Body.

## Direct properties

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--usernames--encryption_type"></a>

### encryption_type property

Type: `"string"`. Computed.

\[Enum: PlainText|Hashed\] Encryption Type for username reporting - PlainText: plaintext - Hashed:
hashed. Possible values are \`PlainText\`, \`Hashed\`. Defaults to \`PlainText\`.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--usernames--username_reporting"></a>

### username_reporting property

Type: `"string"`. Computed.

Field. Human-readable name for the resource

## Next pages

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
