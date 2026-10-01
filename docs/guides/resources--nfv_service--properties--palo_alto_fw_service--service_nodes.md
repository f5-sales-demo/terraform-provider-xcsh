---
page_title: "palo_alto_fw_service.service_nodes"
subcategory: ""
description: "palo_alto_fw_service.service_nodes for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1349, "body_sha256": "sha256:6a9a2443dbd10008327e8758b4f6ed12640f72976cb31eea447c29a4594339ab", "canonical_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes", "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "path": "docs/guides/resources--nfv_service--properties--palo_alto_fw_service--service_nodes.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["palo_alto_fw_service", "service_nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.service_nodes for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.service_nodes

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- palo_alto_fw_service.service_nodes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for service nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nodes")}
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
service_nodes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md): complete subsection reference.

## Next pages

- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--properties--palo_alto_fw_service--service_nodes--nodes.md)
- [palo_alto_fw_service](resources--nfv_service--properties--palo_alto_fw_service.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
