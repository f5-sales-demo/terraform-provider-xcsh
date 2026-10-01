---
page_title: "api_protection_rules.api_groups_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules.client_matcher.asn_matcher for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2387, "body_sha256": "sha256:9db6fa9ece8e9162713fe7703c778a2aaf626649a2f44944c1e1bd37191e9fd0", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules.client_matcher.asn_matcher for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [api_protection_rules.api_groups_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/asn_matcher/asn_sets/)
- [api_protection_rules.api_groups_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/client_matcher/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
