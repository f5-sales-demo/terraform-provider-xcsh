---
page_title: "cname_pool"
subcategory: ""
description: "Pool for CNAME Record."
xcsh_docs: {"aliases": ["cname pool"], "body_bytes": 2235, "body_sha256": "sha256:dd02c1b590f9ce393673aef96a06dc6a65f99ef731ae6f67799c2cc72929903e", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:health_check", "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "parent_id": "xcsh-docs:resources:dns_lb_pool:reference", "path": "documentation/resources/dns_lb_pool/properties/cname_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2102130031203200-0130012310030313-1203320121212201-1330323012031100-2012202003023332-0303332132013312-0103012331231303-1323123021230220", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cname_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cname_pool:ConflictingObjectAttributes:disable_health_check,health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cname_pool:RequiredObjectAttributes:members", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cname_pool"], "schema_version": 1, "sections": [{"aliases": ["disable health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "disable_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["health check"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cname_pool--health_check--name", "enforcement": "provider-schema", "group": "cname_pool.health_check:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:health_check", "type": "requires"}], "schema_path": ["cname_pool", "health_check"], "syntax": "block", "type": "object"}, {"aliases": ["members"], "anchor": "section", "description": "Configuration parameter for members", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cname_pool--members--domain", "enforcement": "provider-schema", "group": "cname_pool.members:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "type": "requires"}], "schema_path": ["cname_pool", "members"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/cname_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Pool for CNAME Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- cname_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Pool for CNAME Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("members"),
  validators.ConflictingObjectAttributes("disable_health_check",
    "health_check")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

Terraform syntax:

```terraform
cname_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/disable_health_check/): complete subsection reference.

- [health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/health_check/): complete subsection reference.

- [members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/): complete subsection reference.

## Next pages

- [cname_pool.disable_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/disable_health_check/)
- [cname_pool.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/health_check/)
- [cname_pool.members](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/members/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
