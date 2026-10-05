---
page_title: "https_management.advertise_on_sli_vip.tls_certificates.private_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["https management advertise on sli vip tls certificates private key"], "body_bytes": 2754, "body_sha256": "sha256:6df512374277707a44f937648e4fc652f0e5f89c67854440eb7e3ee313bd55c3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates", "path": "documentation/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3202013233313032-1122321011023203-3110102003123020-2123222112301232-3113233132130213-0123211321312332-1213330030132020-2001231220212130", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_sli_vip", "tls_certificates", "private_key"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on sli vip tls certificates private key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_certificates", "private_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on sli vip tls certificates private key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_certificates", "private_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_sli_vip.tls_certificates.private_key

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/)
- [https_management.advertise_on_sli_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/)
- [https_management.advertise_on_sli_vip.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/)
- https_management.advertise_on_sli_vip.tls_certificates.private_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/blindfold_secret_info/)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/private_key/clear_secret_info/)
- [https_management.advertise_on_sli_vip.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
