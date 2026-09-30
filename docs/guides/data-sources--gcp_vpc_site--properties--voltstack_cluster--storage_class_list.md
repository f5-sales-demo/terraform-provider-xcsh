---
page_title: "voltstack_cluster.storage_class_list"
subcategory: "Infrastructure"
description: "voltstack_cluster.storage_class_list for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1076, "body_sha256": "sha256:a1cfb0a72471bea7259912f27900341f298f6f378c038bae399e13b5402b900b", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list:storage_classes"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "storage_class_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.storage_class_list for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.storage_class_list

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.storage_class_list

<a id="section"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this site.

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

## Direct properties

- [storage_classes](data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md): complete subsection reference.

## Next pages

- [voltstack_cluster.storage_class_list.storage_classes](data-sources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list--storage_classes.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
