---
page_title: "lb_algorithm"
subcategory: ""
description: "lb_algorithm for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1008, "body_sha256": "sha256:eef8a209470f0f81519efd0a99debc936b9bd8b128bb3af69ff388f6a71e3bca", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "docs/guides/resources--dns_proxy--properties--lb_algorithm.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["lb_algorithm"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/lb_algorithm/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "lb_algorithm for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# lb_algorithm

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- lb_algorithm

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for lb algorithm.

Upstream description:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

Terraform syntax:

```terraform
lb_algorithm {
  # Configure direct properties listed below.
}
```

## Direct properties

- [round_robin](resources--dns_proxy--properties--lb_algorithm--round_robin.md): complete subsection reference.

## Next pages

- [lb_algorithm.round_robin](resources--dns_proxy--properties--lb_algorithm--round_robin.md)
- [Property reference](resources--dns_proxy--reference.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
