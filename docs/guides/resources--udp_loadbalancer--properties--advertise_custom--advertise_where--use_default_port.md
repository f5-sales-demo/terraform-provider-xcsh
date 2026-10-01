---
page_title: "advertise_custom.advertise_where.use_default_port"
subcategory: ""
description: "advertise_custom.advertise_where.use_default_port for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1222, "body_sha256": "sha256:9e9307ebf877116b3c64fc0eb28b87745084d4618d885e6395f987ea6d9c8a09", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "child_ids": [], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "parent_id": "xcsh-docs:resources:udp_loadbalancer:properties:advertise_custom:advertise_where", "path": "docs/guides/resources--udp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advertise_custom", "advertise_where", "use_default_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advertise_custom.advertise_where.use_default_port for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where.use_default_port

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Property reference](resources--udp_loadbalancer--reference.md)
- [advertise_custom](resources--udp_loadbalancer--properties--advertise_custom.md)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- advertise_custom.advertise_where.use_default_port

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
use_default_port = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advertise_custom.advertise_where](resources--udp_loadbalancer--properties--advertise_custom--advertise_where.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
