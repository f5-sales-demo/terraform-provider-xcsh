---
page_title: "palo_alto_fw_service.panorama_server.authorization_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["palo alto fw service panorama server authorization key"], "body_bytes": 2704, "body_sha256": "sha256:2420fe96ab66f6165fb70e904de72f3330aaf1770e6d79b79e759e1febd6bad6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key", "parent_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server.authorization_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server.authorization_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--panorama_server--authorization_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--panorama_server--authorization_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service.panorama_server.authorization_key

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [palo_alto_fw_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/)
- [palo_alto_fw_service.panorama_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/)
- palo_alto_fw_service.panorama_server.authorization_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
authorization_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/clear_secret_info/): complete subsection reference.

## Next pages

- [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/blindfold_secret_info/)
- [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/clear_secret_info/)
- [palo_alto_fw_service.panorama_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
