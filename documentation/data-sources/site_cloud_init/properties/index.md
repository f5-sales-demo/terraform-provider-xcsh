---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_cloud_init."
xcsh_docs: {"aliases": [], "body_bytes": 1939, "body_sha256": "sha256:2769a0aa560c7cddd3df8791a5ddc07aae767e85765db0f1fda1bb2d0d073f18", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_cloud_init:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_cloud_init:reference", "parent_id": "xcsh-docs:data-sources:site_cloud_init:fundamentals", "path": "documentation/data-sources/site_cloud_init/properties/index.md", "provider_name": "site_cloud_init", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_cloud_init/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_cloud_init.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_site_cloud_init](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/)
- Property reference

## Direct properties

<a id="schema-cloud_init_config"></a>

### cloud_init_config property

Type: `"string"`. Computed, Sensitive.

Cloud-init template with an unresolved token placeholder; substitute a separately issued site-bound
JWT before deployment. This sensitive value is stored in Terraform state; protect state access
accordingly.

<a id="schema-enable_management_network"></a>

### enable_management_network property

Type: `"bool"`. Optional.

Management network choice for this cloud-init config.

<a id="schema-provider_ref"></a>

### provider_ref property

Type: `"string"`. Required.

Provider for that cloud-init config.

<a id="schema-site_name"></a>

### site_name property

Type: `"string"`. Required.

Site name for this cloud-init config.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cloud_init_config` | [cloud_init_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-cloud_init_config) |
| `enable_management_network` | [enable_management_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-enable_management_network) |
| `provider_ref` | [provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-provider_ref) |
| `site_name` | [site_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/properties/#schema-site_name) |

## Next pages

- [xcsh_site_cloud_init](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/)
