---
page_title: "rules.ja4_tls_fingerprint"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules ja4 tls fingerprint"], "body_bytes": 976, "body_sha256": "sha256:5a19d1573720d75ec1e480107df2edeb8578114d1c9a4f62810a1802e23657dc", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:user_identification:properties:rules:ja4_tls_fingerprint", "parent_id": "xcsh-docs:data-sources:user_identification:properties:rules", "path": "documentation/data-sources/user_identification/properties/rules/ja4_tls_fingerprint/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1333132302231133-1012010222321122-3100100122132320-2333322113103220-3113003000022031-2300110231011113-2222222311221022-3001230000222003", "registry_path": "docs/guides/data-sources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "ja4_tls_fingerprint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/properties/rules/ja4_tls_fingerprint/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["user_identificationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ja4_tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/)
- rules.ja4_tls_fingerprint

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ja4 tls fingerprint.

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

This is an empty object or choice marker. It has no direct properties.
