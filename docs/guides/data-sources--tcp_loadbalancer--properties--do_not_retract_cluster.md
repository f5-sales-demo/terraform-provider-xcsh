---
page_title: "do_not_retract_cluster"
subcategory: "Load Balancing"
description: "do_not_retract_cluster for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1231, "body_sha256": "sha256:f0b936168d6f5625cd98bd370de9d11b1558b953743346d5d128eecb3e567f85", "canonical_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:do_not_retract_cluster", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:do_not_retract_cluster", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:reference", "path": "docs/guides/data-sources--tcp_loadbalancer--properties--do_not_retract_cluster.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["do_not_retract_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/do_not_retract_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "do_not_retract_cluster for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# do_not_retract_cluster

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- do_not_retract_cluster

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

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

OneOf alternatives in this subsection:

- [do_not_retract_cluster](data-sources--tcp_loadbalancer--properties--do_not_retract_cluster.md#section)
- [retract_cluster](data-sources--tcp_loadbalancer--properties--retract_cluster.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md)
