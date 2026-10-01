---
page_title: "gre.gre_parameters.site_local_network"
subcategory: ""
description: "gre.gre_parameters.site_local_network for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1134, "body_sha256": "sha256:d647bcae1f88c618f1d64571256728d3b82bbd787263b8c293cb348a5f93d212", "canonical_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "child_ids": [], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "parent_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "path": "docs/guides/resources--external_connector--properties--gre--gre_parameters--site_local_network.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gre", "gre_parameters", "site_local_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/site_local_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre.gre_parameters.site_local_network for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.site_local_network

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [gre](resources--external_connector--properties--gre.md)
- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- gre.gre_parameters.site_local_network

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
site_local_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
