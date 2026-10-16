---
page_title: "authentication.cookie_params.kms_key_hmac"
subcategory: ""
description: "Reference to KMS Key Object."
xcsh_docs: {"aliases": ["authentication cookie params kms key hmac"], "body_bytes": 1119, "body_sha256": "sha256:86915f21566d216fc936b118c80158ce3c93afc467da731bb9686324899a199f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params:kms_key_hmac", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:authentication:cookie_params", "path": "documentation/data-sources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3310021112030201-0031131211202322-2133131213000101-1010122133011302-2132212001303011-0213113310233213-2322012011303002-3120321021213033", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "kms_key_hmac"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Reference to KMS Key Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.kms_key_hmac

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/authentication/cookie_params/)
- authentication.cookie_params.kms_key_hmac

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for kms key hmac.

Additional upstream details:

Reference to KMS Key Object.

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

This is an empty object or choice marker. It has no direct properties.
