---
page_title: "rules.ip_and_tls_fingerprint"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules ip and tls fingerprint"], "body_bytes": 1008, "body_sha256": "sha256:801e2aec226bda53d8bc7592d83c25700e60845b593681d4f93b34cd3fb5055d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:ip_and_tls_fingerprint", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "documentation/resources/user_identification/properties/rules/ip_and_tls_fingerprint/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2311120321201223-3201102323200101-2023020101303322-1232011222231021-1121223301223211-1231030121302331-3123211030232003-0323001332231000", "registry_path": "docs/guides/resources--user_identification--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "ip_and_tls_fingerprint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/ip_and_tls_fingerprint/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["user_identificationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.ip_and_tls_fingerprint

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/rules/)
- rules.ip_and_tls_fingerprint

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
ip_and_tls_fingerprint = {}
```

This is an empty object or choice marker. It has no direct properties.
