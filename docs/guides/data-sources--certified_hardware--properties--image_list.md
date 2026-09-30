---
page_title: "image_list"
subcategory: ""
description: "image_list for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 1389, "body_sha256": "sha256:492b623e6c39f0c2f635248ac9f237905ef84b08b20e577e2e10fa39ea69a974", "canonical_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list", "child_ids": ["xcsh-docs:data-sources:certified_hardware:properties:image_list:aws", "xcsh-docs:data-sources:certified_hardware:properties:image_list:azure", "xcsh-docs:data-sources:certified_hardware:properties:image_list:gcp"], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:image_list", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "docs/guides/data-sources--certified_hardware--properties--image_list.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["image_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/image_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "image_list for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# image_list

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
- [Property reference](data-sources--certified_hardware--reference.md)
- image_list

<a id="section"></a>

Type: `"list"`. Computed.

List of image names with providers for this certified hardware, e.g. AWS ami-0f99d090261d2acd5.

## Direct properties

- [aws](data-sources--certified_hardware--properties--image_list--aws.md): complete subsection reference.

- [azure](data-sources--certified_hardware--properties--image_list--azure.md): complete subsection reference.

- [gcp](data-sources--certified_hardware--properties--image_list--gcp.md): complete subsection reference.

<a id="schema-image_list--name"></a>

### name property

Type: `"string"`. Computed.

Name. Image name to use for this hardware.

<a id="schema-image_list--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Image provider F5 Distributed Cloud, Cloud provider like AWS or Azure.

## Next pages

- [image_list.aws](data-sources--certified_hardware--properties--image_list--aws.md)
- [image_list.azure](data-sources--certified_hardware--properties--image_list--azure.md)
- [image_list.gcp](data-sources--certified_hardware--properties--image_list--gcp.md)
- [Property reference](data-sources--certified_hardware--reference.md)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
