---
page_title: "palo_alto_fw_service.panorama_server.authorization_key"
subcategory: ""
description: "palo_alto_fw_service.panorama_server.authorization_key for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2420, "body_sha256": "sha256:020b1d94196f237dcfa40dd603b7b138e9091473d3a3691c1683b675a3e2b077", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "palo_alto_fw_service.panorama_server.authorization_key for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.panorama_server.authorization_key

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.panorama_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/)
- palo_alto_fw_service.panorama_server.authorization_key

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/clear_secret_info/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/blindfold_secret_info/)
- [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/clear_secret_info/)
- [palo_alto_fw_service.panorama_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
