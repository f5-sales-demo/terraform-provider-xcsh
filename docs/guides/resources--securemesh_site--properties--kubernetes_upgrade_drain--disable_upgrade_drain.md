---
page_title: "kubernetes_upgrade_drain.disable_upgrade_drain"
subcategory: ""
description: "kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1129, "body_sha256": "sha256:58c2bca4096a71660b886f075936768e38fd1e77db1148f598bbf97d4dc136cd", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "parent_id": "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain", "path": "docs/guides/resources--securemesh_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.disable_upgrade_drain

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [kubernetes_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain.md)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kubernetes_upgrade_drain](resources--securemesh_site--properties--kubernetes_upgrade_drain.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
