---
page_title: "palo_alto_fw_service.panorama_server.authorization_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["palo alto fw service panorama server authorization key"], "body_bytes": 1632, "body_sha256": "sha256:75c01b4b82f1f30a9fae1cd1c117e0b24cd85c681de315a45d52195d4ea0b861", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server", "path": "documentation/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0002100233231111-3231112032100013-0122210001230213-3012100312210232-3210112021330202-3133123131231210-3312230210131132-3012131020131121", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key"], "schema_version": 1, "sections": [{"aliases": ["palo alto fw service panorama server authorization key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["palo alto fw service panorama server authorization key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:palo_alto_fw_service:panorama_server:authorization_key:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["palo_alto_fw_service", "panorama_server", "authorization_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/palo_alto_fw_service/panorama_server/authorization_key/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
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
