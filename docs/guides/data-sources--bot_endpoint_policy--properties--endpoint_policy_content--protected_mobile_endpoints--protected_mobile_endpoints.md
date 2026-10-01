---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints"
subcategory: ""
description: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 7683, "body_sha256": "sha256:76509d21fc17ff372a0d8b369e4d8405b9a597a47eb8b6462da7def0b3a93c87", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:allow_deny", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:block", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:continue", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:domain", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:flow_label_choice", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:header", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:metadata", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:path", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:query", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:regular_request", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:request_body", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:transaction_result_criteria", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:transform", "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:usernames"], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints.md)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints

<a id="section"></a>

Type: `"list"`. Computed.

Protected Endpoints. Endpoint or connection point

## Direct properties

- [allow_deny](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny.md): complete subsection reference.

- [block](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--block.md): complete subsection reference.

- [continue](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--continue.md): complete subsection reference.

- [domain](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain.md): complete subsection reference.

- [flow_label_choice](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--flow_label_choice.md): complete subsection reference.

- [header](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--header.md): complete subsection reference.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--http_methods"></a>

### http_methods property

Type: `["list", "string"]`. Computed.

\[Enum:
BP\_METHOD\_GET|BP\_METHOD\_POST|BP\_METHOD\_PUT|BP\_METHOD\_PATCH|BP\_METHOD\_DELETE|BP\_METHOD\_GET\_DOCUMENT|BP\_METHOD\_HEAD|BP\_METHOD\_OPTIONS|BP\_METHOD\_TRACE\]
HTTP Methods. List of HTTP methods. Possible values are \`BP\_METHOD\_GET\`, \`BP\_METHOD\_POST\`,
\`BP\_METHOD\_PUT\`, \`BP\_METHOD\_PATCH\`, \`BP\_METHOD\_DELETE\`, \`BP\_METHOD\_GET\_DOCUMENT\`,
\`BP\_METHOD\_HEAD\`, \`BP\_METHOD\_OPTIONS\`, \`BP\_METHOD\_TRACE\`. Defaults to
\`BP\_METHOD\_GET\`.

- [metadata](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--metadata.md): complete subsection reference.

- [path](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--path.md): complete subsection reference.

- [query](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--query.md): complete subsection reference.

- [regular_request](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--regular_request.md): complete subsection reference.

- [request_body](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--request_body.md): complete subsection reference.

- [transaction_result_criteria](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria.md): complete subsection reference.

- [transform](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transform.md): complete subsection reference.

- [usernames](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--usernames.md): complete subsection reference.

## Next pages

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.allow_deny](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--allow_deny.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.block](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--block.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.continue](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--continue.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.flow_label_choice](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--flow_label_choice.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.header](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--header.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.metadata](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--metadata.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.path](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--path.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.query](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--query.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.regular_request](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--regular_request.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.request_body](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--request_body.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transform](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transform.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.usernames](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--usernames.md)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
