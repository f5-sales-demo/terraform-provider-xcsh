---
page_title: "rule_list.rules.ip_prefix_set"
subcategory: "DNS"
description: "rule_list.rules.ip_prefix_set for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 2091, "body_sha256": "sha256:45929f7650686ae8771dd05215d9ef94a2e4b81940abec11592302e89957e146", "canonical_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set:prefix_sets"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "path": "docs/guides/resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_set.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.ip_prefix_set for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.ip_prefix_set

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
- [Property reference](resources--dns_load_balancer--reference.md)
- [rule_list](resources--dns_load_balancer--properties--rule_list.md)
- [rule_list.rules](resources--dns_load_balancer--properties--rule_list--rules.md)
- rule_list.rules.ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--ip_prefix_set--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_set--prefix_sets.md): complete subsection reference.

## Next pages

- [rule_list.rules.ip_prefix_set.prefix_sets](resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_set--prefix_sets.md)
- [rule_list.rules](resources--dns_load_balancer--properties--rule_list--rules.md)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
