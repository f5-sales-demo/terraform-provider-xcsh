---
page_title: "sriov_interfaces"
subcategory: ""
description: "sriov_interfaces for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1078, "body_sha256": "sha256:4ff4132c4ac6be91ab637408f1c7d9462b0530d4eb24bf3b0842cc253b3d96a8", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:sriov_interfaces:sriov_interface"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:sriov_interfaces", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "docs/guides/resources--voltstack_site--properties--sriov_interfaces.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sriov_interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/sriov_interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sriov_interfaces for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sriov_interfaces

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- sriov_interfaces

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

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

Terraform syntax:

```terraform
sriov_interfaces {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sriov_interface](resources--voltstack_site--properties--sriov_interfaces--sriov_interface.md): complete subsection reference.

## Next pages

- [sriov_interfaces.sriov_interface](resources--voltstack_site--properties--sriov_interfaces--sriov_interface.md)
- [Property reference](resources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
