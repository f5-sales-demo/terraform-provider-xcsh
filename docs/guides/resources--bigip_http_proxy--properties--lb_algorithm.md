---
page_title: "lb_algorithm"
subcategory: ""
description: "lb_algorithm for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1163, "body_sha256": "sha256:0b169538ef80c99c4a1fb37a9746185e74f10ab8052aeec58e0a0addb6bc8f93", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm:round_robin"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "docs/guides/resources--bigip_http_proxy--properties--lb_algorithm.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["lb_algorithm"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/lb_algorithm/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "lb_algorithm for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# lb_algorithm

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
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

- [round_robin](resources--bigip_http_proxy--properties--lb_algorithm--round_robin.md): complete subsection reference.

## Next pages

- [lb_algorithm.round_robin](resources--bigip_http_proxy--properties--lb_algorithm--round_robin.md)
- [Property reference](resources--bigip_http_proxy--reference.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
