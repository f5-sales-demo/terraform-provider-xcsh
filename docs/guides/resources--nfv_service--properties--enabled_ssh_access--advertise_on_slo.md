---
page_title: "enabled_ssh_access.advertise_on_slo"
subcategory: ""
description: "enabled_ssh_access.advertise_on_slo for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1045, "body_sha256": "sha256:72236c86902133d1b1959acf30c1c9735e059d7b064fcf50e504ada96dd426eb", "canonical_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:advertise_on_slo", "parent_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "path": "docs/guides/resources--nfv_service--properties--enabled_ssh_access--advertise_on_slo.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enabled_ssh_access", "advertise_on_slo"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/enabled_ssh_access/advertise_on_slo/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enabled_ssh_access.advertise_on_slo for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled_ssh_access.advertise_on_slo

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [enabled_ssh_access](resources--nfv_service--properties--enabled_ssh_access.md)
- enabled_ssh_access.advertise_on_slo

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on slo.

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
advertise_on_slo = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enabled_ssh_access](resources--nfv_service--properties--enabled_ssh_access.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
