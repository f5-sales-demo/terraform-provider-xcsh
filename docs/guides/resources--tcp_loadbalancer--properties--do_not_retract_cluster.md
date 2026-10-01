---
page_title: "do_not_retract_cluster"
subcategory: "Load Balancing"
description: "do_not_retract_cluster for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1278, "body_sha256": "sha256:f2c2aaf60fdeb4de6505fc4aedcedc3e7cd275c582fa6c67a521abaccd0ddf3f", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:properties:do_not_retract_cluster", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:do_not_retract_cluster", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "path": "docs/guides/resources--tcp_loadbalancer--properties--do_not_retract_cluster.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["do_not_retract_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/do_not_retract_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "do_not_retract_cluster for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# do_not_retract_cluster

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Property reference](resources--tcp_loadbalancer--reference.md)
- do_not_retract_cluster

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

- [do_not_retract_cluster](resources--tcp_loadbalancer--properties--do_not_retract_cluster.md#section)
- [retract_cluster](resources--tcp_loadbalancer--properties--retract_cluster.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_retract_cluster = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--tcp_loadbalancer--reference.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
