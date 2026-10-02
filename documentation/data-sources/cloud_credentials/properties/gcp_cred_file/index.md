---
page_title: "gcp_cred_file"
subcategory: "Infrastructure"
description: "GCP Credentials type."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "gcp cred file"], "body_bytes": 1325, "body_sha256": "sha256:d41e0e6935cf9c30e30e9fe80f688ded01e2fc2346ceaf5c578cd9a08962f3ca", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file", "parent_id": "xcsh-docs:data-sources:cloud_credentials:reference", "path": "documentation/data-sources/cloud_credentials/properties/gcp_cred_file/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2031333012222021-0100332121232333-2302122322021100-2113313022131132-1233130233203333-0122233121322110-2303200103012231-1203013021303220", "registry_path": "docs/guides/data-sources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_cred_file"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential file", "credential setup", "credentials"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:cloud_credentials:properties:gcp_cred_file:credential_file", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp_cred_file", "credential_file"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/properties/gcp_cred_file/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "GCP Credentials type.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_cred_file

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
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

- [credential_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/gcp_cred_file/credential_file/): complete subsection reference.

## Next pages

- [gcp_cred_file.credential_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/gcp_cred_file/credential_file/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/properties/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
