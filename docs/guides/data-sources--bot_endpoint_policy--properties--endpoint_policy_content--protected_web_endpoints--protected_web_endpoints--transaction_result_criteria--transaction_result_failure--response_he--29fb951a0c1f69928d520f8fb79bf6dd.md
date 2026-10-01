---
page_title: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header.header_and"
subcategory: ""
description: "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header.header_and for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 4700, "body_sha256": "sha256:3a48b3d01d3dcf9ca2b68d113c97e8a9036b79d994c75f7e39d90cf1262e2660", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_failure:response_header_v2:response_header_and:response_header_operator:header:header_and", "child_ids": ["xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_failure:response_header_v2:response_header_and:response_header_operator:header:header_and:response_header_match_v2"], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_failure:response_header_v2:response_header_and:response_header_operator:header:header_and", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_web_endpoints:protected_web_endpoints:transaction_result_criteria:transaction_result_failure:response_header_v2:response_header_and:response_header_operator:header", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_he--29fb951a0c1f69928d520f8fb79bf6dd.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_web_endpoints", "protected_web_endpoints", "transaction_result_criteria", "transaction_result_failure", "response_header_v2", "response_header_and", "response_header_operator", "header", "header_and"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_web_endpoints/protected_web_endpoints/transaction_result_criteria/transaction_result_failure/response_header_v2/response_header_and/response_header_operator/header/header_and/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header.header_and for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header.header_and

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
- [endpoint_policy_content.protected_web_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_header_v2.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_header_v2--response_header_and.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_he--8f61b2f13de9897b5a89b9aa523a766a.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_he--760e5899ccb7d2e7356e70daafd76393.md)
- endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header.header_and

<a id="section"></a>

Type: `"single"`. Computed.

Header Matcher. Response Header matcher Choice.

## Direct properties

- [response_header_match_v2](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_he--6f638f24e03c99e10085f6f8f586f05a.md): complete subsection reference.

## Next pages

- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_he--6f638f24e03c99e10085f6f8f586f05a.md)
- [endpoint_policy_content.protected_web_endpoints.protected_web_endpoints.transaction_result_criteria.transaction_result_failure.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_web_endpoints--protected_web_endpoints--transaction_result_criteria--transaction_result_failure--response_he--760e5899ccb7d2e7356e70daafd76393.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
