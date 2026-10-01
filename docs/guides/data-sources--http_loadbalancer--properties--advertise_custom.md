---
page_title: "advertise_custom"
subcategory: "Load Balancing"
description: "advertise_custom for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2042, "body_sha256": "sha256:8783d4638166db3b7a2a636e38037475a5540337c7f0c365c2e0183c1aee9dbf", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom:advertise_where"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:advertise_custom", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--advertise_custom.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/advertise_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- advertise_custom

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: advertise\_custom, advertise\_dualstack\_on\_public, advertise\_on\_public,
advertise\_on\_public\_default\_vip, advertise\_v6\_on\_public, do\_not\_advertise; Default:
advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

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

- [advertise_custom](data-sources--http_loadbalancer--properties--advertise_custom.md#section)
- [advertise_dualstack_on_public](data-sources--http_loadbalancer--properties--advertise_dualstack_on_public.md#section)
- [advertise_on_public](data-sources--http_loadbalancer--properties--advertise_on_public.md#section)
- [advertise_on_public_default_vip](data-sources--http_loadbalancer--properties--advertise_on_public_default_vip.md#section)
- [advertise_v6_on_public](data-sources--http_loadbalancer--properties--advertise_v6_on_public.md#section)
- [do_not_advertise](data-sources--http_loadbalancer--properties--do_not_advertise.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [advertise_where](data-sources--http_loadbalancer--properties--advertise_custom--advertise_where.md): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where](data-sources--http_loadbalancer--properties--advertise_custom--advertise_where.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
