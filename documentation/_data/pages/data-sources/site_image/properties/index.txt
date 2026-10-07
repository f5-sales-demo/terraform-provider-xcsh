---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_image."
xcsh_docs: {"aliases": ["site image"], "body_bytes": 1794, "body_sha256": "sha256:6ba4835322996bf51aafc05e9b769ef3263e822efa2f4fbaf5ab5627b2cde97a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:reference", "parent_id": "xcsh-docs:data-sources:site_image:fundamentals", "path": "documentation/data-sources/site_image/properties/index.md", "product": "distributed-cloud", "provider_name": "site_image", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3230203203002131-2312110021103212-3200131121210123-1302203223102322-2302331310010222-3030113311121323-3030201010300123-3222123123130310", "registry_path": "docs/guides/data-sources--site_image--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["image download url"], "anchor": "schema-image_download_url", "description": "Validated HTTPS image URL. Protect Terraform state.", "document_id": "xcsh-docs:data-sources:site_image:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_download_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["image md5 sum"], "anchor": "schema-image_md5_sum", "description": "Expected artifact MD5, which the consumer must verify before boot.", "document_id": "xcsh-docs:data-sources:site_image:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_md5_sum"], "syntax": "attribute", "type": "string"}, {"aliases": ["image name"], "anchor": "schema-image_name", "description": "Image name returned by the Site-UID query. May contain a download URL; protect Terraform state.", "document_id": "xcsh-docs:data-sources:site_image:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["site name"], "anchor": "schema-site_name", "description": "Existing KVM SMSv2 configuration name in system. Ownership is revalidated on each read.", "document_id": "xcsh-docs:data-sources:site_image:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/properties/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Property reference for xcsh_site_image.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/)
- Property reference

## Direct properties

<a id="schema-image_download_url"></a>

### image_download_url property

Type: `"string"`. Computed, Sensitive.

Validated HTTPS image URL. Protect Terraform state.

<a id="schema-image_md5_sum"></a>

### image_md5_sum property

Type: `"string"`. Computed.

Expected artifact MD5, which the consumer must verify before boot.

<a id="schema-image_name"></a>

### image_name property

Type: `"string"`. Computed, Sensitive.

Image name returned by the Site-UID query. May contain a download URL; protect Terraform state.

<a id="schema-site_name"></a>

### site_name property

Type: `"string"`. Required.

Existing KVM SMSv2 configuration name in system. Ownership is revalidated on each read.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `image_download_url` | [image_download_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/#schema-image_download_url) |
| `image_md5_sum` | [image_md5_sum](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/#schema-image_md5_sum) |
| `image_name` | [image_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/#schema-image_name) |
| `site_name` | [site_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/#schema-site_name) |
