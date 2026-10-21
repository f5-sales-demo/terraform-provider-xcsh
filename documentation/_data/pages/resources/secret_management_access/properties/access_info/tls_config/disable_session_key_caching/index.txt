---
page_title: "access_info.tls_config.disable_session_key_caching"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["access info tls config disable session key caching"], "body_bytes": 1281, "body_sha256": "sha256:1f857b79feda8f29f7e6983deae5275ff395760caa0e9e878022669eded2713b", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:disable_session_key_caching", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/disable_session_key_caching/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2203002320210003-3112320211113020-1231121133010311-1233120131013313-3001311220012032-3222010112020103-2122300120312001-3113010110113312", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "disable_session_key_caching"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/disable_session_key_caching/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.disable_session_key_caching

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- access_info.tls_config.disable_session_key_caching

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.
