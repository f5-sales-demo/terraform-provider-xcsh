---
page_title: "cname_pool"
subcategory: ""
description: "cname_pool for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2235, "body_sha256": "sha256:dd02c1b590f9ce393673aef96a06dc6a65f99ef731ae6f67799c2cc72929903e", "child_ids": ["xcsh-docs:resources:dns_lb_pool:properties:cname_pool:disable_health_check", "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:health_check", "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members"], "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "parent_id": "xcsh-docs:resources:dns_lb_pool:reference", "path": "documentation/resources/dns_lb_pool/properties/cname_pool/index.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["cname_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/cname_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cname_pool for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
