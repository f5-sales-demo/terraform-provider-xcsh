---
page_title: "rule_list.rules.asn_matcher"
subcategory: "DNS"
description: "rule_list.rules.asn_matcher for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1453, "body_sha256": "sha256:36f81e78c4110bef096c8c7f8966586c654281931e59a3dd0137ceff396ec836", "canonical_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher:asn_sets"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "path": "docs/guides/resources--dns_load_balancer--properties--rule_list--rules--asn_matcher.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "asn_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.asn_matcher for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.asn_matcher

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
- [Property reference](resources--dns_load_balancer--reference.md)
- [rule_list](resources--dns_load_balancer--properties--rule_list.md)
- [rule_list.rules](resources--dns_load_balancer--properties--rule_list--rules.md)
- rule_list.rules.asn_matcher

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

- [asn_sets](resources--dns_load_balancer--properties--rule_list--rules--asn_matcher--asn_sets.md): complete subsection reference.

## Next pages

- [rule_list.rules.asn_matcher.asn_sets](resources--dns_load_balancer--properties--rule_list--rules--asn_matcher--asn_sets.md)
- [rule_list.rules](resources--dns_load_balancer--properties--rule_list--rules.md)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
