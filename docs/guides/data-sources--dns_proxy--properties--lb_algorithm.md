---
page_title: "lb_algorithm"
subcategory: ""
description: "lb_algorithm for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 905, "body_sha256": "sha256:6d3cc4d941a3caaf020ac6b5f642d4c3a4f5e4bbbb839e92831afb734b6b79bd", "canonical_id": "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm:round_robin"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "docs/guides/data-sources--dns_proxy--properties--lb_algorithm.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["lb_algorithm"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/lb_algorithm/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "lb_algorithm for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# lb_algorithm

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- lb_algorithm

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [round_robin](data-sources--dns_proxy--properties--lb_algorithm--round_robin.md): complete subsection reference.

## Next pages

- [lb_algorithm.round_robin](data-sources--dns_proxy--properties--lb_algorithm--round_robin.md)
- [Property reference](data-sources--dns_proxy--reference.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
