---
page_title: "advertise_custom"
subcategory: "Load Balancing"
description: "This defines a way to advertise a VIP on specific sites."
xcsh_docs: {"aliases": ["advertise custom"], "body_bytes": 2899, "body_sha256": "sha256:496b285cc8893804d44dfc1ac5af512dd06482c489d80e1a81f64bf40ef5737a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/advertise_custom/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom:RequiredObjectAttributes:advertise_where", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom"], "schema_version": 1, "sections": [{"aliases": ["advertise where"], "anchor": "section", "description": "Where should this load balancer be available.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:advertise_custom:advertise_where", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["advertise_custom", "advertise_where"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/advertise_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a way to advertise a VIP on specific sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- advertise_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: advertise\_custom, advertise\_dualstack\_on\_public, advertise\_on\_public,
advertise\_on\_public\_default\_vip, advertise\_v6\_on\_public, do\_not\_advertise; Default:
advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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

OneOf alternatives in this subsection:

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_custom/#section)
- [advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_dualstack_on_public/#section)
- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_on_public/#section)
- [advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_on_public_default_vip/#section)
- [advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_v6_on_public/#section)
- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/do_not_advertise/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_custom/advertise_where/): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/advertise_custom/advertise_where/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
