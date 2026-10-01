---
page_title: "api_protection_rules.api_groups_rules.client_matcher"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules.client_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5805, "body_sha256": "sha256:c23f7ded2582ccfa0e1550c86b7e8aa98da94eda8c0de2707b32798622943c33", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_client", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:any_ip", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:client_selector", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_prefix_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:ip_threat_category_list", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:tls_fingerprint_matcher"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules.client_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.client_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- api_protection_rules.api_groups_rules.client_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

## Direct properties

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/any_client/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/any_ip/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/): complete subsection reference.

- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/client_selector/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_prefix_list/): complete subsection reference.

- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_threat_category_list/): complete subsection reference.

- [tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/tls_fingerprint_matcher/): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.client_matcher.any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/any_client/)
- [api_protection_rules.api_groups_rules.client_matcher.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/any_ip/)
- [api_protection_rules.api_groups_rules.client_matcher.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_list/)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/)
- [api_protection_rules.api_groups_rules.client_matcher.client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/client_selector/)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_matcher/)
- [api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_prefix_list/)
- [api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/ip_threat_category_list/)
- [api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/tls_fingerprint_matcher/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
