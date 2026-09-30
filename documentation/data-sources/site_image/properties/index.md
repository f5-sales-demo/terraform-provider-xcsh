---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_image."
xcsh_docs: {"aliases": [], "body_bytes": 1813, "body_sha256": "sha256:7397ba63b98e5bec0b748867008b0373e04ef425b7b42f050130a7b238bf15ce", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:reference", "parent_id": "xcsh-docs:data-sources:site_image:fundamentals", "path": "documentation/data-sources/site_image/properties/index.md", "provider_name": "site_image", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_image.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
