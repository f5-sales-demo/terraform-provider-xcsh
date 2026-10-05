---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_image."
xcsh_docs: {"aliases": ["site image"], "body_bytes": 1912, "body_sha256": "sha256:a36ccf2bc427ee5f9bab7a16ba207f70cd8c62e46d533f0c4c1a441038bf1e05", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:reference", "parent_id": "xcsh-docs:data-sources:site_image:fundamentals", "path": "documentation/data-sources/site_image/properties/index.md", "product": "distributed-cloud", "provider_name": "site_image", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3230203203002131-2312110021103212-3200131121210123-1302203223102322-2302331310010222-3030113311121323-3030201010300123-3222123123130310", "registry_path": "docs/guides/data-sources--site_image--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["image download url"], "anchor": "schema-image_download_url", "description": "Validated HTTPS image URL. Protect Terraform state.", "document_id": "xcsh-docs:data-sources:site_image:reference", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_download_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["image md5 sum"], "anchor": "schema-image_md5_sum", "description": "Expected artifact MD5, which the consumer must verify before boot.", "document_id": "xcsh-docs:data-sources:site_image:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_md5_sum"], "syntax": "attribute", "type": "string"}, {"aliases": ["image name"], "anchor": "schema-image_name", "description": "Image name returned by the Site-UID query. May contain a download URL; protect Terraform state.", "document_id": "xcsh-docs:data-sources:site_image:reference", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["site name"], "anchor": "schema-site_name", "description": "Existing KVM SMSv2 configuration name in system. Ownership is revalidated on each read.", "document_id": "xcsh-docs:data-sources:site_image:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/properties/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Property reference for xcsh_site_image.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [xcsh_site_image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/)
