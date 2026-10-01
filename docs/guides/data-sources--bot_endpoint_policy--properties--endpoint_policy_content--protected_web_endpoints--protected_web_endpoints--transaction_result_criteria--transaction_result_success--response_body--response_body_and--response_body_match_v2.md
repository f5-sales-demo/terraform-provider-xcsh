---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and.response_body_match_v2"
subcategory: ""
description: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and.response_body_match_v2 for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4724, "body_sha256": "sha256:88107dcf18a802343471ed8ea9bf44b49471d1992e419fee918a32927972e288", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_success:response_body:response_body_and:response_body_match_v2", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_success:response_body:response_body_and:response_body_match_v2", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_success:response_body:response_body_and", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and--response_body_match_v2.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "transaction_result_criteria", "transaction_result_success", "response_body", "response_body_and", "response_body_match_v2"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/response_body/response_body_and/response_body_match_v2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and.response_body_match_v2 for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and.response_body_match_v2

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
- [endpoint_policy_content.protected_web_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and.md)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and.response_body_match_v2

<a id="section"></a>

Type: `"list"`. Computed.

Response Body Matcher(s). Response Body Matchers.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and--response_body_match_v2--case_sensitive"></a>

### case_sensitive property

Type: `"bool"`. Computed.

Configuration parameter for case sensitive.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and--response_body_match_v2--not"></a>

### not property

Type: `"bool"`. Computed.

Not(!). Configuration parameter for not

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and--response_body_match_v2--operator"></a>

### operator property

Type: `"string"`. Computed.

\[Enum:
RESPONSE\_OPERATOR\_EQUALS\_TO|RESPONSE\_OPERATOR\_CONTAINS|RESPONSE\_OPERATOR\_STARTS\_WITH|RESPONSE\_OPERATOR\_ENDS\_WITH\]
&#8203;- RESPONSE\_OPERATOR\_EQUALS\_TO: EQUALS\_TO value - RESPONSE\_OPERATOR\_CONTAINS: CONTAINS value -
RESPONSE\_OPERATOR\_STARTS\_WITH: STARTS\_WITH value - RESPONSE\_OPERATOR\_ENDS\_WITH: ENDS\_WITH
value. Possible values are \`RESPONSE\_OPERATOR\_EQUALS\_TO\`, \`RESPONSE\_OPERATOR\_CONTAINS\`,
\`RESPONSE\_OPERATOR\_STARTS\_WITH\`, \`RESPONSE\_OPERATOR\_ENDS\_WITH\`. Defaults to
\`RESPONSE\_OPERATOR\_EQUALS\_TO\`.

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and--response_body_match_v2--value"></a>

### value property

Type: `"string"`. Computed.

Value. Configuration parameter for value

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_and](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_body--response_body_and.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
