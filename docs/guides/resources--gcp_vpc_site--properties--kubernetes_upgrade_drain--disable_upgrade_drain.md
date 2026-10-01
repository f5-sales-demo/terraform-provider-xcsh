---
page_title: "kubernetes_upgrade_drain.disable_upgrade_drain"
subcategory: "Infrastructure"
description: "kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1108, "body_sha256": "sha256:730ed9cfa91a7b39b47d232a5cc9bd80b0b473411a5a071a20a3eae450febb9a", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "child_ids": [], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain", "path": "docs/guides/resources--gcp_vpc_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain.disable_upgrade_drain for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.disable_upgrade_drain

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--properties--kubernetes_upgrade_drain.md)
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

- [kubernetes_upgrade_drain](resources--gcp_vpc_site--properties--kubernetes_upgrade_drain.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
