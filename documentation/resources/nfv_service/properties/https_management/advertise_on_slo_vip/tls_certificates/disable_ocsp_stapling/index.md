---
page_title: "https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https management advertise on slo vip tls certificates disable ocsp stapling"], "body_bytes": 1501, "body_sha256": "sha256:a29e421ae6f2fcb8cfc4c3c83f6b52fb8ceb9a769baa7a101b12f0a5fa298902", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates", "path": "documentation/resources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_certificates/disable_ocsp_stapling/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1223312232131130-3323102031233232-3103230230232313-0023320033320202-0210000233222211-3012010210210210-2201000102130223-2220232000002001", "registry_path": "docs/guides/resources--nfv_service--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_vip", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_vip/)
- [https_management.advertise_on_slo_vip.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_certificates/)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.
