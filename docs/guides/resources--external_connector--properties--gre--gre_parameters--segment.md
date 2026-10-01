---
page_title: "gre.gre_parameters.segment"
subcategory: ""
description: "gre.gre_parameters.segment for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1479, "body_sha256": "sha256:afe33d509cb61c8712838bd2fededda053fd3c3df0d24728f040d3284fa007c2", "canonical_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment:refs"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "parent_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "path": "docs/guides/resources--external_connector--properties--gre--gre_parameters--segment.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gre", "gre_parameters", "segment"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/segment/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre.gre_parameters.segment for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.segment

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [gre](resources--external_connector--properties--gre.md)
- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- gre.gre_parameters.segment

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
```

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
segment {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](resources--external_connector--properties--gre--gre_parameters--segment--refs.md): complete subsection reference.

## Next pages

- [gre.gre_parameters.segment.refs](resources--external_connector--properties--gre--gre_parameters--segment--refs.md)
- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
