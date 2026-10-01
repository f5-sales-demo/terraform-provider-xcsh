---
page_title: "gcp"
subcategory: ""
description: "gcp for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1146, "body_sha256": "sha256:059860c478f1899e62aece16a40912a2294631750fa2c9afd2608e60cc02d538", "canonical_id": "xcsh-docs:data-sources:cloud_link:properties:gcp", "child_ids": ["xcsh-docs:data-sources:cloud_link:properties:gcp:byoc", "xcsh-docs:data-sources:cloud_link:properties:gcp:gcp_cred"], "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:gcp", "parent_id": "xcsh-docs:data-sources:cloud_link:reference", "path": "docs/guides/data-sources--cloud_link--properties--gcp.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/gcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md)
- [Property reference](data-sources--cloud_link--reference.md)
- gcp

<a id="section"></a>

Type: `"single"`. Computed.

Google Cloud Platform (GCP) CloudLink Provider. CloudLink for GCP Cloud Provider.

Upstream description:

CloudLink for GCP Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]"
}
```

## Direct properties

- [byoc](data-sources--cloud_link--properties--gcp--byoc.md): complete subsection reference.

- [gcp_cred](data-sources--cloud_link--properties--gcp--gcp_cred.md): complete subsection reference.

## Next pages

- [gcp.byoc](data-sources--cloud_link--properties--gcp--byoc.md)
- [gcp.gcp_cred](data-sources--cloud_link--properties--gcp--gcp_cred.md)
- [Property reference](data-sources--cloud_link--reference.md)
- [xcsh_cloud_link](../data-sources/cloud_link.md)
