---
page_title: "rules.tls_fingerprint"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules tls fingerprint"], "body_bytes": 964, "body_sha256": "sha256:ddb4fbe292c3b14aa24fabc55c009c4f1a81e8a753c4540a6445e858ec2167bf", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:user_identification:properties:rules:tls_fingerprint", "parent_id": "xcsh-docs:data-sources:user_identification:properties:rules", "path": "documentation/data-sources/user_identification/properties/rules/tls_fingerprint/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3223201212222123-1320213332302222-3033301310213001-0232101010110322-1130011101033201-2231133010313101-0010013203310202-0332121310312003", "registry_path": "docs/guides/data-sources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "tls_fingerprint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/properties/rules/tls_fingerprint/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["user_identificationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/properties/rules/)
- rules.tls_fingerprint

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for tls fingerprint.

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
