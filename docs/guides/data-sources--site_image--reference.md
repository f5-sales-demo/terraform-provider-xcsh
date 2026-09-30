---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_image."
xcsh_docs: {"aliases": [], "body_bytes": 1499, "body_sha256": "sha256:b03d33975ae6211e6ece316a4263fb497ffa4762575c6cdc89202a837979c753", "canonical_id": "xcsh-docs:data-sources:site_image:reference", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:reference", "parent_id": "xcsh-docs:data-sources:site_image:fundamentals", "path": "docs/guides/data-sources--site_image--reference.md", "provider_name": "site_image", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_image.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_site_image](../data-sources/site_image.md)
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
| `image_download_url` | [image_download_url](data-sources--site_image--reference.md#schema-image_download_url) |
| `image_md5_sum` | [image_md5_sum](data-sources--site_image--reference.md#schema-image_md5_sum) |
| `image_name` | [image_name](data-sources--site_image--reference.md#schema-image_name) |
| `site_name` | [site_name](data-sources--site_image--reference.md#schema-site_name) |

## Next pages

- [xcsh_site_image](../data-sources/site_image.md)
