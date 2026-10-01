---
page_title: "origin_servers.private_name.snat_pool.no_snat_pool"
subcategory: "Load Balancing"
description: "origin_servers.private_name.snat_pool.no_snat_pool for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1322, "body_sha256": "sha256:e22e2f41db0b424a85a2ed7aeedeca5886fb2dd8c31b73c930ec1b30f40aeecc", "canonical_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool:no_snat_pool", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool:no_snat_pool", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool", "path": "docs/guides/resources--origin_pool--properties--origin_servers--private_name--snat_pool--no_snat_pool.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "private_name", "snat_pool", "no_snat_pool"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/private_name/snat_pool/no_snat_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.private_name.snat_pool.no_snat_pool for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.private_name.snat_pool.no_snat_pool

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [origin_servers.private_name](resources--origin_pool--properties--origin_servers--private_name.md)
- [origin_servers.private_name.snat_pool](resources--origin_pool--properties--origin_servers--private_name--snat_pool.md)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_servers.private_name.snat_pool](resources--origin_pool--properties--origin_servers--private_name--snat_pool.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
