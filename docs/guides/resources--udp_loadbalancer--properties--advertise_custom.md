---
page_title: "advertise_custom"
subcategory: ""
description: "advertise_custom for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1990, "body_sha256": "sha256:bc17e90623deebc356bb96577b6124bfdbc61662f54435392e411d6fe16fe069", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where"], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom", "parent_id": "xcsh-docs:resources:udp_loadbalancer:reference", "path": "docs/guides/resources--udp_loadbalancer--properties--advertise_custom.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/advertise_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- advertise_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

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

- [advertise_custom](resources--udp_loadbalancer--properties--advertise_custom.md#section)
- [advertise_on_public](resources--udp_loadbalancer--properties--advertise_on_public.md#section)
- [advertise_on_public_default_vip](resources--udp_loadbalancer--properties--advertise_on_public_default_vip.md#section)
- [do_not_advertise](resources--udp_loadbalancer--properties--do_not_advertise.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
