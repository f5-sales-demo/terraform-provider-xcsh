---
page_title: "use_default_psp"
subcategory: ""
description: "use_default_psp for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 873, "body_sha256": "sha256:76bd13c46d43aa81ab8e5583a6cce71e9352dea3869e86ba4fd25ca9131eacb6", "canonical_id": "xcsh-docs:resources:k8s_cluster:properties:use_default_psp", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:use_default_psp", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "docs/guides/resources--k8s_cluster--properties--use_default_psp.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_default_psp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/use_default_psp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_default_psp for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# use_default_psp

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Property reference](resources--k8s_cluster--reference.md)
- use_default_psp

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for use default psp. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
use_default_psp = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--k8s_cluster--reference.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
