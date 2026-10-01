---
page_title: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none.domain_match"
subcategory: ""
description: "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none.domain_match for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2933, "body_sha256": "sha256:b0527bf2b20d479e9b985b2e4fa4439d8796907e5d36888831c091e262a4e848", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:domain:domain_none:domain_match", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:domain:domain_none:domain_match", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:properties:endpoint_policy_content:protected_mobile_endpoints:protected_mobile_endpoints:domain:domain_none", "path": "docs/guides/data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain--domain_none--domain_match.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint_policy_content", "protected_mobile_endpoints", "protected_mobile_endpoints", "domain", "domain_none", "domain_match"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/properties/endpoint_policy_content/protected_mobile_endpoints/protected_mobile_endpoints/domain/domain_none/domain_match/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none.domain_match for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none.domain_match

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Property reference](data-sources--bot_endpoint_policy--reference.md)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--properties--endpoint_policy_content.md)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain.md)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain--domain_none.md)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none.domain_match

<a id="section"></a>

Type: `"list"`. Computed.

Domain Matcher(s). Domain Matchers.

## Direct properties

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain--domain_none--domain_match--negation"></a>

### negation property

Type: `"string"`. Computed.

\[Enum: NO|YES\] Select from one of the Negation Operator. - NO: No - YES: Yes. Possible values are
\`NO\`, \`YES\`. Defaults to \`NO\`.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain--domain_none--domain_match--operator"></a>

### operator property

Type: `"string"`. Computed.

\[Enum: EXACT|CONTAIN|START\_WITH|END\_WITH\] Select from one of the Comparison Operator. - EXACT:
exact value - CONTAIN: contain value - START\_WITH: start with value - END\_WITH: end with value.
Possible values are \`EXACT\`, \`CONTAIN\`, \`START\_WITH\`, \`END\_WITH\`. Defaults to \`EXACT\`.

<a id="schema-endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain--domain_none--domain_match--value"></a>

### value property

Type: `"string"`. Computed.

Value. Domain Matcher Value.

## Next pages

- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.domain.domain_none](data-sources--bot_endpoint_policy--properties--endpoint_policy_content--protected_mobile_endpoints--protected_mobile_endpoints--domain--domain_none.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
