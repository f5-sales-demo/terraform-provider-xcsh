---
page_title: "gcp_cred_file"
subcategory: "Infrastructure"
description: "gcp_cred_file for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1017, "body_sha256": "sha256:626483d0fe5ba5cf3e4396d8afc4fa96c5310712d5290f853abeeab0f200a5ba", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "docs/guides/data-sources--cloud_credentials--properties--gcp_cred_file.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp_cred_file"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/gcp_cred_file/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_cred_file for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file

Breadcrumbs:

- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
- gcp_cred_file

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for gcp cred file.

Upstream description:

GCP Credentials type.

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

- [credential_file](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file.md): complete subsection reference.

## Next pages

- [gcp_cred_file.credential_file](data-sources--cloud_credentials--properties--gcp_cred_file--credential_file.md)
- [Property reference](data-sources--cloud_credentials--reference.md)
- [xcsh_cloud_credentials](../data-sources/cloud_credentials.md)
