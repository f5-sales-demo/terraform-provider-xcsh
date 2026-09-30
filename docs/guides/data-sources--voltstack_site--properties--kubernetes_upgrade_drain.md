---
page_title: "kubernetes_upgrade_drain"
subcategory: ""
description: "kubernetes_upgrade_drain for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1353, "body_sha256": "sha256:0133657fbb5c06aac8a844f0a8fdc5ef8ba1085589095d493f9f02541b168c72", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:kubernetes_upgrade_drain", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:data-sources:voltstack_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--kubernetes_upgrade_drain.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kubernetes_upgrade_drain"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kubernetes_upgrade_drain for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- kubernetes_upgrade_drain

<a id="section"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

## Direct properties

- [disable_upgrade_drain](data-sources--voltstack_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md): complete subsection reference.

- [enable_upgrade_drain](data-sources--voltstack_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md): complete subsection reference.

## Next pages

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--voltstack_site--properties--kubernetes_upgrade_drain--disable_upgrade_drain.md)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--voltstack_site--properties--kubernetes_upgrade_drain--enable_upgrade_drain.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
