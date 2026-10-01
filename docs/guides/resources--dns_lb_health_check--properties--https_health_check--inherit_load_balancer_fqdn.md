---
page_title: "https_health_check.inherit_load_balancer_fqdn"
subcategory: ""
description: "https_health_check.inherit_load_balancer_fqdn for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": [], "body_bytes": 1141, "body_sha256": "sha256:fac253a090de4bcd89d82daffcb510256d5625f0c425752dae3e6ced634cebd9", "canonical_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:inherit_load_balancer_fqdn", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check:inherit_load_balancer_fqdn", "parent_id": "xcsh-docs:resources:dns_lb_health_check:properties:https_health_check", "path": "docs/guides/resources--dns_lb_health_check--properties--https_health_check--inherit_load_balancer_fqdn.md", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_health_check", "inherit_load_balancer_fqdn"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/https_health_check/inherit_load_balancer_fqdn/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_health_check.inherit_load_balancer_fqdn for xcsh_dns_lb_health_check.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_health_check.inherit_load_balancer_fqdn

Breadcrumbs:

- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md)
- [Property reference](resources--dns_lb_health_check--reference.md)
- [https_health_check](resources--dns_lb_health_check--properties--https_health_check.md)
- https_health_check.inherit_load_balancer_fqdn

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit load balancer fqdn.

Upstream description:

This can be used for messages where no values are needed.

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
inherit_load_balancer_fqdn = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [https_health_check](resources--dns_lb_health_check--properties--https_health_check.md)
- [xcsh_dns_lb_health_check](../resources/dns_lb_health_check.md)
