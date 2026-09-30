---
page_title: "https_management.advertise_on_internet"
subcategory: ""
description: "https_management.advertise_on_internet for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1450, "body_sha256": "sha256:34a9b8d5f48d2a0e70ed02cfd6169d59fedf537b8f298746651038251f1978fb", "canonical_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet", "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet:public_ip"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_internet", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management", "path": "docs/guides/resources--nfv_service--properties--https_management--advertise_on_internet.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_internet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_internet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_internet for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_management.advertise_on_internet

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- https_management.advertise_on_internet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_internet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [public_ip](resources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_internet.public_ip](resources--nfv_service--properties--https_management--advertise_on_internet--public_ip.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
