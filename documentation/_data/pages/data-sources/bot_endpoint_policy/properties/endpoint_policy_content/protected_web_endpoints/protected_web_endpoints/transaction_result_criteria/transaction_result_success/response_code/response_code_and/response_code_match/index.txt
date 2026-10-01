---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match"
subcategory: ""
description: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4738, "body_sha256": "sha256:c28a8379fe4d08e58123a33b60f0ad4d71ce39c9ee3a9cafdae1a0a3403e63a3", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_success:response_code:response_code_and:response_code_match", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_success:response_code:response_code_and", "path": "documentation/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/response_code/response_code_and/response_code_match/index.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "transaction_result_criteria", "transaction_result_success", "response_code", "response_code_and", "response_code_match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/response_code/response_code_and/response_code_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/)
- [endpoint_policy_content](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/)
- [endpoint_policy_content.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/response_code/)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/response_code/response_code_and/)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match

<a id="section"></a>

Type: `"list"`. Computed.

Response Code Matcher(s). Response Code Matchers.

## Direct properties

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_and--response_code_match--operator"></a>

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

<a id="schema-endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_success--response_code--response_code_and--response_code_match--value"></a>

### value property

Type: `"number"`. Computed.

Value. Configuration parameter for value

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_success/response_code/response_code_and/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
