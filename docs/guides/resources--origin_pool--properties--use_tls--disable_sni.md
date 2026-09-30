---
page_title: "use_tls.disable_sni"
subcategory: "Load Balancing"
description: "use_tls.disable_sni for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 860, "body_sha256": "sha256:079fffc1f881781de0f4131d94354ce1f46e33f1e4737efcaa139d0af28df141", "canonical_id": "xcsh-docs:resources:origin_pool:properties:use_tls:disable_sni", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:disable_sni", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "path": "docs/guides/resources--origin_pool--properties--use_tls--disable_sni.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "disable_sni"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/disable_sni/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.disable_sni for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# use_tls.disable_sni

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- use_tls.disable_sni

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [use_tls](resources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
