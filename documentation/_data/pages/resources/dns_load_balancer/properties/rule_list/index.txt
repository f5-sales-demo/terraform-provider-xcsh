---
page_title: "rule_list"
subcategory: "DNS"
description: "List of the Load Balancing Rules."
xcsh_docs: {"aliases": ["rule list"], "body_bytes": 1538, "body_sha256": "sha256:492ce1bcf897b77ee4e1d618629a2b025e94a823a9f1bbdb5419dcfc05968547", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list", "parent_id": "xcsh-docs:resources:dns_load_balancer:reference", "path": "documentation/resources/dns_load_balancer/properties/rule_list/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list"], "schema_version": 1, "sections": [{"aliases": ["rules"], "anchor": "section", "description": "Rules to perform load balancing.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of the Load Balancing Rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- rule_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Load Balancing Rule List. List of the Load Balancing Rules.

Upstream description:

List of the Load Balancing Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
rule_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/): complete subsection reference.

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
