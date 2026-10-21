---
page_title: "cookie_params.kms_key_hmac"
subcategory: ""
description: "Reference to KMS Key Object."
xcsh_docs: {"aliases": ["cookie params kms key hmac"], "body_bytes": 938, "body_sha256": "sha256:3113a0ecad88c158de35d54e488db53fe66fb9afd4a54a8ea21ae8759001bd3d", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:properties:cookie_params:kms_key_hmac", "parent_id": "xcsh-docs:data-sources:authentication:properties:cookie_params", "path": "documentation/data-sources/authentication/properties/cookie_params/kms_key_hmac/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1122032132211133-3221102303211012-3310003020230231-2020330020233321-2301211333312333-2220222021012130-3120332310133212-1100101331103330", "registry_path": "docs/guides/data-sources--authentication--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cookie_params", "kms_key_hmac"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/properties/cookie_params/kms_key_hmac/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reference to KMS Key Object.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["authenticationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cookie_params.kms_key_hmac

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/)
- [cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/cookie_params/)
- cookie_params.kms_key_hmac

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
