---
page_title: "lb_algorithm.round_robin"
subcategory: ""
description: "lb_algorithm.round_robin for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 889, "body_sha256": "sha256:70b776b95e4b2d1bcf36d91ed2971534440ca54e4b0c507a1035b98f1aed6141", "canonical_id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin", "parent_id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "path": "docs/guides/resources--dns_proxy--properties--lb_algorithm--round_robin.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["lb_algorithm", "round_robin"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/lb_algorithm/round_robin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "lb_algorithm.round_robin for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# lb_algorithm.round_robin

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- [Property reference](resources--dns_proxy--reference.md)
- [lb_algorithm](resources--dns_proxy--properties--lb_algorithm.md)
- lb_algorithm.round_robin

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

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
round_robin {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [lb_algorithm](resources--dns_proxy--properties--lb_algorithm.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
