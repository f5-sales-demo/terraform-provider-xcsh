---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match"
subcategory: ""
description: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4296, "body_sha256": "sha256:282c5b1a63e4b69620d91c57f61b8d8be3b236b219cd309a2559c1510c0d2ab8", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:transaction_result_criteria:transaction_result_success:response_code:response_code_or:response_code_match", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:transaction_result_criteria:transaction_result_success:response_code:response_code_or:response_code_match", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:transaction_result_criteria:transaction_result_success:response_code:response_code_or", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_or--response_code_match.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "transaction_result_criteria", "transaction_result_success", "response_code", "response_code_or", "response_code_match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/transaction_result_criteria/transaction_result_success/response_code/response_code_or/response_code_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success--response_code.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_or.md)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match

<a id="section"></a>

Type: `"list"`. Computed.

Response Code Matcher(s). Response Code Matchers.

## Direct properties

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_or--response_code_match--operator"></a>

### operator property

Type: `"string"`. Computed.

\[Enum:
CODE\_EQUALS|CODE\_NOT\_EQUAL\_TO|CODE\_LESS\_THAN|CODE\_GREATER\_THAN|CODE\_LESS\_THAN\_OR\_EQUAlS\_TO|CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\]
&#8203;- CODE\_EQUALS: EQUALS value - CODE\_NOT\_EQUAL\_TO: NOT\_EQUAL\_TO value - CODE\_LESS\_THAN:
LESS\_THAN value - CODE\_GREATER\_THAN: GREATER\_THAN value - CODE\_LESS\_THAN\_OR\_EQUAlS\_TO:
LESS\_THAN\_OR\_EQUAlS\_TO value - CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO:
GREATER\_THAN\_OR\_EQUAlS\_TO value. Possible values are \`CODE\_EQUALS\`, \`CODE\_NOT\_EQUAL\_TO\`,
\`CODE\_LESS\_THAN\`, \`CODE\_GREATER\_THAN\`, \`CODE\_LESS\_THAN\_OR\_EQUAlS\_TO\`,
\`CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\`. Defaults to \`CODE\_EQUALS\`.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_or--response_code_match--value"></a>

### value property

Type: `"number"`. Computed.

Value. Configuration parameter for value

## Next pages

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_or.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
